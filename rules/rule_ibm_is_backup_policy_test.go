package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMBackupPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: match_resource_type",
			Content: `resource "ibm_is_backup_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyRule(),
					Message: "match_resource_type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 40},
					},
				},
			},
		},
		{
			Name: "invalid enum value for match_resource_type",
			Content: `resource "ibm_is_backup_policy" "test" {
  match_resource_type = "invalid-value"
  match_user_tags = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyRule(),
					Message: "'match_resource_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 40},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for match_resource_type",
			Content: `resource "ibm_is_backup_policy" "test" {
  match_resource_type = "invalid-format"
  match_user_tags = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMBackupPolicyRule(),
					Message: "'match_resource_type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 40},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_backup_policy" "test" {
  match_resource_type = "share"
  match_user_tags = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMBackupPolicyRule()

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
