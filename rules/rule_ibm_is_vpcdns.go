package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpcdnsRule checks the configuration of ibm_is_vpcdns resource
type IBMVpcdnsRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpcdnsRule returns a new rule
func NewIBMVpcdnsRule() *IBMVpcdnsRule {
	return &IBMVpcdnsRule{
		resourceType: "ibm_is_vpcdns",
	}
}

// Name returns the rule name
func (r *IBMVpcdnsRule) Name() string {
	return "vpcdns"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpcdnsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpcdnsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpcdnsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpcdnsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "resolver"},
			{Name: "enable_hub"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
