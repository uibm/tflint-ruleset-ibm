package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVolumeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: profile",
			Content: `resource "ibm_is_volume" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVolumeRule(),
					Message: "profile attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_volume" "test" {
  profile = "test-value"
  zone = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVolumeRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_volume" "test" {
  profile = "test-value"
  zone = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVolumeRule()

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
