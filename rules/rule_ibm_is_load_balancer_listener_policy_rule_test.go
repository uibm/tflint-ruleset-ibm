package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerListenerPolicyRuleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: condition",
			Content: `resource "ibm_is_load_balancer_listener_policy_rule" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRuleRule(),
					Message: "condition attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "invalid enum value for condition",
			Content: `resource "ibm_is_load_balancer_listener_policy_rule" "test" {
  condition = "invalid-value"
  type = "body"
  value = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRuleRule(),
					Message: "'condition' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for condition",
			Content: `resource "ibm_is_load_balancer_listener_policy_rule" "test" {
  condition = "invalid-format"
  type = "body"
  value = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRuleRule(),
					Message: "'condition' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_listener_policy_rule" "test" {
  condition = "contains"
  type = "body"
  value = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerListenerPolicyRuleRule()

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
