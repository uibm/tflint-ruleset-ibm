package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBackupPolicyPlanClonePolicyRule checks the configuration of ibm_is_backup_policy_plan_clone_policy resource
type IBMBackupPolicyPlanClonePolicyRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBackupPolicyPlanClonePolicyRule returns a new rule
func NewIBMBackupPolicyPlanClonePolicyRule() *IBMBackupPolicyPlanClonePolicyRule {
	return &IBMBackupPolicyPlanClonePolicyRule{
		resourceType: "ibm_is_backup_policy_plan_clone_policy",
	}
}

// Name returns the rule name
func (r *IBMBackupPolicyPlanClonePolicyRule) Name() string {
	return "backup_policy_plan_clone_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBackupPolicyPlanClonePolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBackupPolicyPlanClonePolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBackupPolicyPlanClonePolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBackupPolicyPlanClonePolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "max_snapshots"},
			{Name: "zones"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["zones"]; !exists {
			if err := runner.EmitIssue(
				r,
				"zones attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["max_snapshots"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'max_snapshots' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["max_snapshots"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'max_snapshots' value must not exceed 1000",
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
