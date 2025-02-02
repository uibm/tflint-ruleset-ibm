package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVirtualNetworkInterfaceFloatingIPRule checks virtual network interface floating IP configuration
type IBMIsVirtualNetworkInterfaceFloatingIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsVirtualNetworkInterfaceFloatingIPRule() *IBMIsVirtualNetworkInterfaceFloatingIPRule {
	return &IBMIsVirtualNetworkInterfaceFloatingIPRule{}
}

func (r *IBMIsVirtualNetworkInterfaceFloatingIPRule) Name() string {
	return "ibm_is_virtual_network_interface_floating_ip"
}

func (r *IBMIsVirtualNetworkInterfaceFloatingIPRule) Enabled() bool {
	return true
}

func (r *IBMIsVirtualNetworkInterfaceFloatingIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVirtualNetworkInterfaceFloatingIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVirtualNetworkInterfaceFloatingIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_virtual_network_interface_floating_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "virtual_network_interface"},
			{Name: "floating_ip"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"virtual_network_interface", "floating_ip"}
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
