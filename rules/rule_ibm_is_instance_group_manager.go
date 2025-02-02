package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceGroupManagerRule checks instance group manager configuration
type IBMIsInstanceGroupManagerRule struct {
	tflint.DefaultRule
}

// NewIBMIsInstanceGroupManagerRule returns new rule
func NewIBMIsInstanceGroupManagerRule() *IBMIsInstanceGroupManagerRule {
	return &IBMIsInstanceGroupManagerRule{}
}

// Name returns the rule name
func (r *IBMIsInstanceGroupManagerRule) Name() string {
	return "ibm_is_instance_group_manager"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceGroupManagerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceGroupManagerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceGroupManagerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIsInstanceGroupManagerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_group_manager", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "instance_group"},
			{Name: "manager_type"},
			{Name: "enable_manager"},
			{Name: "max_membership_count"},
			{Name: "min_membership_count"},
			{Name: "cooldown"},
			{Name: "aggregation_window"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "instance_group", "manager_type", "enable_manager"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					"`"+attr+"` attribute must be specified",
					resource.DefRange,
				)
			}
		}

		// Validate manager_type
		if attr, exists := resource.Body.Attributes["manager_type"]; exists {
			var managerType string
			err := runner.EvaluateExpr(attr.Expr, &managerType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"autoscale": true,
				"scheduled": true,
			}

			if !validTypes[managerType] {
				runner.EmitIssue(
					r,
					"manager_type must be either 'autoscale' or 'scheduled'",
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

func (r *IBMIsInstanceGroupManagerRule) validateMembershipCounts(runner tflint.Runner, resource *hclext.Block) error {
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
