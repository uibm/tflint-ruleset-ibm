package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMImageExportJobRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: storage_bucket",
			Content: `resource "ibm_is_image_export_job" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageExportJobRule(),
					Message: "storage_bucket attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 43},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_image_export_job" "test" {
  storage_bucket = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageExportJobRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 43},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_image_export_job" "test" {
  storage_bucket = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMImageExportJobRule()

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
