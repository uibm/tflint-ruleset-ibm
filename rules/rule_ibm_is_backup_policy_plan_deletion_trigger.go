package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBackupPolicyPlanDeletionTriggerRule checks the configuration of ibm_is_backup_policy_plan_deletion_trigger resource
type IBMBackupPolicyPlanDeletionTriggerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBackupPolicyPlanDeletionTriggerRule returns a new rule
func NewIBMBackupPolicyPlanDeletionTriggerRule() *IBMBackupPolicyPlanDeletionTriggerRule {
	return &IBMBackupPolicyPlanDeletionTriggerRule{
		resourceType: "ibm_is_backup_policy_plan_deletion_trigger",
	}
}

// Name returns the rule name
func (r *IBMBackupPolicyPlanDeletionTriggerRule) Name() string {
	return "backup_policy_plan_deletion_trigger"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBackupPolicyPlanDeletionTriggerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBackupPolicyPlanDeletionTriggerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBackupPolicyPlanDeletionTriggerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBackupPolicyPlanDeletionTriggerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "delete_after"},
			{Name: "delete_over_count"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if attr, exists := resource.Body.Attributes["delete_over_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'delete_over_count' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["delete_after"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'delete_after' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["delete_after"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'delete_after' value must not exceed 1000",
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
