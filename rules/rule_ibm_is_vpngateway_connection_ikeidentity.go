package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionIkeidentityRule checks the configuration of ibm_is_vpngateway_connection_ikeidentity resource
type IBMVpngatewayConnectionIkeidentityRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionIkeidentityRule returns a new rule
func NewIBMVpngatewayConnectionIkeidentityRule() *IBMVpngatewayConnectionIkeidentityRule {
	return &IBMVpngatewayConnectionIkeidentityRule{
		resourceType: "ibm_is_vpngateway_connection_ikeidentity",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionIkeidentityRule) Name() string {
	return "vpngateway_connection_ikeidentity"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionIkeidentityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionIkeidentityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionIkeidentityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionIkeidentityRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "type"},
			{Name: "value"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["value"]; !exists {
			if err := runner.EmitIssue(
				r,
				"value attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
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
					"fqdn",
					"hostname",
					"ipv4_address",
					"key_id",
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
		if attr, exists := resource.Body.Attributes["value"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 4 {
					if err := runner.EmitIssue(
						r,
						"'value' value must be at least 4 characters",
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
		if attr, exists := resource.Body.Attributes["value"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for value: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'value' value must match pattern: ^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$",
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
