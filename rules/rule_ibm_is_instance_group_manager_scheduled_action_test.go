package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceGroupManagerScheduledActionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: group",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionRule(),
					Message: "group attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 66},
					},
				},
			},
		},
		{
			Name: "string too short for cron_spec",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action" "test" {
  group = "test-value"
  manager = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionRule(),
					Message: "'cron_spec' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 66},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for run_at",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action" "test" {
  group = "test-value"
  manager = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionRule(),
					Message: "'run_at' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 66},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action" "test" {
  group = "test-value"
  manager = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceGroupManagerScheduledActionRule()

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
