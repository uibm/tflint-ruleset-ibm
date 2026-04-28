package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMVpngatewayConnectionIpsecPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_vpngateway_connection_ipsec_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIpsecPolicyRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_vpngateway_connection_ipsec_policy" "test" {
  id = "r006-51eae621-dbbc-4c47-b623-b57a43c19876"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIpsecPolicyRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_vpngateway_connection_ipsec_policy" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/r006-51eae621-dbbc-4c47-b623-b57a43c19876"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMVpngatewayConnectionIpsecPolicyRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 61},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_vpngateway_connection_ipsec_policy" "test" {
  id = "r006-51eae621-dbbc-4c47-b623-b57a43c19876"
  href = "https://us-south.iaas.cloud.ibm.com/v1/ipsec_policies/r006-51eae621-dbbc-4c47-b623-b57a43c19876"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMVpngatewayConnectionIpsecPolicyRule()

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
