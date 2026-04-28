package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerPoolFailsafePolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid enum value for action",
			Content: `resource "ibm_is_load_balancer_pool_failsafe_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolFailsafePolicyRule(),
					Message: "'action' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for action",
			Content: `resource "ibm_is_load_balancer_pool_failsafe_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolFailsafePolicyRule(),
					Message: "'action' value must match pattern",
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
			Content: `resource "ibm_is_load_balancer_pool_failsafe_policy" "test" {
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerPoolFailsafePolicyRule()

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
