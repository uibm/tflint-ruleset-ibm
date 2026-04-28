package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerPoolHealthMonitorTypeTcpRule checks the configuration of ibm_is_load_balancer_pool_health_monitor_type_tcp resource
type IBMLoadBalancerPoolHealthMonitorTypeTcpRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule returns a new rule
func NewIBMLoadBalancerPoolHealthMonitorTypeTcpRule() *IBMLoadBalancerPoolHealthMonitorTypeTcpRule {
	return &IBMLoadBalancerPoolHealthMonitorTypeTcpRule{
		resourceType: "ibm_is_load_balancer_pool_health_monitor_type_tcp",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerPoolHealthMonitorTypeTcpRule) Name() string {
	return "load_balancer_pool_health_monitor_type_tcp"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerPoolHealthMonitorTypeTcpRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerPoolHealthMonitorTypeTcpRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerPoolHealthMonitorTypeTcpRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerPoolHealthMonitorTypeTcpRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "type"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["type"]; !exists {
			if err := runner.EmitIssue(
				r,
				"type attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"tcp",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'type' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'type' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'type' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for type: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'type' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
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
