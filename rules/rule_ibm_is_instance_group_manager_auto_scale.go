package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceGroupManagerAutoScaleRule checks the configuration of ibm_is_instance_group_manager_auto_scale resource
type IBMInstanceGroupManagerAutoScaleRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceGroupManagerAutoScaleRule returns a new rule
func NewIBMInstanceGroupManagerAutoScaleRule() *IBMInstanceGroupManagerAutoScaleRule {
	return &IBMInstanceGroupManagerAutoScaleRule{
		resourceType: "ibm_is_instance_group_manager_auto_scale",
	}
}

// Name returns the rule name
func (r *IBMInstanceGroupManagerAutoScaleRule) Name() string {
	return "instance_group_manager_auto_scale"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceGroupManagerAutoScaleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceGroupManagerAutoScaleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceGroupManagerAutoScaleRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceGroupManagerAutoScaleRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "min_membership_count"},
			{Name: "aggregation_window"},
			{Name: "cooldown"},
			{Name: "manager_type"},
			{Name: "max_membership_count"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["manager_type"]; !exists {
			if err := runner.EmitIssue(
				r,
				"manager_type attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["max_membership_count"]; !exists {
			if err := runner.EmitIssue(
				r,
				"max_membership_count attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["manager_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"autoscale",
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
						fmt.Sprintf("'manager_type' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["manager_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'manager_type' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["manager_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'manager_type' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["manager_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for manager_type: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'manager_type' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["aggregation_window"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 90 {
					if err := runner.EmitIssue(
						r,
						"'aggregation_window' value must be at least 90",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["aggregation_window"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 600 {
					if err := runner.EmitIssue(
						r,
						"'aggregation_window' value must not exceed 600",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["cooldown"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 120 {
					if err := runner.EmitIssue(
						r,
						"'cooldown' value must be at least 120",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["cooldown"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 3600 {
					if err := runner.EmitIssue(
						r,
						"'cooldown' value must not exceed 3600",
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
