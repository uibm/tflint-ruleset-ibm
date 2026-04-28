package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMEndpointGatewayTargetResourceTypeProviderInfrastructureServiceRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: name",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_provider_infrastructure_service" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypeProviderInfrastructureServiceRule(),
					Message: "name attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 96},
					},
				},
			},
		},
		{
			Name: "invalid enum value for resource_type",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_provider_infrastructure_service" "test" {
  name = "ibm-ntp-server"
  resource_type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypeProviderInfrastructureServiceRule(),
					Message: "'resource_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 96},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_provider_infrastructure_service" "test" {
  name = "invalid-format"
  resource_type = "provider_infrastructure_service"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypeProviderInfrastructureServiceRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 96},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_provider_infrastructure_service" "test" {
  name = "ibm-ntp-server"
  resource_type = "provider_infrastructure_service"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMEndpointGatewayTargetResourceTypeProviderInfrastructureServiceRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
