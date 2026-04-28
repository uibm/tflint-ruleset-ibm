package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMSecurityGroupRuleLocalRule checks the configuration of ibm_is_security_group_rule_local resource
type IBMSecurityGroupRuleLocalRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMSecurityGroupRuleLocalRule returns a new rule
func NewIBMSecurityGroupRuleLocalRule() *IBMSecurityGroupRuleLocalRule {
	return &IBMSecurityGroupRuleLocalRule{
		resourceType: "ibm_is_security_group_rule_local",
	}
}

// Name returns the rule name
func (r *IBMSecurityGroupRuleLocalRule) Name() string {
	return "security_group_rule_local"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMSecurityGroupRuleLocalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMSecurityGroupRuleLocalRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMSecurityGroupRuleLocalRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMSecurityGroupRuleLocalRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "address"},
			{Name: "cidr_block"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["address"]; !exists {
			if err := runner.EmitIssue(
				r,
				"address attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["cidr_block"]; !exists {
			if err := runner.EmitIssue(
				r,
				"cidr_block attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
