package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayPolicyModeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: mode",
			Content: `resource "ibm_is_vpngateway_policy_mode" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayPolicyModeRule(),
					Message: "mode attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "invalid enum value for mode",
			Content: `resource "ibm_is_vpngateway_policy_mode" "test" {
  mode = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayPolicyModeRule(),
					Message: "'mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for mode",
			Content: `resource "ibm_is_vpngateway_policy_mode" "test" {
  mode = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayPolicyModeRule(),
					Message: "'mode' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_policy_mode" "test" {
  mode = "policy"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayPolicyModeRule()

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
