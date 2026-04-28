package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMIkepolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: authentication_algorithm",
			Content: `resource "ibm_is_ikepolicy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIkepolicyRule(),
					Message: "authentication_algorithm attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "invalid enum value for authentication_algorithm",
			Content: `resource "ibm_is_ikepolicy" "test" {
  authentication_algorithm = "invalid-value"
  dh_group = 1
  encryption_algorithm = "aes128"
  ike_version = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIkepolicyRule(),
					Message: "'authentication_algorithm' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for encryption_algorithm",
			Content: `resource "ibm_is_ikepolicy" "test" {
  authentication_algorithm = "sha256"
  dh_group = 1
  encryption_algorithm = "invalid-format"
  ike_version = 1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIkepolicyRule(),
					Message: "'encryption_algorithm' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 36},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_ikepolicy" "test" {
  authentication_algorithm = "sha256"
  dh_group = 1
  encryption_algorithm = "aes128"
  ike_version = 1
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMIkepolicyRule()

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
