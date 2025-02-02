package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsLBPoolMemberRule checks load balancer pool member configuration
type IBMIsLBPoolMemberRule struct {
	tflint.DefaultRule
}

func NewIBMIsLBPoolMemberRule() *IBMIsLBPoolMemberRule {
	return &IBMIsLBPoolMemberRule{}
}

func (r *IBMIsLBPoolMemberRule) Name() string {
	return "ibm_is_lb_pool_member"
}

func (r *IBMIsLBPoolMemberRule) Enabled() bool {
	return true
}

func (r *IBMIsLBPoolMemberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsLBPoolMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsLBPoolMemberRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_lb_pool_member", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "lb"},
			{Name: "pool"},
			{Name: "port"},
			{Name: "target_address"},
			{Name: "target_id"},
			{Name: "weight"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"lb", "pool", "port"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check if either target_address or target_id is specified
		hasTarget := false
		if _, exists := resource.Body.Attributes["target_address"]; exists {
			hasTarget = true
		}
		if _, exists := resource.Body.Attributes["target_id"]; exists {
			if hasTarget {
				runner.EmitIssue(
					r,
					"only one of target_address or target_id can be specified",
					resource.DefRange,
				)
			}
			hasTarget = true
		}

		if !hasTarget {
			runner.EmitIssue(
				r,
				"either target_address or target_id must be specified",
				resource.DefRange,
			)
		}

		// Validate port range
		if attr, exists := resource.Body.Attributes["port"]; exists {
			var port int
			err := runner.EvaluateExpr(attr.Expr, &port, nil)
			if err != nil {
				return err
			}

			if port < 1 || port > 65535 {
				runner.EmitIssue(
					r,
					"port must be between 1 and 65535",
					attr.Expr.Range(),
				)
			}
		}

		// Validate weight if specified
		if attr, exists := resource.Body.Attributes["weight"]; exists {
			var weight int
			err := runner.EvaluateExpr(attr.Expr, &weight, nil)
			if err != nil {
				return err
			}

			if weight < 0 || weight > 100 {
				runner.EmitIssue(
					r,
					"weight must be between 0 and 100",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
