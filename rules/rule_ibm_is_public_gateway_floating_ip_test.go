package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMPublicGatewayFloatingIpRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_public_gateway_floating_ip" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPublicGatewayFloatingIpRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 53},
					},
				},
			},
		},
		{
			Name: "string too short for href",
			Content: `resource "ibm_is_public_gateway_floating_ip" "test" {
  id = "r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::floating-ip:r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  href = "x"
  address = "203.0.113.1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPublicGatewayFloatingIpRule(),
					Message: "'href' value must be at least 10 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 53},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for href",
			Content: `resource "ibm_is_public_gateway_floating_ip" "test" {
  id = "r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::floating-ip:r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  href = "invalid-format"
  address = "203.0.113.1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMPublicGatewayFloatingIpRule(),
					Message: "'href' value must match pattern",
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
			Content: `resource "ibm_is_public_gateway_floating_ip" "test" {
  id = "r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::floating-ip:r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  href = "https://us-south.iaas.cloud.ibm.com/v1/floating_ips/r006-f45e0d90-12a8-4460-8210-290ff2ab75cd"
  address = "203.0.113.1"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMPublicGatewayFloatingIpRule()

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
