package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPCRoutingTableRule checks VPC routing table configuration
type IBMIsVPCRoutingTableRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPCRoutingTableRule() *IBMIsVPCRoutingTableRule {
	return &IBMIsVPCRoutingTableRule{}
}

func (r *IBMIsVPCRoutingTableRule) Name() string {
	return "ibm_is_vpc_routing_table"
}

func (r *IBMIsVPCRoutingTableRule) Enabled() bool {
	return true
}

func (r *IBMIsVPCRoutingTableRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPCRoutingTableRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPCRoutingTableRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpc_routing_table", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "vpc"},
			{Name: "name"},
			{Name: "route_direct_link_ingress"},
			{Name: "route_transit_gateway_ingress"},
			{Name: "route_vpc_zone_ingress"},
			{Name: "accept_routes_from_resource_type"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"vpc", "name"}
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

		// Validate accept_routes_from_resource_type if specified
		if attr, exists := resource.Body.Attributes["accept_routes_from_resource_type"]; exists {
			var resourceType []string
			err := runner.EvaluateExpr(attr.Expr, &resourceType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"vpn_server":  true,
				"vpn_gateway": true,
			}

			for _, rType := range resourceType {
				if !validTypes[rType] {
					runner.EmitIssue(
						r,
						"accept_routes_from_resource_type must be either 'vpn_server' or 'vpn_gateway'",
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	return nil
}
