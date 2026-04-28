package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMRouteRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: destination",
			Content: `resource "ibm_is_route" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteRule(),
					Message: "destination attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "invalid enum value for action",
			Content: `resource "ibm_is_route" "test" {
  destination = "test-value"
  zone = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteRule(),
					Message: "'action' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_route" "test" {
  destination = "test-value"
  zone = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMRouteRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_route" "test" {
  destination = "test-value"
  zone = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMRouteRule()

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
