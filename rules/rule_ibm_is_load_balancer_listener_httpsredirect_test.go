package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerListenerHttpsredirectRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: http_status_code",
			Content: `resource "ibm_is_load_balancer_listener_httpsredirect" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerHttpsredirectRule(),
					Message: "http_status_code attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for uri",
			Content: `resource "ibm_is_load_balancer_listener_httpsredirect" "test" {
  http_status_code = "test-value"
  listener = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerHttpsredirectRule(),
					Message: "'uri' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_load_balancer_listener_httpsredirect" "test" {
  http_status_code = "test-value"
  listener = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerListenerHttpsredirectRule()

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
