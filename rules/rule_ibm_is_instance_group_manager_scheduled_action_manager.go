package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceGroupManagerScheduledActionManagerRule checks the configuration of ibm_is_instance_group_manager_scheduled_action_manager resource
type IBMInstanceGroupManagerScheduledActionManagerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceGroupManagerScheduledActionManagerRule returns a new rule
func NewIBMInstanceGroupManagerScheduledActionManagerRule() *IBMInstanceGroupManagerScheduledActionManagerRule {
	return &IBMInstanceGroupManagerScheduledActionManagerRule{
		resourceType: "ibm_is_instance_group_manager_scheduled_action_manager",
	}
}

// Name returns the rule name
func (r *IBMInstanceGroupManagerScheduledActionManagerRule) Name() string {
	return "instance_group_manager_scheduled_action_manager"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceGroupManagerScheduledActionManagerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceGroupManagerScheduledActionManagerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceGroupManagerScheduledActionManagerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceGroupManagerScheduledActionManagerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "id"},
			{Name: "href"},
			{Name: "max_membership_count"},
			{Name: "min_membership_count"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["id"]; !exists {
			if err := runner.EmitIssue(
				r,
				"id attribute must be specified",
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
		if attr, exists := resource.Body.Attributes["max_membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'max_membership_count' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["max_membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'max_membership_count' value must not exceed 1000",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["min_membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'min_membership_count' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["min_membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'min_membership_count' value must not exceed 1000",
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
