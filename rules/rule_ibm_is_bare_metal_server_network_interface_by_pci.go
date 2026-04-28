package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMBareMetalServerNetworkInterfaceByPciRule checks the configuration of ibm_is_bare_metal_server_network_interface_by_pci resource
type IBMBareMetalServerNetworkInterfaceByPciRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMBareMetalServerNetworkInterfaceByPciRule returns a new rule
func NewIBMBareMetalServerNetworkInterfaceByPciRule() *IBMBareMetalServerNetworkInterfaceByPciRule {
	return &IBMBareMetalServerNetworkInterfaceByPciRule{
		resourceType: "ibm_is_bare_metal_server_network_interface_by_pci",
	}
}

// Name returns the rule name
func (r *IBMBareMetalServerNetworkInterfaceByPciRule) Name() string {
	return "bare_metal_server_network_interface_by_pci"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMBareMetalServerNetworkInterfaceByPciRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMBareMetalServerNetworkInterfaceByPciRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMBareMetalServerNetworkInterfaceByPciRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMBareMetalServerNetworkInterfaceByPciRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "allowed_vlans"},
			{Name: "interface_type"},
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
		if attr, exists := resource.Body.Attributes["interface_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"pci",
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
	}

	return nil
}
