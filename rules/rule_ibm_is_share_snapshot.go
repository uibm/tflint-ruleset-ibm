package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsShareSnapshotRule checks share snapshot configuration
type IBMIsShareSnapshotRule struct {
	tflint.DefaultRule
}

func NewIBMIsShareSnapshotRule() *IBMIsShareSnapshotRule {
	return &IBMIsShareSnapshotRule{}
}

func (r *IBMIsShareSnapshotRule) Name() string {
	return "ibm_is_share_snapshot"
}

func (r *IBMIsShareSnapshotRule) Enabled() bool {
	return true
}

func (r *IBMIsShareSnapshotRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsShareSnapshotRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsShareSnapshotRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_share_snapshot", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "share"},
			{Name: "tags"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "share"}
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

		// Validate tags if specified
		if attr, exists := resource.Body.Attributes["tags"]; exists {
			var tags []string
			err := runner.EvaluateExpr(attr.Expr, &tags, nil)
			if err != nil {
				return err
			}

			for _, tag := range tags {
				if len(tag) > 128 {
					runner.EmitIssue(
						r,
						"tag length cannot exceed 128 characters",
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	return nil
}
