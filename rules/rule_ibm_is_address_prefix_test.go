package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMAddressPrefixRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: cidr",
			Content: `resource "ibm_is_address_prefix" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMAddressPrefixRule(),
					Message: "cidr attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "string too short for cidr",
			Content: `resource "ibm_is_address_prefix" "test" {
  cidr = "x"
  zone = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMAddressPrefixRule(),
					Message: "'cidr' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for name",
			Content: `resource "ibm_is_address_prefix" "test" {
  cidr = "10.240.0.0/18"
  zone = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMAddressPrefixRule(),
					Message: "'name' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 41},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_address_prefix" "test" {
  cidr = "10.240.0.0/18"
  zone = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMAddressPrefixRule()

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
