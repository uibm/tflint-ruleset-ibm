package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMIpsecPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: authentication_algorithm",
			Content: `resource "ibm_is_ipsec_policy" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIpsecPolicyRule(),
					Message: "authentication_algorithm attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 39},
					},
				},
			},
		},
		{
			Name: "invalid enum value for pfs",
			Content: `resource "ibm_is_ipsec_policy" "test" {
  authentication_algorithm = "disabled"
  encryption_algorithm = "aes128"
  pfs = "invalid-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIpsecPolicyRule(),
					Message: "'pfs' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 39},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for authentication_algorithm",
			Content: `resource "ibm_is_ipsec_policy" "test" {
  authentication_algorithm = "invalid-format"
  encryption_algorithm = "aes128"
  pfs = "disabled"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMIpsecPolicyRule(),
					Message: "'authentication_algorithm' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 39},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_ipsec_policy" "test" {
  authentication_algorithm = "disabled"
  encryption_algorithm = "aes128"
  pfs = "disabled"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMIpsecPolicyRule()

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
