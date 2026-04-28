package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionPolicyModePeerRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: cidrs",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode_peer" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionPolicyModePeerRule(),
					Message: "cidrs attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "string too short for address",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode_peer" "test" {
  cidrs = "test-value"
  address = "x"
  fqdn = "my-service.example.com"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionPolicyModePeerRule(),
					Message: "'address' value must be at least 7 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for fqdn",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode_peer" "test" {
  cidrs = "test-value"
  address = "192.0.2.5"
  fqdn = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionPolicyModePeerRule(),
					Message: "'fqdn' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_connection_policy_mode_peer" "test" {
  cidrs = "test-value"
  address = "192.0.2.5"
  fqdn = "my-service.example.com"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionPolicyModePeerRule()

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
