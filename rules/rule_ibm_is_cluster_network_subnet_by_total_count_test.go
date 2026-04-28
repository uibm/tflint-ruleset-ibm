package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMClusterNetworkSubnetByTotalCountRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: total_ipv4_address_count",
			Content: `resource "ibm_is_cluster_network_subnet_by_total_count" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkSubnetByTotalCountRule(),
					Message: "total_ipv4_address_count attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 64},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_cluster_network_subnet_by_total_count" "test" {
  total_ipv4_address_count = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMClusterNetworkSubnetByTotalCountRule()

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
