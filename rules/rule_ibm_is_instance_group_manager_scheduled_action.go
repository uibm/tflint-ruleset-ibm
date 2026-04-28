package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceGroupManagerScheduledActionRule checks the configuration of ibm_is_instance_group_manager_scheduled_action resource
type IBMInstanceGroupManagerScheduledActionRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceGroupManagerScheduledActionRule returns a new rule
func NewIBMInstanceGroupManagerScheduledActionRule() *IBMInstanceGroupManagerScheduledActionRule {
	return &IBMInstanceGroupManagerScheduledActionRule{
		resourceType: "ibm_is_instance_group_manager_scheduled_action",
	}
}

// Name returns the rule name
func (r *IBMInstanceGroupManagerScheduledActionRule) Name() string {
	return "instance_group_manager_scheduled_action"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceGroupManagerScheduledActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceGroupManagerScheduledActionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceGroupManagerScheduledActionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceGroupManagerScheduledActionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "run_at"},
			{Name: "group"},
			{Name: "manager"},
			{Name: "cron_spec"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["group"]; !exists {
			if err := runner.EmitIssue(
				r,
				"group attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["manager"]; !exists {
			if err := runner.EmitIssue(
				r,
				"manager attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["run_at"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 10 {
					if err := runner.EmitIssue(
						r,
						"'run_at' value must be at least 10 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["run_at"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 64 {
					if err := runner.EmitIssue(
						r,
						"'run_at' value must not exceed 64 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["cron_spec"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 9 {
					if err := runner.EmitIssue(
						r,
						"'cron_spec' value must be at least 9 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["cron_spec"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 63 {
					if err := runner.EmitIssue(
						r,
						"'cron_spec' value must not exceed 63 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["cron_spec"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^(((\\d+,)+\\d+|([\\d\\*]+(\\/|-)\\d+)|\\d+|\\*) ?){5}$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for cron_spec: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'cron_spec' value must match pattern: ^(((\\d+,)+\\d+|([\\d\\*]+(\\/|-)\\d+)|\\d+|\\*) ?){5}$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["run_at"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^((?:(\\d{4}-\\d{2}-\\d{2})T(\\d{2}:\\d{2}:\\d{2}(?:\\.\\d+)?))(Z|[\\+-]\\d{2}:\\d{2})?)$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for run_at: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'run_at' value must match pattern: ^((?:(\\d{4}-\\d{2}-\\d{2})T(\\d{2}:\\d{2}:\\d{2}(?:\\.\\d+)?))(Z|[\\+-]\\d{2}:\\d{2})?)$",
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
