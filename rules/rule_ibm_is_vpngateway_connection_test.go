package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: psk",
			Content: `resource "ibm_is_vpngateway_connection" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionRule(),
					Message: "psk attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "invalid enum value for routing_protocol",
			Content: `resource "ibm_is_vpngateway_connection" "test" {
  psk = "lkj14b1oi0alcniejkso"
  peer = "test-value"
  local = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionRule(),
					Message: "'routing_protocol' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "string too short for psk",
			Content: `resource "ibm_is_vpngateway_connection" "test" {
  psk = "x"
  peer = "test-value"
  local = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionRule(),
					Message: "'psk' value must be at least 6 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for psk",
			Content: `resource "ibm_is_vpngateway_connection" "test" {
  psk = "invalid-format"
  peer = "test-value"
  local = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionRule(),
					Message: "'psk' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_connection" "test" {
  psk = "lkj14b1oi0alcniejkso"
  peer = "test-value"
  local = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionRule()

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
