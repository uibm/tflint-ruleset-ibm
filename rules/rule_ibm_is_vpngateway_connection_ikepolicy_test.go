package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionIkepolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_vpngateway_connection_ikepolicy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkepolicyRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 58},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_vpngateway_connection_ikepolicy" "test" {
  id = "r006-e98f46a3-1e4e-4195-b4e5-b8155192689d"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkepolicyRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 58},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_vpngateway_connection_ikepolicy" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/r006-e98f46a3-1e4e-4195-b4e5-b8155192689d"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkepolicyRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 58},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_connection_ikepolicy" "test" {
  id = "r006-e98f46a3-1e4e-4195-b4e5-b8155192689d"
  href = "https://us-south.iaas.cloud.ibm.com/v1/ike_policies/r006-e98f46a3-1e4e-4195-b4e5-b8155192689d"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionIkepolicyRule()

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
