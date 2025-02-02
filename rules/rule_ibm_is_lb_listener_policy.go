package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsLBListenerPolicyRule checks load balancer listener policy configuration
type IBMIsLBListenerPolicyRule struct {
	tflint.DefaultRule
}

func NewIBMIsLBListenerPolicyRule() *IBMIsLBListenerPolicyRule {
	return &IBMIsLBListenerPolicyRule{}
}

func (r *IBMIsLBListenerPolicyRule) Name() string {
	return "ibm_is_lb_listener_policy"
}

func (r *IBMIsLBListenerPolicyRule) Enabled() bool {
	return true
}

func (r *IBMIsLBListenerPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsLBListenerPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsLBListenerPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_lb_listener_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "lb"},
			{Name: "listener"},
			{Name: "action"},
			{Name: "priority"},
			{Name: "name"},
			{Name: "target_http_status_code"},
			{Name: "target_url"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "rules",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "condition"},
						{Name: "type"},
						{Name: "field"},
						{Name: "value"},
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
		requiredAttrs := []string{"lb", "listener", "action", "priority", "name"}
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
				"forward":  true,
				"redirect": true,
				"reject":   true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be one of: forward, redirect, reject",
					attr.Expr.Range(),
				)
			}

			// Check action-specific requirements
			if action == "redirect" {
				if _, exists := resource.Body.Attributes["target_url"]; !exists {
					runner.EmitIssue(
						r,
						"target_url must be specified when action is redirect",
						attr.Expr.Range(),
					)
				}
				if _, exists := resource.Body.Attributes["target_http_status_code"]; !exists {
					runner.EmitIssue(
						r,
						"target_http_status_code must be specified when action is redirect",
						attr.Expr.Range(),
					)
				}
			}
		}

		// Validate priority
		if attr, exists := resource.Body.Attributes["priority"]; exists {
			var priority int
			err := runner.EvaluateExpr(attr.Expr, &priority, nil)
			if err != nil {
				return err
			}

			if priority < 1 || priority > 10 {
				runner.EmitIssue(
					r,
					"priority must be between 1 and 10",
					attr.Expr.Range(),
				)
			}
		}

		// Validate target_http_status_code if specified
		if attr, exists := resource.Body.Attributes["target_http_status_code"]; exists {
			var statusCode int
			err := runner.EvaluateExpr(attr.Expr, &statusCode, nil)
			if err != nil {
				return err
			}

			validStatusCodes := map[int]bool{
				301: true,
				302: true,
				303: true,
				307: true,
				308: true,
			}

			if !validStatusCodes[statusCode] {
				runner.EmitIssue(
					r,
					"target_http_status_code must be one of: 301, 302, 303, 307, 308",
					attr.Expr.Range(),
				)
			}
		}

		// Validate rules blocks
		for _, block := range resource.Body.Blocks {
			if block.Type == "rules" {
				err := r.validateRule(runner, block)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *IBMIsLBListenerPolicyRule) validateRule(runner tflint.Runner, block *hclext.Block) error {
	// Check required rule attributes
	requiredAttrs := []string{"condition", "type", "field", "value"}
	for _, attr := range requiredAttrs {
		if _, exists := block.Body.Attributes[attr]; !exists {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` attribute must be specified in rules block", attr),
				block.DefRange,
			)
		}
	}

	// Validate condition
	if attr, exists := block.Body.Attributes["condition"]; exists {
		var condition string
		err := runner.EvaluateExpr(attr.Expr, &condition, nil)
		if err != nil {
			return err
		}

		validConditions := map[string]bool{
			"contains":      true,
			"equals":        true,
			"matches_regex": true,
			"not_contains":  true,
			"not_equals":    true,
		}

		if !validConditions[condition] {
			runner.EmitIssue(
				r,
				"condition must be one of: contains, equals, matches_regex, not_contains, not_equals",
				attr.Expr.Range(),
			)
		}
	}

	// Validate type
	if attr, exists := block.Body.Attributes["type"]; exists {
		var ruleType string
		err := runner.EvaluateExpr(attr.Expr, &ruleType, nil)
		if err != nil {
			return err
		}

		validTypes := map[string]bool{
			"header": true,
			"host":   true,
			"path":   true,
			"query":  true,
			"body":   true,
			"cookie": true,
		}

		if !validTypes[ruleType] {
			runner.EmitIssue(
				r,
				"type must be one of: header, host, path, query, body, cookie",
				attr.Expr.Range(),
			)
		}
	}

	return nil
}
