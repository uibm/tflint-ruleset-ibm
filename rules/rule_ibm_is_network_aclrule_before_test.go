package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMNetworkAclruleBeforeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_network_aclrule_before" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleBeforeRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_network_aclrule_before" "test" {
  id = "r006-8daca77a-4980-4d33-8f3e-7038797be8f9"
  href = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleBeforeRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_network_aclrule_before" "test" {
  id = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-3217cb8b-5368-452a-9399-a84f14fb539d/rules/r006-8daca77a-4980-4d33-8f3e-7038797be8f9"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMNetworkAclruleBeforeRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 49},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_network_aclrule_before" "test" {
  id = "r006-8daca77a-4980-4d33-8f3e-7038797be8f9"
  href = "https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-3217cb8b-5368-452a-9399-a84f14fb539d/rules/r006-8daca77a-4980-4d33-8f3e-7038797be8f9"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMNetworkAclruleBeforeRule()

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
