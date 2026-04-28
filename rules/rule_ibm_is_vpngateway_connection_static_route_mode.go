package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionStaticRouteModeRule checks the configuration of ibm_is_vpngateway_connection_static_route_mode resource
type IBMVpngatewayConnectionStaticRouteModeRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionStaticRouteModeRule returns a new rule
func NewIBMVpngatewayConnectionStaticRouteModeRule() *IBMVpngatewayConnectionStaticRouteModeRule {
	return &IBMVpngatewayConnectionStaticRouteModeRule{
		resourceType: "ibm_is_vpngateway_connection_static_route_mode",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionStaticRouteModeRule) Name() string {
	return "vpngateway_connection_static_route_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionStaticRouteModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionStaticRouteModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionStaticRouteModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionStaticRouteModeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "distribute_traffic"},
			{Name: "local"},
			{Name: "peer"},
			{Name: "routing_protocol"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["peer"]; !exists {
			if err := runner.EmitIssue(
				r,
				"peer attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["routing_protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"none",
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
						fmt.Sprintf("'routing_protocol' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["routing_protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'routing_protocol' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["routing_protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'routing_protocol' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["routing_protocol"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for routing_protocol: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'routing_protocol' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
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
