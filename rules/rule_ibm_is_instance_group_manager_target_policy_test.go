package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceGroupManagerTargetPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: metric_type",
			Content: `resource "ibm_is_instance_group_manager_target_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerTargetPolicyRule(),
					Message: "metric_type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "invalid enum value for policy_type",
			Content: `resource "ibm_is_instance_group_manager_target_policy" "test" {
  metric_type = "cpu"
  metric_value = 1
  policy_type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerTargetPolicyRule(),
					Message: "'policy_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for metric_type",
			Content: `resource "ibm_is_instance_group_manager_target_policy" "test" {
  metric_type = "invalid-format"
  metric_value = 1
  policy_type = "target"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceGroupManagerTargetPolicyRule(),
					Message: "'metric_type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_group_manager_target_policy" "test" {
  metric_type = "cpu"
  metric_value = 1
  policy_type = "target"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceGroupManagerTargetPolicyRule()

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
