package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMEndpointGatewayTargetRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: crn",
			Content: `resource "ibm_is_endpoint_gateway_target" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetRule(),
					Message: "crn attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 50},
					},
				},
			},
		},
		{
			Name: "invalid enum value for resource_type",
			Content: `resource "ibm_is_endpoint_gateway_target" "test" {
  crn = "crn:v1:bluemix:public:cloudant:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:3527280b-9327-4411-8020-591092e60353::"
  resource_type = "invalid-value"
  name = "ibm-ntp-server"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetRule(),
					Message: "'resource_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 50},
					},
				},
			},
		},
		{
			Name: "string too short for crn",
			Content: `resource "ibm_is_endpoint_gateway_target" "test" {
  crn = "x"
  resource_type = "provider_cloud_service"
  name = "ibm-ntp-server"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetRule(),
					Message: "'crn' value must be at least 17 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 50},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for crn",
			Content: `resource "ibm_is_endpoint_gateway_target" "test" {
  crn = "invalid-format"
  resource_type = "provider_cloud_service"
  name = "ibm-ntp-server"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMEndpointGatewayTargetRule(),
					Message: "'crn' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 50},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_endpoint_gateway_target" "test" {
  crn = "crn:v1:bluemix:public:cloudant:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34:3527280b-9327-4411-8020-591092e60353::"
  resource_type = "provider_cloud_service"
  name = "ibm-ntp-server"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMEndpointGatewayTargetRule()

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
