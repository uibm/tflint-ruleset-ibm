package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceClusterNetworkAttachmentBeforeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_instance_cluster_network_attachment_before" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceClusterNetworkAttachmentBeforeRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_instance_cluster_network_attachment_before" "test" {
  id = "0717-fb880975-db45-4459-8548-64e3995ac213"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceClusterNetworkAttachmentBeforeRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_instance_cluster_network_attachment_before" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceClusterNetworkAttachmentBeforeRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 69},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance_cluster_network_attachment_before" "test" {
  id = "0717-fb880975-db45-4459-8548-64e3995ac213"
  href = "https://us-south.iaas.cloud.ibm.com/v1/instances/0717_e21b7391-2ca2-4ab5-84a8-b92157a633b0/cluster_network_attachments/0717-fb880975-db45-4459-8548-64e3995ac213"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceClusterNetworkAttachmentBeforeRule()

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
