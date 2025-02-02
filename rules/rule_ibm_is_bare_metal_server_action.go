package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerActionRule checks bare metal server action configuration
type IBMIsBareMetalServerActionRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerActionRule() *IBMIsBareMetalServerActionRule {
	return &IBMIsBareMetalServerActionRule{}
}

func (r *IBMIsBareMetalServerActionRule) Name() string {
	return "ibm_is_bare_metal_server_action"
}

func (r *IBMIsBareMetalServerActionRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerActionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerActionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerActionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_action", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "action"},
			{Name: "stop_type"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "action"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate action
		if attr, exists := resource.Body.Attributes["action"]; exists {
			var action string
			err := runner.EvaluateExpr(attr.Expr, &action, nil)
			if err != nil {
				return err
			}

			validActions := map[string]bool{
				"start":  true,
				"stop":   true,
				"reboot": true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be one of: start, stop, reboot",
					attr.Expr.Range(),
				)
			}

			// Check stop_type when action is stop
			if action == "stop" {
				if stopAttr, exists := resource.Body.Attributes["stop_type"]; exists {
					var stopType string
					err := runner.EvaluateExpr(stopAttr.Expr, &stopType, nil)
					if err != nil {
						return err
					}

					validStopTypes := map[string]bool{
						"hard": true,
						"soft": true,
					}

					if !validStopTypes[stopType] {
						runner.EmitIssue(
							r,
							"stop_type must be either 'hard' or 'soft'",
							stopAttr.Expr.Range(),
						)
					}
				} else {
					runner.EmitIssue(
						r,
						"stop_type must be specified when action is 'stop'",
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	return nil
}
