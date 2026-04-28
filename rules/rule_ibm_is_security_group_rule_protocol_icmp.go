package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMSecurityGroupRuleProtocolIcmpRule checks the configuration of ibm_is_security_group_rule_protocol_icmp resource
type IBMSecurityGroupRuleProtocolIcmpRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMSecurityGroupRuleProtocolIcmpRule returns a new rule
func NewIBMSecurityGroupRuleProtocolIcmpRule() *IBMSecurityGroupRuleProtocolIcmpRule {
	return &IBMSecurityGroupRuleProtocolIcmpRule{
		resourceType: "ibm_is_security_group_rule_protocol_icmp",
	}
}

// Name returns the rule name
func (r *IBMSecurityGroupRuleProtocolIcmpRule) Name() string {
	return "security_group_rule_protocol_icmp"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMSecurityGroupRuleProtocolIcmpRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMSecurityGroupRuleProtocolIcmpRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMSecurityGroupRuleProtocolIcmpRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMSecurityGroupRuleProtocolIcmpRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "type"},
			{Name: "code"},
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
					"icmp",
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
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 0 {
					if err := runner.EmitIssue(
						r,
						"'type' value must be at least 0",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 254 {
					if err := runner.EmitIssue(
						r,
						"'type' value must not exceed 254",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["code"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 0 {
					if err := runner.EmitIssue(
						r,
						"'code' value must be at least 0",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["code"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 255 {
					if err := runner.EmitIssue(
						r,
						"'code' value must not exceed 255",
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
