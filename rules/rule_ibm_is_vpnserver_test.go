package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpnserverRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: certificate",
			Content: `resource "ibm_is_vpnserver" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRule(),
					Message: "certificate attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "invalid enum value for protocol",
			Content: `resource "ibm_is_vpnserver" "test" {
  certificate = "test-value"
  client_authentication = "test-value"
  client_ip_pool = "172.16.0.0/16"
  subnets = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRule(),
					Message: "'protocol' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "string too short for client_ip_pool",
			Content: `resource "ibm_is_vpnserver" "test" {
  certificate = "test-value"
  client_authentication = "test-value"
  client_ip_pool = "x"
  subnets = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRule(),
					Message: "'client_ip_pool' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_vpnserver" "test" {
  certificate = "test-value"
  client_authentication = "test-value"
  client_ip_pool = "172.16.0.0/16"
  subnets = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpnserverRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpnserver" "test" {
  certificate = "test-value"
  client_authentication = "test-value"
  client_ip_pool = "172.16.0.0/16"
  subnets = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpnserverRule()

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
