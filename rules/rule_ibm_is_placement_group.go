package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPlacementGroupRule checks placement group configuration
type IBMIsPlacementGroupRule struct {
	tflint.DefaultRule
}

func NewIBMIsPlacementGroupRule() *IBMIsPlacementGroupRule {
	return &IBMIsPlacementGroupRule{}
}

func (r *IBMIsPlacementGroupRule) Name() string {
	return "ibm_is_placement_group"
}

func (r *IBMIsPlacementGroupRule) Enabled() bool {
	return true
}

func (r *IBMIsPlacementGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPlacementGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPlacementGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_placement_group", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "strategy"},
			{Name: "name"},
			{Name: "resource_group"},
			{Name: "access_tags"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"strategy", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate strategy
		if attr, exists := resource.Body.Attributes["strategy"]; exists {
			var strategy string
			err := runner.EvaluateExpr(attr.Expr, &strategy, nil)
			if err != nil {
				return err
			}

			validStrategies := map[string]bool{
				"host_spread":  true,
				"power_spread": true,
				"rack_spread":  true,
			}

			if !validStrategies[strategy] {
				runner.EmitIssue(
					r,
					"strategy must be one of: host_spread, power_spread, rack_spread",
					attr.Expr.Range(),
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

		// Additional validations could be added for access_tags if needed
	}

	return nil
}
