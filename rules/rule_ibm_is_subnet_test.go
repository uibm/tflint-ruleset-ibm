package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMSubnetRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: vpc",
			Content: `resource "ibm_is_subnet" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSubnetRule(),
					Message: "vpc attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "invalid enum value for ip_version",
			Content: `resource "ibm_is_subnet" "test" {
  vpc = "test-value"
  total_ipv4_address_count = 8
  zone = "test-value"
  ipv4_cidr_block = "10.0.0.0/24"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSubnetRule(),
					Message: "'ip_version' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "string too short for ipv4_cidr_block",
			Content: `resource "ibm_is_subnet" "test" {
  vpc = "test-value"
  total_ipv4_address_count = 8
  zone = "test-value"
  ipv4_cidr_block = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSubnetRule(),
					Message: "'ipv4_cidr_block' value must be at least 9 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for ipv4_cidr_block",
			Content: `resource "ibm_is_subnet" "test" {
  vpc = "test-value"
  total_ipv4_address_count = 8
  zone = "test-value"
  ipv4_cidr_block = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSubnetRule(),
					Message: "'ipv4_cidr_block' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 33},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_subnet" "test" {
  vpc = "test-value"
  total_ipv4_address_count = 8
  zone = "test-value"
  ipv4_cidr_block = "10.0.0.0/24"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMSubnetRule()

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
