package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerNetworkAttachmentRule checks network attachment configuration
type IBMIsBareMetalServerNetworkAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerNetworkAttachmentRule() *IBMIsBareMetalServerNetworkAttachmentRule {
	return &IBMIsBareMetalServerNetworkAttachmentRule{}
}

func (r *IBMIsBareMetalServerNetworkAttachmentRule) Name() string {
	return "ibm_is_bare_metal_server_network_attachment"
}

func (r *IBMIsBareMetalServerNetworkAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerNetworkAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerNetworkAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerNetworkAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_network_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "interface_type"},
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
		requiredAttrs := []string{"bare_metal_server"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate interface_type if specified
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var interfaceType string
			err := runner.EvaluateExpr(attr.Expr, &interfaceType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"vlan": true,
				"pci":  true,
			}

			if !validTypes[interfaceType] {
				runner.EmitIssue(
					r,
					"interface_type must be either 'vlan' or 'pci'",
					attr.Expr.Range(),
				)
			}

			// Check type-specific requirements
			if interfaceType == "vlan" {
				if _, exists := resource.Body.Attributes["vlan"]; !exists {
					runner.EmitIssue(
						r,
						"vlan must be specified when interface_type is 'vlan'",
						attr.Expr.Range(),
					)
				}
			} else if interfaceType == "pci" {
				if _, exists := resource.Body.Attributes["allowed_vlans"]; !exists {
					runner.EmitIssue(
						r,
						"allowed_vlans must be specified when interface_type is 'pci'",
						attr.Expr.Range(),
					)
				}
			}
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
						"all VLAN values in allowed_vlans must be between 1 and 4094",
						attr.Expr.Range(),
					)
					break
				}
			}
		}
	}

	return nil
}
