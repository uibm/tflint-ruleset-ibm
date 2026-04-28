package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBareMetalServerNetworkAttachmentRule checks the configuration of ibm_is_bare_metal_server_network_attachment resource
type IBMBareMetalServerNetworkAttachmentRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBareMetalServerNetworkAttachmentRule returns a new rule
func NewIBMBareMetalServerNetworkAttachmentRule() *IBMBareMetalServerNetworkAttachmentRule {
	return &IBMBareMetalServerNetworkAttachmentRule{
		resourceType: "ibm_is_bare_metal_server_network_attachment",
	}
}

// Name returns the rule name
func (r *IBMBareMetalServerNetworkAttachmentRule) Name() string {
	return "bare_metal_server_network_attachment"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBareMetalServerNetworkAttachmentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBareMetalServerNetworkAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBareMetalServerNetworkAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBareMetalServerNetworkAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "allow_to_float"},
			{Name: "vlan"},
			{Name: "name"},
			{Name: "virtual_network_interface"},
			{Name: "interface_type"},
			{Name: "allowed_vlans"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["interface_type"]; !exists {
			if err := runner.EmitIssue(
				r,
				"interface_type attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["vlan"]; !exists {
			if err := runner.EmitIssue(
				r,
				"vlan attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["virtual_network_interface"]; !exists {
			if err := runner.EmitIssue(
				r,
				"virtual_network_interface attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"vlan",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'interface_type' value must be one of [%s]", strings.Join(validValues, ", ")),
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
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'interface_type' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'interface_type' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for interface_type: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'interface_type' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
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
		if attr, exists := resource.Body.Attributes["vlan"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'vlan' value must be at least 1",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["vlan"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 4094 {
					if err := runner.EmitIssue(
						r,
						"'vlan' value must not exceed 4094",
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
