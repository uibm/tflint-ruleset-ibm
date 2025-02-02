package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerNetworkInterfaceAllowFloatRule checks network interface allow float configuration
type IBMIsBareMetalServerNetworkInterfaceAllowFloatRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerNetworkInterfaceAllowFloatRule() *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule {
	return &IBMIsBareMetalServerNetworkInterfaceAllowFloatRule{}
}

func (r *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule) Name() string {
	return "ibm_is_bare_metal_server_network_interface_allow_float"
}

func (r *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerNetworkInterfaceAllowFloatRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_network_interface_allow_float", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "subnet"},
			{Name: "name"},
			{Name: "vlan"},
			{Name: "allowed_vlans"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "subnet", "name", "vlan"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate VLAN range
		if attr, exists := resource.Body.Attributes["vlan"]; exists {
			var vlan int
			err := runner.EvaluateExpr(attr.Expr, &vlan, nil)
			if err != nil {
				return err
			}

			if vlan < 1 || vlan > 4094 {
				runner.EmitIssue(
					r,
					"vlan must be between 1 and 4094",
					attr.Expr.Range(),
				)
			}
		}

		// Validate name format if needed
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
		}
	}

	return nil
}
