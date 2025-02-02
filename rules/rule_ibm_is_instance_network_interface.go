package rules

import (
	"fmt"
	"net"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceNetworkInterfaceRule checks instance network interface configuration
type IBMIsInstanceNetworkInterfaceRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceNetworkInterfaceRule() *IBMIsInstanceNetworkInterfaceRule {
	return &IBMIsInstanceNetworkInterfaceRule{}
}

func (r *IBMIsInstanceNetworkInterfaceRule) Name() string {
	return "ibm_is_instance_network_interface"
}

func (r *IBMIsInstanceNetworkInterfaceRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceNetworkInterfaceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceNetworkInterfaceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceNetworkInterfaceRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_network_interface", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "instance"},
			{Name: "subnet"},
			{Name: "name"},
			{Name: "allow_ip_spoofing"},
			{Name: "primary_ipv4_address"},
			{Name: "security_groups"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"instance", "subnet", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}

		// Validate primary_ipv4_address if specified
		if attr, exists := resource.Body.Attributes["primary_ipv4_address"]; exists {
			var ipv4 string
			err := runner.EvaluateExpr(attr.Expr, &ipv4, nil)
			if err != nil {
				return err
			}

			if !isValidIPv4(ipv4) {
				runner.EmitIssue(
					r,
					"invalid IPv4 address format",
					attr.Expr.Range(),
				)
			}
		}

		// Validate allow_ip_spoofing if specified
		if attr, exists := resource.Body.Attributes["allow_ip_spoofing"]; exists {
			var allowSpoofing bool
			err := runner.EvaluateExpr(attr.Expr, &allowSpoofing, nil)
			if err != nil {
				return err
			}
		}

		// Additional validation for security_groups if needed
	}

	return nil
}

func isValidIPv4(ip string) bool {
	if ip == "" {
		return false
	}
	return net.ParseIP(ip) != nil
}
