// rules/rule_ibm_is_instance_group_manager_action.go
package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceGroupManagerActionRule checks instance group manager action configuration
type IBMIsInstanceGroupManagerActionRule struct {
	tflint.DefaultRule
}

// NewIBMIsInstanceGroupManagerActionRule returns new rule
func NewIBMIsInstanceGroupManagerActionRule() *IBMIsInstanceGroupManagerActionRule {
	return &IBMIsInstanceGroupManagerActionRule{}
}

// Name returns the rule name
func (r *IBMIsInstanceGroupManagerActionRule) Name() string {
	return "ibm_is_instance_group_manager_action"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceGroupManagerActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceGroupManagerActionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceGroupManagerActionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIsInstanceGroupManagerActionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_group_manager_action", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "instance_group"},
			{Name: "instance_group_manager"},
			{Name: "cron_spec"},
			{Name: "target_manager"},
			{Name: "min_membership_count"},
			{Name: "max_membership_count"},
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
			"instance_group_manager",
			"cron_spec",
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

		// Validate cron_spec format if specified
		if attr, exists := resource.Body.Attributes["cron_spec"]; exists {
			var cronSpec string
			err := runner.EvaluateExpr(attr.Expr, &cronSpec, nil)
			if err != nil {
				return err
			}

			if !isValidCronSpec(cronSpec) {
				runner.EmitIssue(
					r,
					"invalid cron_spec format",
					attr.Expr.Range(),
				)
			}
		}

		// Validate membership counts if specified
		if err := r.validateMembershipCounts(runner, resource); err != nil {
			return err
		}
	}

	return nil
}

func (r *IBMIsInstanceGroupManagerActionRule) validateMembershipCounts(runner tflint.Runner, resource *hclext.Block) error {
	var minCount, maxCount int

	if minAttr, exists := resource.Body.Attributes["min_membership_count"]; exists {
		if err := runner.EvaluateExpr(minAttr.Expr, &minCount, nil); err != nil {
			return err
		}
		if minCount < 0 {
			runner.EmitIssue(
				r,
				"min_membership_count must be greater than or equal to 0",
				minAttr.Expr.Range(),
			)
		}
	}

	if maxAttr, exists := resource.Body.Attributes["max_membership_count"]; exists {
		if err := runner.EvaluateExpr(maxAttr.Expr, &maxCount, nil); err != nil {
			return err
		}
		if maxCount < minCount {
			runner.EmitIssue(
				r,
				"max_membership_count must be greater than or equal to min_membership_count",
				maxAttr.Expr.Range(),
			)
		}
	}

	return nil
}

func isValidCronSpec(spec string) bool {
	// Add proper cron validation logic here
	// This is a basic placeholder - implement full cron validation
	return len(spec) > 0
}
