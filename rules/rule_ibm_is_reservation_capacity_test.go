package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMReservationCapacityRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: total",
			Content: `resource "ibm_is_reservation_capacity" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationCapacityRule(),
					Message: "total attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 47},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_reservation_capacity" "test" {
  total = 1
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMReservationCapacityRule()

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
