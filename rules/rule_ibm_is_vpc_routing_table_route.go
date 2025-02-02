package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPCRoutingTableRouteRule checks VPC routing table route configuration
type IBMIsVPCRoutingTableRouteRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPCRoutingTableRouteRule() *IBMIsVPCRoutingTableRouteRule {
	return &IBMIsVPCRoutingTableRouteRule{}
}

func (r *IBMIsVPCRoutingTableRouteRule) Name() string {
	return "ibm_is_vpc_routing_table_route"
}

func (r *IBMIsVPCRoutingTableRouteRule) Enabled() bool {
	return true
}

func (r *IBMIsVPCRoutingTableRouteRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPCRoutingTableRouteRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPCRoutingTableRouteRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpc_routing_table_route", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "vpc"},
			{Name: "routing_table"},
			{Name: "zone"},
			{Name: "name"},
			{Name: "destination"},
			{Name: "action"},
			{Name: "next_hop"},
			{Name: "advertise"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"vpc", "routing_table", "destination", "action", "zone", "name"}
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

		// Validate action
		if attr, exists := resource.Body.Attributes["action"]; exists {
			var action string
			err := runner.EvaluateExpr(attr.Expr, &action, nil)
			if err != nil {
				return err
			}

			validActions := map[string]bool{
				"deliver": true,
				"drop":    true,
			}

			if !validActions[action] {
				runner.EmitIssue(
					r,
					"action must be either 'deliver' or 'drop'",
					attr.Expr.Range(),
				)
			}

			// If action is deliver, next_hop must be specified
			if action == "deliver" {
				if _, exists := resource.Body.Attributes["next_hop"]; !exists {
					runner.EmitIssue(
						r,
						"next_hop must be specified when action is 'deliver'",
						attr.Expr.Range(),
					)
				}
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
					"invalid destination CIDR format",
					attr.Expr.Range(),
				)
			}
		}

		// Validate zone format
		if attr, exists := resource.Body.Attributes["zone"]; exists {
			var zone string
			err := runner.EvaluateExpr(attr.Expr, &zone, nil)
			if err != nil {
				return err
			}

			if !isValidZone(zone) {
				runner.EmitIssue(
					r,
					"invalid zone format. Must be in format: region-number (e.g., us-south-1)",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
