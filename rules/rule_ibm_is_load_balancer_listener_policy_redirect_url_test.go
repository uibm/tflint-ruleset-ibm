package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMLoadBalancerListenerPolicyRedirectUrlRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: http_status_code",
			Content: `resource "ibm_is_load_balancer_listener_policy_redirect_url" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRedirectUrlRule(),
					Message: "http_status_code attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "string too short for url",
			Content: `resource "ibm_is_load_balancer_listener_policy_redirect_url" "test" {
  http_status_code = "test-value"
  url = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRedirectUrlRule(),
					Message: "'url' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for url",
			Content: `resource "ibm_is_load_balancer_listener_policy_redirect_url" "test" {
  http_status_code = "test-value"
  url = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMLoadBalancerListenerPolicyRedirectUrlRule(),
					Message: "'url' value must match pattern",
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
			Content: `resource "ibm_is_load_balancer_listener_policy_redirect_url" "test" {
  http_status_code = "test-value"
  url = "https://ibm.com:443/new/{path}?{query}"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMLoadBalancerListenerPolicyRedirectUrlRule()

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
