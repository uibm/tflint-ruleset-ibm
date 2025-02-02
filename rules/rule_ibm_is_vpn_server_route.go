package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPNServerRouteRule checks VPN server route configuration
type IBMIsVPNServerRouteRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPNServerRouteRule() *IBMIsVPNServerRouteRule {
	return &IBMIsVPNServerRouteRule{}
}

func (r *IBMIsVPNServerRouteRule) Name() string {
	return "ibm_is_vpn_server_route"
}

func (r *IBMIsVPNServerRouteRule) Enabled() bool {
	return true
}

func (r *IBMIsVPNServerRouteRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPNServerRouteRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPNServerRouteRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpn_server_route", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "vpn_server"},
			{Name: "destination"},
			{Name: "action"},
			{Name: "name"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"vpn_server", "destination", "action", "name"}
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
				"translate": true,
				"drop":      true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be either 'translate' or 'drop'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate destination CIDR
		if attr, exists := resource.Body.Attributes["destination"]; exists {
			var destination string
			err := runner.EvaluateExpr(attr.Expr, &destination, nil)
			if err != nil {
				return err
			}

			if !isValidCIDR(destination) {
				runner.EmitIssue(
					r,
					"destination must be a valid CIDR",
					attr.Expr.Range(),
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
	}

	return nil
}
