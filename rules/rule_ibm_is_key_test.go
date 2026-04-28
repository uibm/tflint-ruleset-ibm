package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMKeyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: public_key",
			Content: `resource "ibm_is_key" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMKeyRule(),
					Message: "public_key attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 30},
					},
				},
			},
		},
		{
			Name: "invalid enum value for type",
			Content: `resource "ibm_is_key" "test" {
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMKeyRule(),
					Message: "'type' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 30},
					},
				},
			},
		},
		{
			Name: "string too short for public_key",
			Content: `resource "ibm_is_key" "test" {
  public_key = "x"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMKeyRule(),
					Message: "'public_key' value must be at least 4 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 30},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for public_key",
			Content: `resource "ibm_is_key" "test" {
  public_key = "invalid-format"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMKeyRule(),
					Message: "'public_key' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 30},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_key" "test" {
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMKeyRule()

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
