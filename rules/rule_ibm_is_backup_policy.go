package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

type IBMIsBackupPolicyRule struct {
	tflint.DefaultRule
}

func NewIBMIsBackupPolicyRule() *IBMIsBackupPolicyRule {
	return &IBMIsBackupPolicyRule{}
}

func (r *IBMIsBackupPolicyRule) Name() string {
	return "ibm_is_backup_policy"
}

func (r *IBMIsBackupPolicyRule) Enabled() bool {
	return true
}

func (r *IBMIsBackupPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBackupPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBackupPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_backup_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "match_resource_type"},
			{Name: "match_user_tags"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Validate `name`
		if _, exists := resource.Body.Attributes["name"]; !exists {
			runner.EmitIssue(
				r,
				"`name` attribute must be specified",
				resource.DefRange,
			)
		}

		// Validate `match_resource_type`
		if matchResourceTypeAttr, exists := resource.Body.Attributes["match_resource_type"]; exists {
			var matchResourceType string
			err := runner.EvaluateExpr(matchResourceTypeAttr.Expr, &matchResourceType, nil)
			if err != nil {
				return err
			}
			if matchResourceType == "" {
				runner.EmitIssue(
					r,
					"`match_resource_type` attribute cannot be empty",
					matchResourceTypeAttr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
