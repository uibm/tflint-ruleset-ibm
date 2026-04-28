package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMNetworkAclruleProtocolIcmpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: protocol",
			Content: `resource "ibm_is_network_aclrule_protocol_icmp" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleProtocolIcmpRule(),
					Message: "protocol attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "invalid enum value for protocol",
			Content: `resource "ibm_is_network_aclrule_protocol_icmp" "test" {
  protocol = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleProtocolIcmpRule(),
					Message: "'protocol' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for protocol",
			Content: `resource "ibm_is_network_aclrule_protocol_icmp" "test" {
  protocol = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleProtocolIcmpRule(),
					Message: "'protocol' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 56},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_network_aclrule_protocol_icmp" "test" {
  protocol = "icmp"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMNetworkAclruleProtocolIcmpRule()

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
