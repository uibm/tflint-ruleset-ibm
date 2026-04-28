package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMClusterNetworkSubnetByIpv4CidrblockRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: ipv4_cidr_block",
			Content: `resource "ibm_is_cluster_network_subnet_by_ipv4_cidrblock" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkSubnetByIpv4CidrblockRule(),
					Message: "ipv4_cidr_block attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 67},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_cluster_network_subnet_by_ipv4_cidrblock" "test" {
  ipv4_cidr_block = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMClusterNetworkSubnetByIpv4CidrblockRule()

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
