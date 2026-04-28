package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpcdnsresolverRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: manual_servers",
			Content: `resource "ibm_is_vpcdnsresolver" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpcdnsresolverRule(),
					Message: "manual_servers attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "invalid enum value for type",
			Content: `resource "ibm_is_vpcdnsresolver" "test" {
  manual_servers = "test-value"
  type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpcdnsresolverRule(),
					Message: "'type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for type",
			Content: `resource "ibm_is_vpcdnsresolver" "test" {
  manual_servers = "test-value"
  type = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpcdnsresolverRule(),
					Message: "'type' value must match pattern",
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
			Content: `resource "ibm_is_vpcdnsresolver" "test" {
  manual_servers = "test-value"
  type = "system"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpcdnsresolverRule()

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
