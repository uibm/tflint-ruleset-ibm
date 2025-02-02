package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSecurityGroupRuleRule checks security group rule configuration
type IBMIsSecurityGroupRuleRule struct {
	tflint.DefaultRule
}

func NewIBMIsSecurityGroupRuleRule() *IBMIsSecurityGroupRuleRule {
	return &IBMIsSecurityGroupRuleRule{}
}

func (r *IBMIsSecurityGroupRuleRule) Name() string {
	return "ibm_is_security_group_rule"
}

func (r *IBMIsSecurityGroupRuleRule) Enabled() bool {
	return true
}

func (r *IBMIsSecurityGroupRuleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSecurityGroupRuleRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSecurityGroupRuleRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_security_group_rule", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "group"},
			{Name: "direction"},
			{Name: "remote"},
			{Name: "ip_version"},
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
					},
				},
			},
			{
				Type: "udp",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "port_min"},
						{Name: "port_max"},
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
		requiredAttrs := []string{"group", "direction", "remote"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
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

		// Validate protocol blocks
		hasICMP := false
		hasTCP := false
		hasUDP := false

		for _, block := range resource.Body.Blocks {
			switch block.Type {
			case "icmp":
				hasICMP = true
				if err := r.validateICMPBlock(runner, block); err != nil {
					return err
				}
			case "tcp":
				hasTCP = true
				if err := r.validatePortBlock(runner, block, "TCP"); err != nil {
					return err
				}
			case "udp":
				hasUDP = true
				if err := r.validatePortBlock(runner, block, "UDP"); err != nil {
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

		if protocolCount == 0 {
			runner.EmitIssue(
				r,
				"at least one protocol block (icmp, tcp, or udp) must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}

func (r *IBMIsSecurityGroupRuleRule) validateICMPBlock(runner tflint.Runner, block *hclext.Block) error {
	if attr, exists := block.Body.Attributes["type"]; exists {
		var icmpType int
		err := runner.EvaluateExpr(attr.Expr, &icmpType, nil)
		if err != nil {
			return err
		}

		if icmpType < 0 || icmpType > 254 {
			runner.EmitIssue(
				r,
				"ICMP type must be between 0 and 254",
				attr.Expr.Range(),
			)
		}
	}

	if attr, exists := block.Body.Attributes["code"]; exists {
		var icmpCode int
		err := runner.EvaluateExpr(attr.Expr, &icmpCode, nil)
		if err != nil {
			return err
		}

		if icmpCode < 0 || icmpCode > 255 {
			runner.EmitIssue(
				r,
				"ICMP code must be between 0 and 255",
				attr.Expr.Range(),
			)
		}
	}

	return nil
}

// Completing the validatePortBlock function
func (r *IBMIsSecurityGroupRuleRule) validatePortBlock(runner tflint.Runner, block *hclext.Block, protocol string) error {
	var minPort, maxPort int

	if attr, exists := block.Body.Attributes["port_min"]; exists {
		err := runner.EvaluateExpr(attr.Expr, &minPort, nil)
		if err != nil {
			return err
		}

		if minPort < 1 || minPort > 65535 {
			runner.EmitIssue(
				r,
				fmt.Sprintf("%s port_min must be between 1 and 65535", protocol),
				attr.Expr.Range(),
			)
		}
	}

	if attr, exists := block.Body.Attributes["port_max"]; exists {
		err := runner.EvaluateExpr(attr.Expr, &maxPort, nil)
		if err != nil {
			return err
		}

		if maxPort < 1 || maxPort > 65535 {
			runner.EmitIssue(
				r,
				fmt.Sprintf("%s port_max must be between 1 and 65535", protocol),
				attr.Expr.Range(),
			)
		}

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
