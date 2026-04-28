package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMSecurityGroupRuleProtocolAllRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: protocol",
			Content: `resource "ibm_is_security_group_rule_protocol_all" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleProtocolAllRule(),
					Message: "protocol attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "invalid enum value for protocol",
			Content: `resource "ibm_is_security_group_rule_protocol_all" "test" {
  protocol = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleProtocolAllRule(),
					Message: "'protocol' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for protocol",
			Content: `resource "ibm_is_security_group_rule_protocol_all" "test" {
  protocol = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleProtocolAllRule(),
					Message: "'protocol' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_security_group_rule_protocol_all" "test" {
  protocol = "all"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMSecurityGroupRuleProtocolAllRule()

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
