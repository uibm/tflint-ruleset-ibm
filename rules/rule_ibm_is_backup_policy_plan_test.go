package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMBackupPolicyPlanRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: cron_spec",
			Content: `resource "ibm_is_backup_policy_plan" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyPlanRule(),
					Message: "cron_spec attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "string too short for cron_spec",
			Content: `resource "ibm_is_backup_policy_plan" "test" {
  cron_spec = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyPlanRule(),
					Message: "'cron_spec' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_backup_policy_plan" "test" {
  cron_spec = "30 */2 * * 1-5"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyPlanRule(),
					Message: "'name' value must match pattern",
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
			Content: `resource "ibm_is_backup_policy_plan" "test" {
  cron_spec = "30 */2 * * 1-5"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMBackupPolicyPlanRule()

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
