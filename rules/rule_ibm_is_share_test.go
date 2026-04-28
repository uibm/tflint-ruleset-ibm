package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMShareRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: profile",
			Content: `resource "ibm_is_share" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareRule(),
					Message: "profile attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "invalid enum value for access_control_mode",
			Content: `resource "ibm_is_share" "test" {
  profile = "test-value"
  size = 1
  zone = "test-value"
  replication_cron_spec = "0 */5 * * *"
  source_share = "test-value"
  origin_share = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareRule(),
					Message: "'access_control_mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "string too short for replication_cron_spec",
			Content: `resource "ibm_is_share" "test" {
  profile = "test-value"
  size = 1
  zone = "test-value"
  replication_cron_spec = "x"
  source_share = "test-value"
  origin_share = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareRule(),
					Message: "'replication_cron_spec' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_share" "test" {
  profile = "test-value"
  size = 1
  zone = "test-value"
  replication_cron_spec = "0 */5 * * *"
  source_share = "test-value"
  origin_share = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 32},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_share" "test" {
  profile = "test-value"
  size = 1
  zone = "test-value"
  replication_cron_spec = "0 */5 * * *"
  source_share = "test-value"
  origin_share = "test-value"
  source_snapshot = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMShareRule()

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
