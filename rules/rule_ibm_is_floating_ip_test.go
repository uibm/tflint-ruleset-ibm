package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMFloatingIpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: zone",
			Content: `resource "ibm_is_floating_ip" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMFloatingIpRule(),
					Message: "zone attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 38},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_floating_ip" "test" {
  zone = "test-value"
  target = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMFloatingIpRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 38},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_floating_ip" "test" {
  zone = "test-value"
  target = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMFloatingIpRule()

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
