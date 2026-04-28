package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceGroupManagerScheduledActionGroupRule checks the configuration of ibm_is_instance_group_manager_scheduled_action_group resource
type IBMInstanceGroupManagerScheduledActionGroupRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceGroupManagerScheduledActionGroupRule returns a new rule
func NewIBMInstanceGroupManagerScheduledActionGroupRule() *IBMInstanceGroupManagerScheduledActionGroupRule {
	return &IBMInstanceGroupManagerScheduledActionGroupRule{
		resourceType: "ibm_is_instance_group_manager_scheduled_action_group",
	}
}

// Name returns the rule name
func (r *IBMInstanceGroupManagerScheduledActionGroupRule) Name() string {
	return "instance_group_manager_scheduled_action_group"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceGroupManagerScheduledActionGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceGroupManagerScheduledActionGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceGroupManagerScheduledActionGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceGroupManagerScheduledActionGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "membership_count"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["membership_count"]; !exists {
			if err := runner.EmitIssue(
				r,
				"membership_count attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 0 {
					if err := runner.EmitIssue(
						r,
						"'membership_count' value must be at least 0",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["membership_count"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 100 {
					if err := runner.EmitIssue(
						r,
						"'membership_count' value must not exceed 100",
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
