package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceGroupRule checks instance group configuration
type IBMIsInstanceGroupRule struct {
	tflint.DefaultRule
}

// NewIBMIsInstanceGroupRule returns new rule
func NewIBMIsInstanceGroupRule() *IBMIsInstanceGroupRule {
	return &IBMIsInstanceGroupRule{}
}

// Name returns the rule name
func (r *IBMIsInstanceGroupRule) Name() string {
	return "ibm_is_instance_group"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIsInstanceGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_group", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "instance_template"},
			{Name: "instance_count"},
			{Name: "subnets"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "instance_template", "subnets"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					"`"+attr+"` attribute must be specified",
					resource.DefRange,
				)
			}
		}

		// Validate instance_count if specified
		if attr, exists := resource.Body.Attributes["instance_count"]; exists {
			var count int
			err := runner.EvaluateExpr(attr.Expr, &count, nil)
			if err != nil {
				return err
			}

			if count < 0 {
				runner.EmitIssue(
					r,
					"instance_count must be greater than or equal to 0",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
