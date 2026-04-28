package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpnserverRouteRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: destination",
			Content: `resource "ibm_is_vpnserver_route" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRouteRule(),
					Message: "destination attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "invalid enum value for action",
			Content: `resource "ibm_is_vpnserver_route" "test" {
  destination = "172.16.0.0/16"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRouteRule(),
					Message: "'action' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "string too short for destination",
			Content: `resource "ibm_is_vpnserver_route" "test" {
  destination = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRouteRule(),
					Message: "'destination' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for action",
			Content: `resource "ibm_is_vpnserver_route" "test" {
  destination = "172.16.0.0/16"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRouteRule(),
					Message: "'action' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpnserver_route" "test" {
  destination = "172.16.0.0/16"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpnserverRouteRule()

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
