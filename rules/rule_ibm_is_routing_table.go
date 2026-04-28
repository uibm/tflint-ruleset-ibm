package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMRoutingTableRule checks the configuration of ibm_is_routing_table resource
type IBMRoutingTableRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMRoutingTableRule returns a new rule
func NewIBMRoutingTableRule() *IBMRoutingTableRule {
	return &IBMRoutingTableRule{
		resourceType: "ibm_is_routing_table",
	}
}

// Name returns the rule name
func (r *IBMRoutingTableRule) Name() string {
	return "routing_table"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMRoutingTableRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMRoutingTableRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMRoutingTableRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMRoutingTableRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "route_direct_link_ingress"},
			{Name: "route_internet_ingress"},
			{Name: "route_transit_gateway_ingress"},
			{Name: "route_vpc_zone_ingress"},
			{Name: "routes"},
			{Name: "accept_routes_from"},
			{Name: "advertise_routes_to"},
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
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
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$",
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
