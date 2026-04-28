package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerListenerRule checks the configuration of ibm_is_load_balancer_listener resource
type IBMLoadBalancerListenerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerListenerRule returns a new rule
func NewIBMLoadBalancerListenerRule() *IBMLoadBalancerListenerRule {
	return &IBMLoadBalancerListenerRule{
		resourceType: "ibm_is_load_balancer_listener",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerListenerRule) Name() string {
	return "load_balancer_listener"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerListenerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerListenerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerListenerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerListenerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "policies"},
			{Name: "accept_proxy_protocol"},
			{Name: "connection_limit"},
			{Name: "idle_connection_timeout"},
			{Name: "certificate_instance"},
			{Name: "default_pool"},
			{Name: "port"},
			{Name: "port_max"},
			{Name: "https_redirect"},
			{Name: "port_min"},
			{Name: "protocol"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["protocol"]; !exists {
			if err := runner.EmitIssue(
				r,
				"protocol attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"http",
					"https",
					"tcp",
					"udp",
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
						fmt.Sprintf("'protocol' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'protocol' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'protocol' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for protocol: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'protocol' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["port_max"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'port_max' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["port_max"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 65535 {
					if err := runner.EmitIssue(
						r,
						"'port_max' value must not exceed 65535",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["port_min"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'port_min' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["port_min"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 65535 {
					if err := runner.EmitIssue(
						r,
						"'port_min' value must not exceed 65535",
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
		if attr, exists := resource.Body.Attributes["connection_limit"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'connection_limit' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["connection_limit"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 15000 {
					if err := runner.EmitIssue(
						r,
						"'connection_limit' value must not exceed 15000",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["idle_connection_timeout"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 50 {
					if err := runner.EmitIssue(
						r,
						"'idle_connection_timeout' value must be at least 50",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["idle_connection_timeout"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 7200 {
					if err := runner.EmitIssue(
						r,
						"'idle_connection_timeout' value must not exceed 7200",
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
