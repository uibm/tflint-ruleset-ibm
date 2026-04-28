package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionDpdRule checks the configuration of ibm_is_vpngateway_connection_dpd resource
type IBMVpngatewayConnectionDpdRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionDpdRule returns a new rule
func NewIBMVpngatewayConnectionDpdRule() *IBMVpngatewayConnectionDpdRule {
	return &IBMVpngatewayConnectionDpdRule{
		resourceType: "ibm_is_vpngateway_connection_dpd",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionDpdRule) Name() string {
	return "vpngateway_connection_dpd"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionDpdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionDpdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionDpdRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionDpdRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "action"},
			{Name: "interval"},
			{Name: "timeout"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if attr, exists := resource.Body.Attributes["action"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"clear",
					"hold",
					"none",
					"restart",
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
						fmt.Sprintf("'action' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["timeout"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 2 {
					if err := runner.EmitIssue(
						r,
						"'timeout' value must be at least 2",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["timeout"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 86399 {
					if err := runner.EmitIssue(
						r,
						"'timeout' value must not exceed 86399",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["interval"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'interval' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["interval"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 86399 {
					if err := runner.EmitIssue(
						r,
						"'interval' value must not exceed 86399",
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
