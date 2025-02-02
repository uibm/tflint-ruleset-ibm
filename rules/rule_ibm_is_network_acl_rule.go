package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsNetworkACLRuleRule checks network ACL rule configuration
type IBMIsNetworkACLRuleRule struct {
	tflint.DefaultRule
}

func NewIBMIsNetworkACLRuleRule() *IBMIsNetworkACLRuleRule {
	return &IBMIsNetworkACLRuleRule{}
}

func (r *IBMIsNetworkACLRuleRule) Name() string {
	return "ibm_is_network_acl_rule"
}

func (r *IBMIsNetworkACLRuleRule) Enabled() bool {
	return true
}

func (r *IBMIsNetworkACLRuleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsNetworkACLRuleRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsNetworkACLRuleRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_network_acl_rule", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "network_acl"},
			{Name: "name"},
			{Name: "action"},
			{Name: "source"},
			{Name: "destination"},
			{Name: "direction"},
			{Name: "protocol"},
			{Name: "before"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "icmp",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "code"},
						{Name: "type"},
					},
				},
			},
			{
				Type: "tcp",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "port_min"},
						{Name: "port_max"},
						{Name: "source_port_min"},
						{Name: "source_port_max"},
					},
				},
			},
			{
				Type: "udp",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "port_min"},
						{Name: "port_max"},
						{Name: "source_port_min"},
						{Name: "source_port_max"},
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
			"network_acl",
			"name",
			"action",
			"source",
			"destination",
			"direction",
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

		// Validate action
		if attr, exists := resource.Body.Attributes["action"]; exists {
			var action string
			err := runner.EvaluateExpr(attr.Expr, &action, nil)
			if err != nil {
				return err
			}

			validActions := map[string]bool{
				"allow": true,
				"deny":  true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be either 'allow' or 'deny'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate direction
		if attr, exists := resource.Body.Attributes["direction"]; exists {
			var direction string
			err := runner.EvaluateExpr(attr.Expr, &direction, nil)
			if err != nil {
				return err
			}

			validDirections := map[string]bool{
				"inbound":  true,
				"outbound": true,
			}

			if !validDirections[direction] {
				runner.EmitIssue(
					r,
					"direction must be either 'inbound' or 'outbound'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate protocol specifics
		hasICMP := false
		hasTCP := false
		hasUDP := false

		for _, block := range resource.Body.Blocks {
			switch block.Type {
			case "icmp":
				hasICMP = true
				err := r.validateICMPBlock(runner, block)
				if err != nil {
					return err
				}
			case "tcp":
				hasTCP = true
				err := r.validateTCPBlock(runner, block)
				if err != nil {
					return err
				}
			case "udp":
				hasUDP = true
				err := r.validateUDPBlock(runner, block)
				if err != nil {
					return err
				}
			}
		}

		// Check that only one protocol block is specified
		protocolCount := 0
		if hasICMP {
			protocolCount++
		}
		if hasTCP {
			protocolCount++
		}
		if hasUDP {
			protocolCount++
		}

		if protocolCount > 1 {
			runner.EmitIssue(
				r,
				"only one protocol block (icmp, tcp, or udp) can be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}

func (r *IBMIsNetworkACLRuleRule) validateICMPBlock(runner tflint.Runner, block *hclext.Block) error {
	// Validate ICMP type and code
	if attr, exists := block.Body.Attributes["type"]; exists {
		var icmpType int
		if err := runner.EvaluateExpr(attr.Expr, &icmpType, nil); err == nil {
			if icmpType < 0 || icmpType > 254 {
				runner.EmitIssue(
					r,
					"ICMP type must be between 0 and 254",
					attr.Expr.Range(),
				)
			}
		}
	}

	if attr, exists := block.Body.Attributes["code"]; exists {
		var icmpCode int
		if err := runner.EvaluateExpr(attr.Expr, &icmpCode, nil); err == nil {
			if icmpCode < 0 || icmpCode > 255 {
				runner.EmitIssue(
					r,
					"ICMP code must be between 0 and 255",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}

func (r *IBMIsNetworkACLRuleRule) validateTCPBlock(runner tflint.Runner, block *hclext.Block) error {
	return r.validatePortBlock(runner, block, "TCP")
}

func (r *IBMIsNetworkACLRuleRule) validateUDPBlock(runner tflint.Runner, block *hclext.Block) error {
	return r.validatePortBlock(runner, block, "UDP")
}

func (r *IBMIsNetworkACLRuleRule) validatePortBlock(runner tflint.Runner, block *hclext.Block, protocol string) error {
	// Validate port ranges
	portAttrs := []string{"port_min", "port_max", "source_port_min", "source_port_max"}
	for _, attrName := range portAttrs {
		if attr, exists := block.Body.Attributes[attrName]; exists {
			var port int
			if err := runner.EvaluateExpr(attr.Expr, &port, nil); err == nil {
				if port < 1 || port > 65535 {
					runner.EmitIssue(
						r,
						fmt.Sprintf("%s port must be between 1 and 65535", protocol),
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	// Validate port ranges are consistent
	var minPort, maxPort int
	if attr, exists := block.Body.Attributes["port_min"]; exists {
		runner.EvaluateExpr(attr.Expr, &minPort, nil)
	}
	if attr, exists := block.Body.Attributes["port_max"]; exists {
		runner.EvaluateExpr(attr.Expr, &maxPort, nil)
		if maxPort < minPort {
			runner.EmitIssue(
				r,
				fmt.Sprintf("%s port_max cannot be less than port_min", protocol),
				attr.Expr.Range(),
			)
		}
	}

	return nil
}
