package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVirtualNetworkInterfaceRule checks virtual network interface configuration
type IBMIsVirtualNetworkInterfaceRule struct {
	tflint.DefaultRule
}

func NewIBMIsVirtualNetworkInterfaceRule() *IBMIsVirtualNetworkInterfaceRule {
	return &IBMIsVirtualNetworkInterfaceRule{}
}

func (r *IBMIsVirtualNetworkInterfaceRule) Name() string {
	return "ibm_is_virtual_network_interface"
}

func (r *IBMIsVirtualNetworkInterfaceRule) Enabled() bool {
	return true
}

func (r *IBMIsVirtualNetworkInterfaceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVirtualNetworkInterfaceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVirtualNetworkInterfaceRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_virtual_network_interface", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "subnet"},
			{Name: "allow_ip_spoofing"},
			{Name: "auto_delete"},
			{Name: "enable_infrastructure_nat"},
			{Name: "protocol_state_filtering_mode"},
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

		// Validate protocol_state_filtering_mode if specified
		if attr, exists := resource.Body.Attributes["protocol_state_filtering_mode"]; exists {
			var mode string
			err := runner.EvaluateExpr(attr.Expr, &mode, nil)
			if err != nil {
				return err
			}

			validModes := map[string]bool{
				"enabled":  true,
				"disabled": true,
			}

			if !validModes[mode] {
				runner.EmitIssue(
					r,
					"protocol_state_filtering_mode must be either 'enabled' or 'disabled'",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
