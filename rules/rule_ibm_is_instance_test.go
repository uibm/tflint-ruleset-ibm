package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMInstanceRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: image",
			Content: `resource "ibm_is_instance" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceRule(),
					Message: "image attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 35},
					},
				},
			},
		},
		{
			Name: "invalid enum value for confidential_compute_mode",
			Content: `resource "ibm_is_instance" "test" {
  image = "test-value"
  zone = "test-value"
  primary_network_attachment = "test-value"
  primary_network_interface = "test-value"
  catalog_offering = "test-value"
  boot_volume_attachment = "test-value"
  source_template = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceRule(),
					Message: "'confidential_compute_mode' value must be one of",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 35},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for confidential_compute_mode",
			Content: `resource "ibm_is_instance" "test" {
  image = "test-value"
  zone = "test-value"
  primary_network_attachment = "test-value"
  primary_network_interface = "test-value"
  catalog_offering = "test-value"
  boot_volume_attachment = "test-value"
  source_template = "test-value"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMInstanceRule(),
					Message: "'confidential_compute_mode' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 35},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_instance" "test" {
  image = "test-value"
  zone = "test-value"
  primary_network_attachment = "test-value"
  primary_network_interface = "test-value"
  catalog_offering = "test-value"
  boot_volume_attachment = "test-value"
  source_template = "test-value"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMInstanceRule()

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
