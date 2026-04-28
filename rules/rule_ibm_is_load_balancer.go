package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerRule checks the configuration of ibm_is_load_balancer resource
type IBMLoadBalancerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerRule returns a new rule
func NewIBMLoadBalancerRule() *IBMLoadBalancerRule {
	return &IBMLoadBalancerRule{
		resourceType: "ibm_is_load_balancer",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerRule) Name() string {
	return "load_balancer"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "pools"},
			{Name: "security_groups"},
			{Name: "dns"},
			{Name: "profile"},
			{Name: "resource_group"},
			{Name: "is_public"},
			{Name: "logging"},
			{Name: "is_private_path"},
			{Name: "listeners"},
			{Name: "route_mode"},
			{Name: "subnets"},
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["subnets"]; !exists {
			if err := runner.EmitIssue(
				r,
				"subnets attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["is_public"]; !exists {
			if err := runner.EmitIssue(
				r,
				"is_public attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'name' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 63 {
					if err := runner.EmitIssue(
						r,
						"'name' value must not exceed 63 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$",
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
