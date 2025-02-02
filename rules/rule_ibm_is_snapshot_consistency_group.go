package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSnapshotConsistencyGroupRule checks snapshot consistency group configuration
type IBMIsSnapshotConsistencyGroupRule struct {
	tflint.DefaultRule
}

func NewIBMIsSnapshotConsistencyGroupRule() *IBMIsSnapshotConsistencyGroupRule {
	return &IBMIsSnapshotConsistencyGroupRule{}
}

func (r *IBMIsSnapshotConsistencyGroupRule) Name() string {
	return "ibm_is_snapshot_consistency_group"
}

func (r *IBMIsSnapshotConsistencyGroupRule) Enabled() bool {
	return true
}

func (r *IBMIsSnapshotConsistencyGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSnapshotConsistencyGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSnapshotConsistencyGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_snapshot_consistency_group", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "delete_snapshots_on_delete"},
			{Name: "resource_group"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "snapshots",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
						{Name: "source_volume"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}

		// Check snapshots blocks
		hasSnapshots := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "snapshots" {
				hasSnapshots = true

				// Check required attributes in snapshots block
				requiredSnapshotAttrs := []string{"name", "source_volume"}
				for _, attr := range requiredSnapshotAttrs {
					if _, exists := block.Body.Attributes[attr]; !exists {
						runner.EmitIssue(
							r,
							fmt.Sprintf("`%s` attribute must be specified in snapshots block", attr),
							block.DefRange,
						)
					}
				}
			}
		}

		if !hasSnapshots {
			runner.EmitIssue(
				r,
				"at least one snapshots block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
