package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVirtualEndpointGatewayIPRule checks virtual endpoint gateway IP configuration
type IBMIsVirtualEndpointGatewayIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsVirtualEndpointGatewayIPRule() *IBMIsVirtualEndpointGatewayIPRule {
	return &IBMIsVirtualEndpointGatewayIPRule{}
}

func (r *IBMIsVirtualEndpointGatewayIPRule) Name() string {
	return "ibm_is_virtual_endpoint_gateway_ip"
}

func (r *IBMIsVirtualEndpointGatewayIPRule) Enabled() bool {
	return true
}

func (r *IBMIsVirtualEndpointGatewayIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVirtualEndpointGatewayIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVirtualEndpointGatewayIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_virtual_endpoint_gateway_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "gateway"},
			{Name: "reserved_ip"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"gateway", "reserved_ip"}
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
