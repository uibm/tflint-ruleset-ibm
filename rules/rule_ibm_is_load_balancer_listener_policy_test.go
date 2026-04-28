package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerListenerPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: action",
			Content: `resource "ibm_is_load_balancer_listener_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRule(),
					Message: "action attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "invalid enum value for action",
			Content: `resource "ibm_is_load_balancer_listener_policy" "test" {
  action = "invalid-value"
  priority = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRule(),
					Message: "'action' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for action",
			Content: `resource "ibm_is_load_balancer_listener_policy" "test" {
  action = "invalid-format"
  priority = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRule(),
					Message: "'action' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_listener_policy" "test" {
  action = "forward_to_listener"
  priority = 1
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerListenerPolicyRule()

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
