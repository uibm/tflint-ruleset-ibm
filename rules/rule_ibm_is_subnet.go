package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetRule checks subnet configuration
type IBMIsSubnetRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetRule() *IBMIsSubnetRule {
	return &IBMIsSubnetRule{}
}

func (r *IBMIsSubnetRule) Name() string {
	return "ibm_is_subnet"
}

func (r *IBMIsSubnetRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpc"},
			{Name: "zone"},
			{Name: "ipv4_cidr_block"},
			{Name: "total_ipv4_address_count"},
			{Name: "resource_group"},
			{Name: "public_gateway"},
			{Name: "routing_table"},
			{Name: "network_acl"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "vpc", "zone"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check if either ipv4_cidr_block or total_ipv4_address_count is specified
		hasCIDR := false
		hasAddressCount := false

		if _, exists := resource.Body.Attributes["ipv4_cidr_block"]; exists {
			hasCIDR = true
		}
		if _, exists := resource.Body.Attributes["total_ipv4_address_count"]; exists {
			hasAddressCount = true
		}

		if !hasCIDR && !hasAddressCount {
			runner.EmitIssue(
				r,
				"either ipv4_cidr_block or total_ipv4_address_count must be specified",
				resource.DefRange,
			)
		}

		if hasCIDR && hasAddressCount {
			runner.EmitIssue(
				r,
				"cannot specify both ipv4_cidr_block and total_ipv4_address_count",
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

		// Validate CIDR block if specified
		if attr, exists := resource.Body.Attributes["ipv4_cidr_block"]; exists {
			var cidr string
			err := runner.EvaluateExpr(attr.Expr, &cidr, nil)
			if err != nil {
				return err
			}

			if !isValidCIDR(cidr) {
				runner.EmitIssue(
					r,
					"invalid CIDR block format",
					attr.Expr.Range(),
				)
			}
		}

		// Validate total_ipv4_address_count if specified
		if attr, exists := resource.Body.Attributes["total_ipv4_address_count"]; exists {
			var count int
			err := runner.EvaluateExpr(attr.Expr, &count, nil)
			if err != nil {
				return err
			}

			if count < 8 || count > 8192 || !isPowerOfTwo(count) {
				runner.EmitIssue(
					r,
					"total_ipv4_address_count must be a power of 2 between 8 and 8192",
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
