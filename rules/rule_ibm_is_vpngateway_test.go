package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: subnet",
			Content: `resource "ibm_is_vpngateway" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayRule(),
					Message: "subnet attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 37},
					},
				},
			},
		},
		{
			Name: "invalid enum value for mode",
			Content: `resource "ibm_is_vpngateway" "test" {
  subnet = "test-value"
  mode = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayRule(),
					Message: "'mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 37},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for mode",
			Content: `resource "ibm_is_vpngateway" "test" {
  subnet = "test-value"
  mode = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayRule(),
					Message: "'mode' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 37},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway" "test" {
  subnet = "test-value"
  mode = "policy"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayRule()

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
