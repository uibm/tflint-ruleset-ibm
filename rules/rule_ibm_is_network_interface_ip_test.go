package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMNetworkInterfaceIpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_network_interface_ip" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkInterfaceIpRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 47},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_network_interface_ip" "test" {
  id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkInterfaceIpRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 47},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_network_interface_ip" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkInterfaceIpRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 47},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_network_interface_ip" "test" {
  id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
  href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMNetworkInterfaceIpRule()

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
