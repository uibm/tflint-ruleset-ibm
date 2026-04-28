package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMSecurityGroupRuleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: direction",
			Content: `resource "ibm_is_security_group_rule" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRule(),
					Message: "direction attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "invalid enum value for direction",
			Content: `resource "ibm_is_security_group_rule" "test" {
  direction = "invalid-value"
  protocol = "tcp"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRule(),
					Message: "'direction' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 46},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for direction",
			Content: `resource "ibm_is_security_group_rule" "test" {
  direction = "invalid-format"
  protocol = "tcp"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRule(),
					Message: "'direction' value must match pattern",
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
			Content: `resource "ibm_is_security_group_rule" "test" {
  direction = "inbound"
  protocol = "tcp"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMSecurityGroupRuleRule()

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
