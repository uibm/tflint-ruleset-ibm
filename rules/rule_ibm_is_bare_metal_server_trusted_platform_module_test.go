package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMBareMetalServerTrustedPlatformModuleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid enum value for mode",
			Content: `resource "ibm_is_bare_metal_server_trusted_platform_module" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBareMetalServerTrustedPlatformModuleRule(),
					Message: "'mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 68},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for mode",
			Content: `resource "ibm_is_bare_metal_server_trusted_platform_module" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBareMetalServerTrustedPlatformModuleRule(),
					Message: "'mode' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 68},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_bare_metal_server_trusted_platform_module" "test" {
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMBareMetalServerTrustedPlatformModuleRule()

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
