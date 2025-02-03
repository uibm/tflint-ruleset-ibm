package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerInitializationRule checks bare metal server initialization configuration
type IBMIsBareMetalServerInitializationRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerInitializationRule() *IBMIsBareMetalServerInitializationRule {
	return &IBMIsBareMetalServerInitializationRule{}
}

func (r *IBMIsBareMetalServerInitializationRule) Name() string {
	return "ibm_is_bare_metal_server_initialization"
}

func (r *IBMIsBareMetalServerInitializationRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerInitializationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerInitializationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerInitializationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_initialization", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "image"},
			{Name: "keys"},
			{Name: "user_data"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "image", "keys"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate keys array
		if attr, exists := resource.Body.Attributes["keys"]; exists {
			var keys []string
			err := runner.EvaluateExpr(attr.Expr, &keys, nil)
			if err != nil {
				return err
			}

			if len(keys) == 0 {
				runner.EmitIssue(
					r,
					"at least one SSH key must be specified",
					attr.Expr.Range(),
				)
			}
		}

		// Validate user_data if specified
		if attr, exists := resource.Body.Attributes["user_data"]; exists {
			var userData string
			err := runner.EvaluateExpr(attr.Expr, &userData, nil)
			if err != nil {
				return err
			}

			if len(userData) > 16384 { // 16KB limit
				runner.EmitIssue(
					r,
					"user_data exceeds maximum size of 16KB",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
