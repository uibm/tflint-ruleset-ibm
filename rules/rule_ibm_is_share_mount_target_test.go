package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMShareMountTargetRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: virtual_network_interface",
			Content: `resource "ibm_is_share_mount_target" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetRule(),
					Message: "virtual_network_interface attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "invalid enum value for transit_encryption",
			Content: `resource "ibm_is_share_mount_target" "test" {
  virtual_network_interface = "test-value"
  vpc = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetRule(),
					Message: "'transit_encryption' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for transit_encryption",
			Content: `resource "ibm_is_share_mount_target" "test" {
  virtual_network_interface = "test-value"
  vpc = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetRule(),
					Message: "'transit_encryption' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_share_mount_target" "test" {
  virtual_network_interface = "test-value"
  vpc = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMShareMountTargetRule()

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
