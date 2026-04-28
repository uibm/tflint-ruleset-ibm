package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMSecurityGroupRuleRemoteRule checks the configuration of ibm_is_security_group_rule_remote resource
type IBMSecurityGroupRuleRemoteRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMSecurityGroupRuleRemoteRule returns a new rule
func NewIBMSecurityGroupRuleRemoteRule() *IBMSecurityGroupRuleRemoteRule {
	return &IBMSecurityGroupRuleRemoteRule{
		resourceType: "ibm_is_security_group_rule_remote",
	}
}

// Name returns the rule name
func (r *IBMSecurityGroupRuleRemoteRule) Name() string {
	return "security_group_rule_remote"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMSecurityGroupRuleRemoteRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMSecurityGroupRuleRemoteRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMSecurityGroupRuleRemoteRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMSecurityGroupRuleRemoteRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "address"},
			{Name: "cidr_block"},
			{Name: "id"},
			{Name: "crn"},
			{Name: "href"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["cidr_block"]; !exists {
			if err := runner.EmitIssue(
				r,
				"cidr_block attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["id"]; !exists {
			if err := runner.EmitIssue(
				r,
				"id attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["crn"]; !exists {
			if err := runner.EmitIssue(
				r,
				"crn attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["href"]; !exists {
			if err := runner.EmitIssue(
				r,
				"href attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["address"]; !exists {
			if err := runner.EmitIssue(
				r,
				"address attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["id"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'id' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["id"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 64 {
					if err := runner.EmitIssue(
						r,
						"'id' value must not exceed 64 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 17 {
					if err := runner.EmitIssue(
						r,
						"'crn' value must be at least 17 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 512 {
					if err := runner.EmitIssue(
						r,
						"'crn' value must not exceed 512 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["href"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 10 {
					if err := runner.EmitIssue(
						r,
						"'href' value must be at least 10 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["href"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 8000 {
					if err := runner.EmitIssue(
						r,
						"'href' value must not exceed 8000 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["href"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^http(s)?:\\/\\/([^\\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for href: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'href' value must match pattern: ^http(s)?:\\/\\/([^\\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["id"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[-0-9a-z_]+$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for id: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'id' value must match pattern: ^[-0-9a-z_]+$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\\/]*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for crn: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'crn' value must match pattern: ^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\\/]*$",
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
