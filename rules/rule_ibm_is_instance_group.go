package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceGroupRule checks the configuration of ibm_is_instance_group resource
type IBMInstanceGroupRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceGroupRule returns a new rule
func NewIBMInstanceGroupRule() *IBMInstanceGroupRule {
	return &IBMInstanceGroupRule{
		resourceType: "ibm_is_instance_group",
	}
}

// Name returns the rule name
func (r *IBMInstanceGroupRule) Name() string {
	return "instance_group"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "membership_count"},
			{Name: "name"},
			{Name: "resource_group"},
			{Name: "subnets"},
			{Name: "application_port"},
			{Name: "instance_template"},
			{Name: "load_balancer"},
			{Name: "load_balancer_pool"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["subnets"]; !exists {
			if err := runner.EmitIssue(
				r,
				"subnets attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["instance_template"]; !exists {
			if err := runner.EmitIssue(
				r,
				"instance_template attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'name' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 63 {
					if err := runner.EmitIssue(
						r,
						"'name' value must not exceed 63 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["application_port"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'application_port' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["application_port"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 65535 {
					if err := runner.EmitIssue(
						r,
						"'application_port' value must not exceed 65535",
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
				if val > 1000 {
					if err := runner.EmitIssue(
						r,
						"'membership_count' value must not exceed 1000",
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
