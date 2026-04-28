package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMPlacementGroupRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: strategy",
			Content: `resource "ibm_is_placement_group" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPlacementGroupRule(),
					Message: "strategy attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "invalid enum value for strategy",
			Content: `resource "ibm_is_placement_group" "test" {
  strategy = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPlacementGroupRule(),
					Message: "'strategy' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_placement_group" "test" {
  strategy = "host_spread"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPlacementGroupRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_placement_group" "test" {
  strategy = "host_spread"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMPlacementGroupRule()

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
