package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerNetworkInterfaceFloatingIPRule checks network interface floating IP configuration
type IBMIsBareMetalServerNetworkInterfaceFloatingIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerNetworkInterfaceFloatingIPRule() *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule {
	return &IBMIsBareMetalServerNetworkInterfaceFloatingIPRule{}
}

func (r *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule) Name() string {
	return "ibm_is_bare_metal_server_network_interface_floating_ip"
}

func (r *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerNetworkInterfaceFloatingIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_network_interface_floating_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "network_interface"},
			{Name: "floating_ip"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "network_interface", "floating_ip"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Additional validations can be added here, for example:
		// - Validate floating IP format
		// - Check if network interface exists
		// - Verify floating IP availability
	}

	return nil
}
