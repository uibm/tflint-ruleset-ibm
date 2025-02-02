package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPNServerRule checks VPN server configuration
type IBMIsVPNServerRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPNServerRule() *IBMIsVPNServerRule {
	return &IBMIsVPNServerRule{}
}

func (r *IBMIsVPNServerRule) Name() string {
	return "ibm_is_vpn_server"
}

func (r *IBMIsVPNServerRule) Enabled() bool {
	return true
}

func (r *IBMIsVPNServerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPNServerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPNServerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpn_server", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "certificate_crn"},
			{Name: "client_ip_pool"},
			{Name: "client_dns_server_ips"},
			{Name: "client_idle_timeout"},
			{Name: "enable_split_tunneling"},
			{Name: "port"},
			{Name: "protocol"},
			{Name: "subnets"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "client_authentication",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "method"},
						{Name: "client_ca_crn"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{
			"name",
			"certificate_crn",
			"client_ip_pool",
			"subnets",
		}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check client_authentication block
		hasClientAuth := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "client_authentication" {
				hasClientAuth = true
				if err := r.validateClientAuthentication(runner, block); err != nil {
					return err
				}
			}
		}

		if !hasClientAuth {
			runner.EmitIssue(
				r,
				"client_authentication block must be specified",
				resource.DefRange,
			)
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

		// Validate protocol
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var protocol string
			err := runner.EvaluateExpr(attr.Expr, &protocol, nil)
			if err != nil {
				return err
			}

			validProtocols := map[string]bool{
				"udp": true,
				"tcp": true,
			}

			if !validProtocols[protocol] {
				runner.EmitIssue(
					r,
					"protocol must be either 'udp' or 'tcp'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate port
		if attr, exists := resource.Body.Attributes["port"]; exists {
			var port int
			err := runner.EvaluateExpr(attr.Expr, &port, nil)
			if err != nil {
				return err
			}

			if port < 1 || port > 65535 {
				runner.EmitIssue(
					r,
					"port must be between 1 and 65535",
					attr.Expr.Range(),
				)
			}
		}

		// Validate client_ip_pool CIDR
		if attr, exists := resource.Body.Attributes["client_ip_pool"]; exists {
			var cidr string
			err := runner.EvaluateExpr(attr.Expr, &cidr, nil)
			if err != nil {
				return err
			}

			if !isValidCIDR(cidr) {
				runner.EmitIssue(
					r,
					"client_ip_pool must be a valid CIDR",
					attr.Expr.Range(),
				)
			}
		}

		// Validate client_dns_server_ips
		if attr, exists := resource.Body.Attributes["client_dns_server_ips"]; exists {
			var dnsIPs []string
			err := runner.EvaluateExpr(attr.Expr, &dnsIPs, nil)
			if err != nil {
				return err
			}

			for _, ip := range dnsIPs {
				if !isValidIPv4(ip) {
					runner.EmitIssue(
						r,
						"client_dns_server_ips must contain valid IPv4 addresses",
						attr.Expr.Range(),
					)
					break
				}
			}
		}
	}

	return nil
}

func (r *IBMIsVPNServerRule) validateClientAuthentication(runner tflint.Runner, block *hclext.Block) error {
	// Check required attributes
	requiredAttrs := []string{"method"}
	for _, attr := range requiredAttrs {
		if _, exists := block.Body.Attributes[attr]; !exists {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` attribute must be specified in client_authentication block", attr),
				block.DefRange,
			)
		}
	}

	// Validate authentication method
	if attr, exists := block.Body.Attributes["method"]; exists {
		var method string
		err := runner.EvaluateExpr(attr.Expr, &method, nil)
		if err != nil {
			return err
		}

		validMethods := map[string]bool{
			"certificate": true,
			"username":    true,
		}

		if !validMethods[method] {
			runner.EmitIssue(
				r,
				"method must be either 'certificate' or 'username'",
				attr.Expr.Range(),
			)
		}

		// Check for required attributes based on method
		if method == "certificate" {
			if _, exists := block.Body.Attributes["client_ca_crn"]; !exists {
				runner.EmitIssue(
					r,
					"client_ca_crn must be specified when method is 'certificate'",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
