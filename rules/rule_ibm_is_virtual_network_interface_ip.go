package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVirtualNetworkInterfaceIPRule checks virtual network interface IP configuration
type IBMIsVirtualNetworkInterfaceIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsVirtualNetworkInterfaceIPRule() *IBMIsVirtualNetworkInterfaceIPRule {
	return &IBMIsVirtualNetworkInterfaceIPRule{}
}

func (r *IBMIsVirtualNetworkInterfaceIPRule) Name() string {
	return "ibm_is_virtual_network_interface_ip"
}

func (r *IBMIsVirtualNetworkInterfaceIPRule) Enabled() bool {
	return true
}

func (r *IBMIsVirtualNetworkInterfaceIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVirtualNetworkInterfaceIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVirtualNetworkInterfaceIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_virtual_network_interface_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "reserved_ip"},
			{Name: "virtual_network_interface"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"reserved_ip", "virtual_network_interface"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}
	}

	return nil
}
