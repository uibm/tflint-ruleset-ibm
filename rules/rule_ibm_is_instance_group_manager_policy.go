package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceGroupManagerPolicyRule checks instance group manager policy configuration
type IBMIsInstanceGroupManagerPolicyRule struct {
	tflint.DefaultRule
}

// NewIBMIsInstanceGroupManagerPolicyRule returns new rule
func NewIBMIsInstanceGroupManagerPolicyRule() *IBMIsInstanceGroupManagerPolicyRule {
	return &IBMIsInstanceGroupManagerPolicyRule{}
}

// Name returns the rule name
func (r *IBMIsInstanceGroupManagerPolicyRule) Name() string {
	return "ibm_is_instance_group_manager_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceGroupManagerPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceGroupManagerPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceGroupManagerPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIsInstanceGroupManagerPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_group_manager_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "instance_group"},
			{Name: "instance_group_manager"},
			{Name: "metric_type"},
			{Name: "metric_value"},
			{Name: "policy_type"},
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
			"metric_type",
			"metric_value",
			"policy_type",
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

		// Validate metric_type
		if attr, exists := resource.Body.Attributes["metric_type"]; exists {
			var metricType string
			err := runner.EvaluateExpr(attr.Expr, &metricType, nil)
			if err != nil {
				return err
			}

			validMetricTypes := map[string]bool{
				"cpu":         true,
				"memory":      true,
				"network_in":  true,
				"network_out": true,
				"custom":      true,
			}

			if !validMetricTypes[metricType] {
				runner.EmitIssue(
					r,
					"invalid metric_type. Must be one of: cpu, memory, network_in, network_out, custom",
					attr.Expr.Range(),
				)
			}
		}

		// Validate policy_type
		if attr, exists := resource.Body.Attributes["policy_type"]; exists {
			var policyType string
			err := runner.EvaluateExpr(attr.Expr, &policyType, nil)
			if err != nil {
				return err
			}

			validPolicyTypes := map[string]bool{
				"target":     true,
				"range":      true,
				"percentage": true,
			}

			if !validPolicyTypes[policyType] {
				runner.EmitIssue(
					r,
					"invalid policy_type. Must be one of: target, range, percentage",
					attr.Expr.Range(),
				)
			}
		}

		// Validate metric_value
		if attr, exists := resource.Body.Attributes["metric_value"]; exists {
			var metricValue float64
			err := runner.EvaluateExpr(attr.Expr, &metricValue, nil)
			if err != nil {
				return err
			}

			if metricValue <= 0 {
				runner.EmitIssue(
					r,
					"metric_value must be greater than 0",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
