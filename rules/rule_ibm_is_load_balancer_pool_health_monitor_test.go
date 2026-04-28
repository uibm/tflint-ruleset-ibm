package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerPoolHealthMonitorRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: delay",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorRule(),
					Message: "delay attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "invalid enum value for type",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor" "test" {
  delay = 2
  max_retries = 1
  timeout = 1
  type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorRule(),
					Message: "'type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for type",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor" "test" {
  delay = 2
  max_retries = 1
  timeout = 1
  type = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorRule(),
					Message: "'type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 60},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor" "test" {
  delay = 2
  max_retries = 1
  timeout = 1
  type = "http"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerPoolHealthMonitorRule()

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
