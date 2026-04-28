package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceGroupManagerAutoScaleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: manager_type",
			Content: `resource "ibm_is_instance_group_manager_auto_scale" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerAutoScaleRule(),
					Message: "manager_type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "invalid enum value for manager_type",
			Content: `resource "ibm_is_instance_group_manager_auto_scale" "test" {
  manager_type = "invalid-value"
  max_membership_count = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerAutoScaleRule(),
					Message: "'manager_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for manager_type",
			Content: `resource "ibm_is_instance_group_manager_auto_scale" "test" {
  manager_type = "invalid-format"
  max_membership_count = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerAutoScaleRule(),
					Message: "'manager_type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_group_manager_auto_scale" "test" {
  manager_type = "autoscale"
  max_membership_count = 1
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceGroupManagerAutoScaleRule()

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
