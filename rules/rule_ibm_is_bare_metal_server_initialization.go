package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBareMetalServerInitializationRule checks the configuration of ibm_is_bare_metal_server_initialization resource
type IBMBareMetalServerInitializationRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBareMetalServerInitializationRule returns a new rule
func NewIBMBareMetalServerInitializationRule() *IBMBareMetalServerInitializationRule {
	return &IBMBareMetalServerInitializationRule{
		resourceType: "ibm_is_bare_metal_server_initialization",
	}
}

// Name returns the rule name
func (r *IBMBareMetalServerInitializationRule) Name() string {
	return "bare_metal_server_initialization"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBareMetalServerInitializationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBareMetalServerInitializationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBareMetalServerInitializationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBareMetalServerInitializationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "image"},
			{Name: "keys"},
			{Name: "user_data"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["image"]; !exists {
			if err := runner.EmitIssue(
				r,
				"image attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["keys"]; !exists {
			if err := runner.EmitIssue(
				r,
				"keys attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["user_data"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 65535 {
					if err := runner.EmitIssue(
						r,
						"'user_data' value must not exceed 65535 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["user_data"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[\\s\\S]*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for user_data: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'user_data' value must match pattern: ^[\\s\\S]*$",
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
