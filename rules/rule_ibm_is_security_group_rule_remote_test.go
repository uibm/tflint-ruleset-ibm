package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMSecurityGroupRuleRemoteRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: address",
			Content: `resource "ibm_is_security_group_rule_remote" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRemoteRule(),
					Message: "address attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 53},
					},
				},
			},
		},
		{
			Name: "string too short for crn",
			Content: `resource "ibm_is_security_group_rule_remote" "test" {
  address = "test-value"
  cidr_block = "test-value"
  id = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
  crn = "x"
  href = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRemoteRule(),
					Message: "'crn' value must be at least 17 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 53},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for crn",
			Content: `resource "ibm_is_security_group_rule_remote" "test" {
  address = "test-value"
  cidr_block = "test-value"
  id = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
  crn = "invalid-format"
  href = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMSecurityGroupRuleRemoteRule(),
					Message: "'crn' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 53},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_security_group_rule_remote" "test" {
  address = "test-value"
  cidr_block = "test-value"
  id = "r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
  crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::security-group:r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
  href = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-be5df5ca-12a0-494b-907e-aa6ec2bfa271"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMSecurityGroupRuleRemoteRule()

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
