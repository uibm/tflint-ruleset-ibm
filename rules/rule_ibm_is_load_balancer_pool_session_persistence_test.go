package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerPoolSessionPersistenceRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: type",
			Content: `resource "ibm_is_load_balancer_pool_session_persistence" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolSessionPersistenceRule(),
					Message: "type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "invalid enum value for type",
			Content: `resource "ibm_is_load_balancer_pool_session_persistence" "test" {
  type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolSessionPersistenceRule(),
					Message: "'type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for cookie_name",
			Content: `resource "ibm_is_load_balancer_pool_session_persistence" "test" {
  type = "app_cookie"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolSessionPersistenceRule(),
					Message: "'cookie_name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 65},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_pool_session_persistence" "test" {
  type = "app_cookie"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerPoolSessionPersistenceRule()

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
