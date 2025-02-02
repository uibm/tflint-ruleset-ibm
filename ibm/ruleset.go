package ibm

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/rules"
)

type RuleSet struct {
	tflint.BuiltinRuleSet
	config *Config
}

func NewRuleSet() *RuleSet {
	return &RuleSet{
		BuiltinRuleSet: tflint.BuiltinRuleSet{
			Name:    "ibm",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewIBMIsInstanceRule(),
				rules.NewIBMIsVPCRule(),
			},
		},
	}
}

func (r *RuleSet) ConfigSchema() *hclext.BodySchema {
	return &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "ibmcloud_api_key", Required: false},
			{Name: "region", Required: false},
		},
	}
}

func (r *RuleSet) ApplyConfig(body *hclext.BodyContent) error {
	if r.config == nil {
		r.config = &Config{}
	}

	diags := hclext.DecodeBody(body, nil, r.config)
	if diags.HasErrors() {
		return fmt.Errorf("failed to decode configuration: %w", diags.Errs()[0])
	}

	// Validate the configuration
	if r.config.IBMCloudApiKey == "" {
		return fmt.Errorf("ibmcloud_api_key is required")
	}
	if r.config.Region == "" {
		return fmt.Errorf("region is required")
	}

	return nil
}

func (r *RuleSet) NewRunner(runner tflint.Runner) (tflint.Runner, error) {
	return NewRunner(runner, r.config)
}
