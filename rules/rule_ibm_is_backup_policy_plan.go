package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBackupPolicyPlanRule checks the configuration of ibm_is_backup_policy_plan resource
type IBMBackupPolicyPlanRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBackupPolicyPlanRule returns a new rule
func NewIBMBackupPolicyPlanRule() *IBMBackupPolicyPlanRule {
	return &IBMBackupPolicyPlanRule{
		resourceType: "ibm_is_backup_policy_plan",
	}
}

// Name returns the rule name
func (r *IBMBackupPolicyPlanRule) Name() string {
	return "backup_policy_plan"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBackupPolicyPlanRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBackupPolicyPlanRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBackupPolicyPlanRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBackupPolicyPlanRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "active"},
			{Name: "attach_user_tags"},
			{Name: "clone_policy"},
			{Name: "copy_user_tags"},
			{Name: "cron_spec"},
			{Name: "deletion_trigger"},
			{Name: "name"},
			{Name: "remote_region_policies"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["cron_spec"]; !exists {
			if err := runner.EmitIssue(
				r,
				"cron_spec attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
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
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'name' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 63 {
					if err := runner.EmitIssue(
						r,
						"'name' value must not exceed 63 characters",
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
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$",
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
