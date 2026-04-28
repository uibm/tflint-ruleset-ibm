package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMClusterNetworkInterfacePrimaryIpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_cluster_network_interface_primary_ip" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfacePrimaryIpRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "string too short for address",
			Content: `resource "ibm_is_cluster_network_interface_primary_ip" "test" {
  id = "0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
  href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfacePrimaryIpRule(),
					Message: "'address' value must be at least 7 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_cluster_network_interface_primary_ip" "test" {
  id = "0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
  href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfacePrimaryIpRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 63},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_cluster_network_interface_primary_ip" "test" {
  id = "0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
  href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930/reserved_ips/0717-d4d6489a-3bf5-4104-a33a-3572faf2d117"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMClusterNetworkInterfacePrimaryIpRule()

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
