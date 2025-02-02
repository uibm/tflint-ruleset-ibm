package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetReservedIPRule checks subnet reserved IP configuration
type IBMIsSubnetReservedIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetReservedIPRule() *IBMIsSubnetReservedIPRule {
	return &IBMIsSubnetReservedIPRule{}
}

func (r *IBMIsSubnetReservedIPRule) Name() string {
	return "ibm_is_subnet_reserved_ip"
}

func (r *IBMIsSubnetReservedIPRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetReservedIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetReservedIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetReservedIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet_reserved_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "subnet"},
			{Name: "name"},
			{Name: "address"},
			{Name: "auto_delete"},
			{Name: "target"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"subnet"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format if specified
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}

		// Validate IP address if specified
		if attr, exists := resource.Body.Attributes["address"]; exists {
			var address string
			err := runner.EvaluateExpr(attr.Expr, &address, nil)
			if err != nil {
				return err
			}

			if !isValidIPv4(address) {
				runner.EmitIssue(
					r,
					"invalid IPv4 address format",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
