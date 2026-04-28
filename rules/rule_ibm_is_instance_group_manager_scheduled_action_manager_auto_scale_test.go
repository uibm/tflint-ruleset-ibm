package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceGroupManagerScheduledActionManagerAutoScaleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_manager_auto_scale" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionManagerAutoScaleRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 85},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_manager_auto_scale" "test" {
  id = "r006-bd21f31f-8a05-4451-836d-ab0347e91fea"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionManagerAutoScaleRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 85},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_manager_auto_scale" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/instance_groups/r006-7b3ac170-01f3-43d6-87ec-f0ed11ed3f60/managers/r006-bd21f31f-8a05-4451-836d-ab0347e91fea"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerScheduledActionManagerAutoScaleRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 85},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_group_manager_scheduled_action_manager_auto_scale" "test" {
  id = "r006-bd21f31f-8a05-4451-836d-ab0347e91fea"
  href = "https://us-south.iaas.cloud.ibm.com/v1/instance_groups/r006-7b3ac170-01f3-43d6-87ec-f0ed11ed3f60/managers/r006-bd21f31f-8a05-4451-836d-ab0347e91fea"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceGroupManagerScheduledActionManagerAutoScaleRule()

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
