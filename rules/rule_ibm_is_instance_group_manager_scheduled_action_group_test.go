package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceGroupManagerScheduledActionGroupRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: membership_count",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_group" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionGroupRule(),
					Message: "membership_count attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 72},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_group" "test" {
  membership_count = 1
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceGroupManagerScheduledActionGroupRule()

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
