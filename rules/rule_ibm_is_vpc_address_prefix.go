package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPCAddressPrefixRule checks VPC address prefix configuration
type IBMIsVPCAddressPrefixRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPCAddressPrefixRule() *IBMIsVPCAddressPrefixRule {
	return &IBMIsVPCAddressPrefixRule{}
}

func (r *IBMIsVPCAddressPrefixRule) Name() string {
	return "ibm_is_vpc_address_prefix"
}

func (r *IBMIsVPCAddressPrefixRule) Enabled() bool {
	return true
}

func (r *IBMIsVPCAddressPrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPCAddressPrefixRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPCAddressPrefixRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpc_address_prefix", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "zone"},
			{Name: "vpc"},
			{Name: "cidr"},
			{Name: "is_default"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "zone", "vpc", "cidr"}
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

		// Validate CIDR format
		if attr, exists := resource.Body.Attributes["cidr"]; exists {
			var cidr string
			err := runner.EvaluateExpr(attr.Expr, &cidr, nil)
			if err != nil {
				return err
			}

			if !isValidCIDR(cidr) {
				runner.EmitIssue(
					r,
					"invalid CIDR format",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
