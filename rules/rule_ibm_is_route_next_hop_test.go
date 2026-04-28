package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMRouteNextHopRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: address",
			Content: `resource "ibm_is_route_next_hop" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteNextHopRule(),
					Message: "address attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_route_next_hop" "test" {
  address = "test-value"
  id = "0717-52f69dc3-6a5c-4bcf-b264-e7fae279b15c"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteNextHopRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for href",
			Content: `resource "ibm_is_route_next_hop" "test" {
  address = "test-value"
  id = "0717-52f69dc3-6a5c-4bcf-b264-e7fae279b15c"
  href = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteNextHopRule(),
					Message: "'href' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_route_next_hop" "test" {
  address = "test-value"
  id = "0717-52f69dc3-6a5c-4bcf-b264-e7fae279b15c"
  href = "https://us-south.iaas.cloud.ibm.com/v1/vpn_gateways/0717-ddf51bec-3424-11e8-b467-0ed5f89f718b/connections/0717-52f69dc3-6a5c-4bcf-b264-e7fae279b15c"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMRouteNextHopRule()

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
