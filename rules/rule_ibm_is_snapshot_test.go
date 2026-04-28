package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMSnapshotRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: source_volume",
			Content: `resource "ibm_is_snapshot" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSnapshotRule(),
					Message: "source_volume attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 35},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_snapshot" "test" {
  source_volume = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSnapshotRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 35},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_snapshot" "test" {
  source_volume = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMSnapshotRule()

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
