package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPNGatewayConnectionsRule checks VPN gateway connections configuration
type IBMIsVPNGatewayConnectionsRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPNGatewayConnectionsRule() *IBMIsVPNGatewayConnectionsRule {
	return &IBMIsVPNGatewayConnectionsRule{}
}

func (r *IBMIsVPNGatewayConnectionsRule) Name() string {
	return "ibm_is_vpn_gateway_connection"
}

func (r *IBMIsVPNGatewayConnectionsRule) Enabled() bool {
	return true
}

func (r *IBMIsVPNGatewayConnectionsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPNGatewayConnectionsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPNGatewayConnectionsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpn_gateway_connection", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpn_gateway"},
			{Name: "preshared_key"},
			{Name: "admin_state_up"},
			{Name: "dead_peer_detection"},
			{Name: "interval"},
			{Name: "timeout"},
			{Name: "action"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "peer",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "address"},
						{Name: "cidrs"},
					},
				},
			},
			{
				Type: "local",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "cidrs"},
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
		requiredAttrs := []string{"name", "vpn_gateway", "preshared_key"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check for either peer block or peer_address/peer_cidrs
		// hasPeerBlock := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "peer" {
				// hasPeerBlock = true
				if err := r.validatePeerBlock(runner, block); err != nil {
					return err
				}
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

		// Validate dead peer detection settings if specified attr to use ?
		if _, exists := resource.Body.Attributes["dead_peer_detection"]; exists {
			if err := r.validateDPD(runner, resource); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *IBMIsVPNGatewayConnectionsRule) validatePeerBlock(runner tflint.Runner, block *hclext.Block) error {
	// Check required peer attributes
	requiredAttrs := []string{"address", "cidrs"}
	for _, attr := range requiredAttrs {
		if _, exists := block.Body.Attributes[attr]; !exists {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` attribute must be specified in peer block", attr),
				block.DefRange,
			)
		}
	}

	// Validate peer address format
	if attr, exists := block.Body.Attributes["address"]; exists {
		var address string
		err := runner.EvaluateExpr(attr.Expr, &address, nil)
		if err != nil {
			return err
		}

		if !isValidIPv4(address) {
			runner.EmitIssue(
				r,
				"invalid peer address format. Must be a valid IPv4 address",
				attr.Expr.Range(),
			)
		}
	}

	return nil
}

func (r *IBMIsVPNGatewayConnectionsRule) validateDPD(runner tflint.Runner, resource *hclext.Block) error {
	dpdFields := []string{"interval", "timeout", "action"}

	for _, field := range dpdFields {
		if _, exists := resource.Body.Attributes[field]; !exists {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` must be specified when dead_peer_detection is configured", field),
				resource.DefRange,
			)
		}
	}

	// Validate action if specified
	if attr, exists := resource.Body.Attributes["action"]; exists {
		var action string
		err := runner.EvaluateExpr(attr.Expr, &action, nil)
		if err != nil {
			return err
		}

		validActions := map[string]bool{
			"restart": true,
			"clear":   true,
			"hold":    true,
			"none":    true,
		}

		if !validActions[action] {
			runner.EmitIssue(
				r,
				"action must be one of: restart, clear, hold, none",
				attr.Expr.Range(),
			)
		}
	}

	return nil
}
