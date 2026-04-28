package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionPolicyModeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: local",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionPolicyModeRule(),
					Message: "local attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode" "test" {
  local = "test-value"
  peer = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionPolicyModeRule()

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
