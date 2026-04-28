package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMNetworkAclruleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: action",
			Content: `resource "ibm_is_network_aclrule" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleRule(),
					Message: "action attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "invalid enum value for direction",
			Content: `resource "ibm_is_network_aclrule" "test" {
  action = "allow"
  destination = "192.168.3.2/32"
  direction = "invalid-value"
  protocol = "tcp"
  source = "192.168.3.2/32"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleRule(),
					Message: "'direction' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 42},
					},
				},
			},
		},
		{
			Name: "string too short for source",
			Content: `resource "ibm_is_network_aclrule" "test" {
  action = "allow"
  destination = "192.168.3.2/32"
  direction = "inbound"
  protocol = "tcp"
  source = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleRule(),
					Message: "'source' value must be at least 7 characters",
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
			Content: `resource "ibm_is_network_aclrule" "test" {
  action = "invalid-format"
  destination = "192.168.3.2/32"
  direction = "inbound"
  protocol = "tcp"
  source = "192.168.3.2/32"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleRule(),
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
			Content: `resource "ibm_is_network_aclrule" "test" {
  action = "allow"
  destination = "192.168.3.2/32"
  direction = "inbound"
  protocol = "tcp"
  source = "192.168.3.2/32"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMNetworkAclruleRule()

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
