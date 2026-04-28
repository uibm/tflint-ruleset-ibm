package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBackupPolicyScopeRule checks the configuration of ibm_is_backup_policy_scope resource
type IBMBackupPolicyScopeRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBackupPolicyScopeRule returns a new rule
func NewIBMBackupPolicyScopeRule() *IBMBackupPolicyScopeRule {
	return &IBMBackupPolicyScopeRule{
		resourceType: "ibm_is_backup_policy_scope",
	}
}

// Name returns the rule name
func (r *IBMBackupPolicyScopeRule) Name() string {
	return "backup_policy_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBackupPolicyScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBackupPolicyScopeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBackupPolicyScopeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBackupPolicyScopeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "crn"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["crn"]; !exists {
			if err := runner.EmitIssue(
				r,
				"crn attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
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
