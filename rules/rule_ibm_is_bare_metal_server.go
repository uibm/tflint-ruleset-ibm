package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBareMetalServerRule checks the configuration of ibm_is_bare_metal_server resource
type IBMBareMetalServerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBareMetalServerRule returns a new rule
func NewIBMBareMetalServerRule() *IBMBareMetalServerRule {
	return &IBMBareMetalServerRule{
		resourceType: "ibm_is_bare_metal_server",
	}
}

// Name returns the rule name
func (r *IBMBareMetalServerRule) Name() string {
	return "bare_metal_server"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBareMetalServerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBareMetalServerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBareMetalServerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBareMetalServerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "reservation_affinity"},
			{Name: "name"},
			{Name: "primary_network_attachment"},
			{Name: "vpc"},
			{Name: "trusted_platform_module"},
			{Name: "primary_network_interface"},
			{Name: "network_attachments"},
			{Name: "network_interfaces"},
			{Name: "bandwidth"},
			{Name: "enable_secure_boot"},
			{Name: "initialization"},
			{Name: "zone"},
			{Name: "profile"},
			{Name: "resource_group"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["primary_network_attachment"]; !exists {
			if err := runner.EmitIssue(
				r,
				"primary_network_attachment attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["primary_network_interface"]; !exists {
			if err := runner.EmitIssue(
				r,
				"primary_network_interface attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["initialization"]; !exists {
			if err := runner.EmitIssue(
				r,
				"initialization attribute must be specified",
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
		if _, exists := resource.Body.Attributes["profile"]; !exists {
			if err := runner.EmitIssue(
				r,
				"profile attribute must be specified",
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
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$",
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
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'bandwidth' value must be at least 1",
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
				if val > 1e+06 {
					if err := runner.EmitIssue(
						r,
						"'bandwidth' value must not exceed 1e+06",
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
