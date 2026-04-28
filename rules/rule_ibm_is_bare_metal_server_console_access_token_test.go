package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMBareMetalServerConsoleAccessTokenRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: console_type",
			Content: `resource "ibm_is_bare_metal_server_console_access_token" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBareMetalServerConsoleAccessTokenRule(),
					Message: "console_type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "invalid enum value for console_type",
			Content: `resource "ibm_is_bare_metal_server_console_access_token" "test" {
  console_type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBareMetalServerConsoleAccessTokenRule(),
					Message: "'console_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for console_type",
			Content: `resource "ibm_is_bare_metal_server_console_access_token" "test" {
  console_type = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBareMetalServerConsoleAccessTokenRule(),
					Message: "'console_type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_bare_metal_server_console_access_token" "test" {
  console_type = "serial"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMBareMetalServerConsoleAccessTokenRule()

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
