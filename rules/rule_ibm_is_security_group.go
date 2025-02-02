package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSecurityGroupRule checks security group configuration
type IBMIsSecurityGroupRule struct {
	tflint.DefaultRule
}

func NewIBMIsSecurityGroupRule() *IBMIsSecurityGroupRule {
	return &IBMIsSecurityGroupRule{}
}

func (r *IBMIsSecurityGroupRule) Name() string {
	return "ibm_is_security_group"
}

func (r *IBMIsSecurityGroupRule) Enabled() bool {
	return true
}

func (r *IBMIsSecurityGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSecurityGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSecurityGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_security_group", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpc"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "vpc"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
