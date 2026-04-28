package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMReservationProfileRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: name",
			Content: `resource "ibm_is_reservation_profile" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationProfileRule(),
					Message: "name attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "invalid enum value for resource_type",
			Content: `resource "ibm_is_reservation_profile" "test" {
  name = "bx2-4x16"
  resource_type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationProfileRule(),
					Message: "'resource_type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_reservation_profile" "test" {
  name = "invalid-format"
  resource_type = "bare_metal_server_profile"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMReservationProfileRule(),
					Message: "'name' value must match pattern",
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
			Content: `resource "ibm_is_reservation_profile" "test" {
  name = "bx2-4x16"
  resource_type = "bare_metal_server_profile"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMReservationProfileRule()

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
