package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerPoolRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: algorithm",
			Content: `resource "ibm_is_load_balancer_pool" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolRule(),
					Message: "algorithm attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "invalid enum value for protocol",
			Content: `resource "ibm_is_load_balancer_pool" "test" {
  algorithm = "least_connections"
  health_monitor = "test-value"
  protocol = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolRule(),
					Message: "'protocol' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for proxy_protocol",
			Content: `resource "ibm_is_load_balancer_pool" "test" {
  algorithm = "least_connections"
  health_monitor = "test-value"
  protocol = "http"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerPoolRule(),
					Message: "'proxy_protocol' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 45},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_pool" "test" {
  algorithm = "least_connections"
  health_monitor = "test-value"
  protocol = "http"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerPoolRule()

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
