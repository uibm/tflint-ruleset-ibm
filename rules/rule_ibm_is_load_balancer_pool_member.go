package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerPoolMemberRule checks the configuration of ibm_is_load_balancer_pool_member resource
type IBMLoadBalancerPoolMemberRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerPoolMemberRule returns a new rule
func NewIBMLoadBalancerPoolMemberRule() *IBMLoadBalancerPoolMemberRule {
	return &IBMLoadBalancerPoolMemberRule{
		resourceType: "ibm_is_load_balancer_pool_member",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerPoolMemberRule) Name() string {
	return "load_balancer_pool_member"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerPoolMemberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerPoolMemberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerPoolMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerPoolMemberRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "weight"},
			{Name: "port"},
			{Name: "target"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["port"]; !exists {
			if err := runner.EmitIssue(
				r,
				"port attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["target"]; !exists {
			if err := runner.EmitIssue(
				r,
				"target attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["port"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'port' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["port"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 65535 {
					if err := runner.EmitIssue(
						r,
						"'port' value must not exceed 65535",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["weight"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 0 {
					if err := runner.EmitIssue(
						r,
						"'weight' value must be at least 0",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["weight"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 100 {
					if err := runner.EmitIssue(
						r,
						"'weight' value must not exceed 100",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
