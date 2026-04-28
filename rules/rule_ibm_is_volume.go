package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVolumeRule checks the configuration of ibm_is_volume resource
type IBMVolumeRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVolumeRule returns a new rule
func NewIBMVolumeRule() *IBMVolumeRule {
	return &IBMVolumeRule{
		resourceType: "ibm_is_volume",
	}
}

// Name returns the rule name
func (r *IBMVolumeRule) Name() string {
	return "volume"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVolumeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVolumeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVolumeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVolumeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bandwidth"},
			{Name: "iops"},
			{Name: "name"},
			{Name: "profile"},
			{Name: "resource_group"},
			{Name: "user_tags"},
			{Name: "zone"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["profile"]; !exists {
			if err := runner.EmitIssue(
				r,
				"profile attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["zone"]; !exists {
			if err := runner.EmitIssue(
				r,
				"zone attribute must be specified",
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
		if attr, exists := resource.Body.Attributes["bandwidth"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1000 {
					if err := runner.EmitIssue(
						r,
						"'bandwidth' value must be at least 1000",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["bandwidth"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 8192 {
					if err := runner.EmitIssue(
						r,
						"'bandwidth' value must not exceed 8192",
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
