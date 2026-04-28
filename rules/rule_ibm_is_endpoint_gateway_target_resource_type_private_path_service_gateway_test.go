package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: crn",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule(),
					Message: "crn attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 93},
					},
				},
			},
		},
		{
			Name: "invalid enum value for resource_type",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway" "test" {
  crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::private-path-service-gateway:r006-7268d425-59b7-48fd-9735-81a7271657d5"
  resource_type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule(),
					Message: "'resource_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 93},
					},
				},
			},
		},
		{
			Name: "string too short for crn",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway" "test" {
  crn = "x"
  resource_type = "private_path_service_gateway"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule(),
					Message: "'crn' value must be at least 17 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 93},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for crn",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway" "test" {
  crn = "invalid-format"
  resource_type = "private_path_service_gateway"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule(),
					Message: "'crn' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 93},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway" "test" {
  crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::private-path-service-gateway:r006-7268d425-59b7-48fd-9735-81a7271657d5"
  resource_type = "private_path_service_gateway"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule()

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
