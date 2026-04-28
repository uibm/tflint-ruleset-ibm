package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerPoolHealthMonitorTypeTcpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: type",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor_type_tcp" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule(),
					Message: "type attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "invalid enum value for type",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor_type_tcp" "test" {
  type = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule(),
					Message: "'type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for type",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor_type_tcp" "test" {
  type = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule(),
					Message: "'type' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_pool_health_monitor_type_tcp" "test" {
  type = "tcp"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule()

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
