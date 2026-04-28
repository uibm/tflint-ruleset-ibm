package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMBackupPolicyScopeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: crn",
			Content: `resource "ibm_is_backup_policy_scope" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyScopeRule(),
					Message: "crn attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "string too short for crn",
			Content: `resource "ibm_is_backup_policy_scope" "test" {
  crn = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyScopeRule(),
					Message: "'crn' value must be at least 17 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for crn",
			Content: `resource "ibm_is_backup_policy_scope" "test" {
  crn = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyScopeRule(),
					Message: "'crn' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_backup_policy_scope" "test" {
  crn = "crn:v1:bluemix:public:enterprise::a/aa2432b1fa4d4ace891e9b80fc104e34::enterprise:ebc2b430240943458b9e91e1432cfcce"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMBackupPolicyScopeRule()

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
