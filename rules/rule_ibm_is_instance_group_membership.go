package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceGroupMembershipRule checks instance group membership configuration
type IBMIsInstanceGroupMembershipRule struct {
	tflint.DefaultRule
}

// NewIBMIsInstanceGroupMembershipRule returns new rule
func NewIBMIsInstanceGroupMembershipRule() *IBMIsInstanceGroupMembershipRule {
	return &IBMIsInstanceGroupMembershipRule{}
}

// Name returns the rule name
func (r *IBMIsInstanceGroupMembershipRule) Name() string {
	return "ibm_is_instance_group_membership"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceGroupMembershipRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceGroupMembershipRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceGroupMembershipRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIsInstanceGroupMembershipRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_group_membership", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "instance_group"},
			{Name: "instance_group_membership"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{
			"name",
			"instance_group",
			"instance_group_membership",
		}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					"`"+attr+"` attribute must be specified",
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
		}

		// Additional validations can be added here based on IBM Cloud requirements
	}

	return nil
}
