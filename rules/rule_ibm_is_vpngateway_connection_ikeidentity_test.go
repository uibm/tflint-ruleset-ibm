package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionIkeidentityRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: type",
			Content: `resource "ibm_is_vpngateway_connection_ikeidentity" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkeidentityRule(),
					Message: "type attribute must be specified",
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
			Content: `resource "ibm_is_vpngateway_connection_ikeidentity" "test" {
  type = "invalid-value"
  value = "qQ+/YEApnl1ZtEgIrfprzb065307thTkzlnLqL5ICpesdbBN03dyCQ=="
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkeidentityRule(),
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
			Name: "string too short for value",
			Content: `resource "ibm_is_vpngateway_connection_ikeidentity" "test" {
  type = "fqdn"
  value = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkeidentityRule(),
					Message: "'value' value must be at least 4 characters",
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
			Content: `resource "ibm_is_vpngateway_connection_ikeidentity" "test" {
  type = "invalid-format"
  value = "qQ+/YEApnl1ZtEgIrfprzb065307thTkzlnLqL5ICpesdbBN03dyCQ=="
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIkeidentityRule(),
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
			Content: `resource "ibm_is_vpngateway_connection_ikeidentity" "test" {
  type = "fqdn"
  value = "qQ+/YEApnl1ZtEgIrfprzb065307thTkzlnLqL5ICpesdbBN03dyCQ=="
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionIkeidentityRule()

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
