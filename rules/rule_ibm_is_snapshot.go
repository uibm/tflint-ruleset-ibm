package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSnapshotRule checks snapshot configuration
type IBMIsSnapshotRule struct {
	tflint.DefaultRule
}

func NewIBMIsSnapshotRule() *IBMIsSnapshotRule {
	return &IBMIsSnapshotRule{}
}

func (r *IBMIsSnapshotRule) Name() string {
	return "ibm_is_snapshot"
}

func (r *IBMIsSnapshotRule) Enabled() bool {
	return true
}

func (r *IBMIsSnapshotRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSnapshotRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSnapshotRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_snapshot", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "source_volume"},
			{Name: "clones"},
			{Name: "resource_group"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "timeouts",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "create"},
						{Name: "delete"},
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
		requiredAttrs := []string{"name", "source_volume"}
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

		// Validate clones if specified
		if attr, exists := resource.Body.Attributes["clones"]; exists {
			var clones []string
			err := runner.EvaluateExpr(attr.Expr, &clones, nil)
			if err != nil {
				return err
			}

			for _, zone := range clones {
				if !isValidZone(zone) {
					runner.EmitIssue(
						r,
						fmt.Sprintf("invalid zone format in clones: %s. Must be in format: region-number (e.g., us-south-1)", zone),
						attr.Expr.Range(),
					)
				}
			}
		}

		// Validate timeouts block if present
		for _, block := range resource.Body.Blocks {
			if block.Type == "timeouts" {
				if err := r.validateTimeouts(runner, block); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *IBMIsSnapshotRule) validateTimeouts(runner tflint.Runner, block *hclext.Block) error {
	timeoutAttrs := []string{"create", "delete"}
	for _, attrName := range timeoutAttrs {
		if attr, exists := block.Body.Attributes[attrName]; exists {
			var timeout string
			err := runner.EvaluateExpr(attr.Expr, &timeout, nil)
			if err != nil {
				return err
			}

			if !isValidTimeout(timeout) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("invalid timeout format for %s: must be a valid duration string (e.g., '15m', '1h')", attrName),
					attr.Expr.Range(),
				)
			}
		}
	}
	return nil
}

func isValidTimeout(timeout string) bool {
	// Add proper timeout validation logic here
	// This is a placeholder - implement proper duration string validation
	return len(timeout) > 0
}
