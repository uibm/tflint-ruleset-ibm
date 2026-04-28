package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMImageRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: file",
			Content: `resource "ibm_is_image" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageRule(),
					Message: "file attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "string too short for obsolescence_at",
			Content: `resource "ibm_is_image" "test" {
  file = "test-value"
  operating_system = "test-value"
  source_volume = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageRule(),
					Message: "'obsolescence_at' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for deprecation_at",
			Content: `resource "ibm_is_image" "test" {
  file = "test-value"
  operating_system = "test-value"
  source_volume = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageRule(),
					Message: "'deprecation_at' value must match pattern",
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
			Content: `resource "ibm_is_image" "test" {
  file = "test-value"
  operating_system = "test-value"
  source_volume = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMImageRule()

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
