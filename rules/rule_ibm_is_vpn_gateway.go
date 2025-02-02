package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPNGatewayRule checks VPN gateway configuration
type IBMIsVPNGatewayRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPNGatewayRule() *IBMIsVPNGatewayRule {
	return &IBMIsVPNGatewayRule{}
}

func (r *IBMIsVPNGatewayRule) Name() string {
	return "ibm_is_vpn_gateway"
}

func (r *IBMIsVPNGatewayRule) Enabled() bool {
	return true
}

func (r *IBMIsVPNGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPNGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPNGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpn_gateway", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "subnet"},
			{Name: "mode"},
			{Name: "resource_group"},
			{Name: "tags"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "subnet"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}

		// Validate mode if specified
		if attr, exists := resource.Body.Attributes["mode"]; exists {
			var mode string
			err := runner.EvaluateExpr(attr.Expr, &mode, nil)
			if err != nil {
				return err
			}

			validModes := map[string]bool{
				"policy": true,
				"route":  true,
			}

			if !validModes[mode] {
				runner.EmitIssue(
					r,
					"mode must be either 'policy' or 'route'",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
