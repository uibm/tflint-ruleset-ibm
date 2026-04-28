package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBackupPolicyPlanRemoteRegionPolicyRule checks the configuration of ibm_is_backup_policy_plan_remote_region_policy resource
type IBMBackupPolicyPlanRemoteRegionPolicyRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBackupPolicyPlanRemoteRegionPolicyRule returns a new rule
func NewIBMBackupPolicyPlanRemoteRegionPolicyRule() *IBMBackupPolicyPlanRemoteRegionPolicyRule {
	return &IBMBackupPolicyPlanRemoteRegionPolicyRule{
		resourceType: "ibm_is_backup_policy_plan_remote_region_policy",
	}
}

// Name returns the rule name
func (r *IBMBackupPolicyPlanRemoteRegionPolicyRule) Name() string {
	return "backup_policy_plan_remote_region_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBackupPolicyPlanRemoteRegionPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBackupPolicyPlanRemoteRegionPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBackupPolicyPlanRemoteRegionPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBackupPolicyPlanRemoteRegionPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "delete_over_count"},
			{Name: "encryption_key"},
			{Name: "region"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["region"]; !exists {
			if err := runner.EmitIssue(
				r,
				"region attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
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
		if attr, exists := resource.Body.Attributes["delete_over_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'delete_over_count' value must not exceed 1000",
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
