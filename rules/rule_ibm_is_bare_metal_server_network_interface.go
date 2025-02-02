package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerNetworkInterfaceRule checks network interface configuration
type IBMIsBareMetalServerNetworkInterfaceRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerNetworkInterfaceRule() *IBMIsBareMetalServerNetworkInterfaceRule {
	return &IBMIsBareMetalServerNetworkInterfaceRule{}
}

func (r *IBMIsBareMetalServerNetworkInterfaceRule) Name() string {
	return "ibm_is_bare_metal_server_network_interface"
}

func (r *IBMIsBareMetalServerNetworkInterfaceRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerNetworkInterfaceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerNetworkInterfaceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerNetworkInterfaceRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_network_interface", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "subnet"},
			{Name: "name"},
			{Name: "allow_ip_spoofing"},
			{Name: "allowed_vlans"},
			{Name: "vlan"},
			{Name: "security_groups"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "subnet", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check for mutually exclusive attributes: allowed_vlans and vlan
		hasAllowedVlans := false
		hasVlan := false

		if _, exists := resource.Body.Attributes["allowed_vlans"]; exists {
			hasAllowedVlans = true
		}
		if _, exists := resource.Body.Attributes["vlan"]; exists {
			hasVlan = true
		}

		if hasAllowedVlans && hasVlan {
			runner.EmitIssue(
				r,
				"cannot specify both allowed_vlans and vlan",
				resource.DefRange,
			)
		}

		// Validate VLAN range if specified
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

		// Validate allowed_vlans if specified
		if attr, exists := resource.Body.Attributes["allowed_vlans"]; exists {
			var vlans []int
			err := runner.EvaluateExpr(attr.Expr, &vlans, nil)
			if err != nil {
				return err
			}

			for _, vlan := range vlans {
				if vlan < 1 || vlan > 4094 {
					runner.EmitIssue(
						r,
						"all VLAN values must be between 1 and 4094",
						attr.Expr.Range(),
					)
					break
				}
			}
		}

		// Validate allow_ip_spoofing if specified
		if attr, exists := resource.Body.Attributes["allow_ip_spoofing"]; exists {
			var allowIpSpoofing bool
			err := runner.EvaluateExpr(attr.Expr, &allowIpSpoofing, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
