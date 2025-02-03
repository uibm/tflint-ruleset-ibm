package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsFlowLogRule checks flow log configuration
type IBMIsFlowLogRule struct {
	tflint.DefaultRule
}

func NewIBMIsFlowLogRule() *IBMIsFlowLogRule {
	return &IBMIsFlowLogRule{}
}

func (r *IBMIsFlowLogRule) Name() string {
	return "ibm_is_flow_log"
}

func (r *IBMIsFlowLogRule) Enabled() bool {
	return true
}

func (r *IBMIsFlowLogRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsFlowLogRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsFlowLogRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_flow_log", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "target"},
			{Name: "active"},
			{Name: "storage_bucket"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "target", "storage_bucket"}
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

		// Validate target ID format
		if attr, exists := resource.Body.Attributes["target"]; exists {
			var target string
			err := runner.EvaluateExpr(attr.Expr, &target, nil)
			if err != nil {
				return err
			}

			if len(target) == 0 {
				runner.EmitIssue(
					r,
					"target cannot be empty",
					attr.Expr.Range(),
				)
			}
		}

		// Validate storage_bucket format
		if attr, exists := resource.Body.Attributes["storage_bucket"]; exists {
			var bucket string
			err := runner.EvaluateExpr(attr.Expr, &bucket, nil)
			if err != nil {
				return err
			}

			if len(bucket) == 0 {
				runner.EmitIssue(
					r,
					"storage_bucket cannot be empty",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
