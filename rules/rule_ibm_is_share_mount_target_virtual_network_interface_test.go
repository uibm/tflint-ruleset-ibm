package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMShareMountTargetVirtualNetworkInterfaceRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_share_mount_target_virtual_network_interface" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetVirtualNetworkInterfaceRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 71},
					},
				},
			},
		},
		{
			Name: "invalid enum value for protocol_state_filtering_mode",
			Content: `resource "ibm_is_share_mount_target_virtual_network_interface" "test" {
  id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetVirtualNetworkInterfaceRule(),
					Message: "'protocol_state_filtering_mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 71},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_share_mount_target_virtual_network_interface" "test" {
  id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  href = "x"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetVirtualNetworkInterfaceRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 71},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for href",
			Content: `resource "ibm_is_share_mount_target_virtual_network_interface" "test" {
  id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  href = "invalid-format"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareMountTargetVirtualNetworkInterfaceRule(),
					Message: "'href' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 71},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_share_mount_target_virtual_network_interface" "test" {
  id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMShareMountTargetVirtualNetworkInterfaceRule()

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
