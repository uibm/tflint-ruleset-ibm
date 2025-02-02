package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceActionRule checks instance action configuration
type IBMIsInstanceActionRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceActionRule() *IBMIsInstanceActionRule {
	return &IBMIsInstanceActionRule{}
}

func (r *IBMIsInstanceActionRule) Name() string {
	return "ibm_is_instance_action"
}

func (r *IBMIsInstanceActionRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceActionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceActionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceActionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_action", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "action"},
			{Name: "instance"},
			{Name: "force"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"action", "instance"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate action type
		if attr, exists := resource.Body.Attributes["action"]; exists {
			var action string
			err := runner.EvaluateExpr(attr.Expr, &action, nil)
			if err != nil {
				return err
			}

			validActions := map[string]bool{
				"start":   true,
				"stop":    true,
				"restart": true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be one of: start, stop, restart",
					attr.Expr.Range(),
				)
			}
		}

		// Validate force flag if specified
		if attr, exists := resource.Body.Attributes["force"]; exists {
			var force bool
			err := runner.EvaluateExpr(attr.Expr, &force, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
