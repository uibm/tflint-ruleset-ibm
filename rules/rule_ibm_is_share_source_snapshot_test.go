package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_IBMShareSourceSnapshotRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "missing required attribute: id",
			Content: `resource "ibm_is_share_source_snapshot" "test" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareSourceSnapshotRule(),
					Message: "id attribute must be specified",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "string too short for crn",
			Content: `resource "ibm_is_share_source_snapshot" "test" {
  id = "r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
  crn = "x"
  href = "https://us-south.iaas.cloud.ibm.com/v1/shares/r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/snapshots/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareSourceSnapshotRule(),
					Message: "'crn' value must be at least 17 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "pattern mismatch for id",
			Content: `resource "ibm_is_share_source_snapshot" "test" {
  id = "invalid-format"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share-snapshot:r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
  href = "https://us-south.iaas.cloud.ibm.com/v1/shares/r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/snapshots/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewIBMShareSourceSnapshotRule(),
					Message: "'id' value must match pattern",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 48},
					},
				},
			},
		},
		{
			Name: "valid configuration",
			Content: `resource "ibm_is_share_source_snapshot" "test" {
  id = "r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
  crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::share-snapshot:r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
  href = "https://us-south.iaas.cloud.ibm.com/v1/shares/r006-0fe9e5d8-0a4d-4818-96ec-e99708644a58/snapshots/r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4"
}`,
			Expected: helper.Issues{
			},
		},
	}

	rule := NewIBMShareSourceSnapshotRule()

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
