package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpnserverAuthenticationByUsernameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: identity_provider",
			Content: `resource "ibm_is_vpnserver_authentication_by_username" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverAuthenticationByUsernameRule(),
					Message: "identity_provider attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpnserver_authentication_by_username" "test" {
  identity_provider = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpnserverAuthenticationByUsernameRule()

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
