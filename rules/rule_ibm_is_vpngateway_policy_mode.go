package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayPolicyModeRule checks the configuration of ibm_is_vpngateway_policy_mode resource
type IBMVpngatewayPolicyModeRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayPolicyModeRule returns a new rule
func NewIBMVpngatewayPolicyModeRule() *IBMVpngatewayPolicyModeRule {
	return &IBMVpngatewayPolicyModeRule{
		resourceType: "ibm_is_vpngateway_policy_mode",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayPolicyModeRule) Name() string {
	return "vpngateway_policy_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayPolicyModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayPolicyModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayPolicyModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayPolicyModeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "mode"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["mode"]; !exists {
			if err := runner.EmitIssue(
				r,
				"mode attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"policy",
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
						fmt.Sprintf("'mode' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'mode' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'mode' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for mode: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'mode' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
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
