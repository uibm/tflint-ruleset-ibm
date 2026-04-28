package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMImageFileRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: href",
			Content: `resource "ibm_is_image_file" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageFileRule(),
					Message: "href attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 37},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_image_file" "test" {
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageFileRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 37},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for href",
			Content: `resource "ibm_is_image_file" "test" {
  href = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMImageFileRule(),
					Message: "'href' value must match pattern",
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
			Content: `resource "ibm_is_image_file" "test" {
  href = "cos://us-south/custom-image-vpc-bucket/customImage-0.vhd"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMImageFileRule()

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
