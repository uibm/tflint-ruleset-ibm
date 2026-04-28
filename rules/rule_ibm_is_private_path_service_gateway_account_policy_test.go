package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMPrivatePathServiceGatewayAccountPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: access_policy",
			Content: `resource "ibm_is_private_path_service_gateway_account_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPrivatePathServiceGatewayAccountPolicyRule(),
					Message: "access_policy attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 70},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_private_path_service_gateway_account_policy" "test" {
  access_policy = "test-value"
  account = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMPrivatePathServiceGatewayAccountPolicyRule()

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
