package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMReservationCommittedUseRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: term",
			Content: `resource "ibm_is_reservation_committed_use" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationCommittedUseRule(),
					Message: "term attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 52},
					},
				},
			},
		},
		{
			Name: "invalid enum value for expiration_policy",
			Content: `resource "ibm_is_reservation_committed_use" "test" {
  term = "test-term"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationCommittedUseRule(),
					Message: "'expiration_policy' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 52},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for expiration_policy",
			Content: `resource "ibm_is_reservation_committed_use" "test" {
  term = "test-term"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationCommittedUseRule(),
					Message: "'expiration_policy' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 52},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_reservation_committed_use" "test" {
  term = "test-term"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMReservationCommittedUseRule()

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
