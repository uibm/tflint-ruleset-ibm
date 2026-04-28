package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMClusterNetworkInterfaceSubnetRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_cluster_network_interface_subnet" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfaceSubnetRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_cluster_network_interface_subnet" "test" {
  id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfaceSubnetRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for href",
			Content: `resource "ibm_is_cluster_network_interface_subnet" "test" {
  id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
  href = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMClusterNetworkInterfaceSubnetRule(),
					Message: "'href' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 59},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_cluster_network_interface_subnet" "test" {
  id = "0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
  href = "https://us-south.iaas.cloud.ibm.com/v1/cluster_networks/0717-da0df18c-7598-4633-a648-fdaac28a5573/subnets/0717-7931845c-65c4-4b0a-80cd-7d9c1a6d7930"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMClusterNetworkInterfaceSubnetRule()

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
