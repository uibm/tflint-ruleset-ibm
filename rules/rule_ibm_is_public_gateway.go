package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPublicGatewayRule checks public gateway configuration
type IBMIsPublicGatewayRule struct {
	tflint.DefaultRule
}

func NewIBMIsPublicGatewayRule() *IBMIsPublicGatewayRule {
	return &IBMIsPublicGatewayRule{}
}

func (r *IBMIsPublicGatewayRule) Name() string {
	return "ibm_is_public_gateway"
}

func (r *IBMIsPublicGatewayRule) Enabled() bool {
	return true
}

func (r *IBMIsPublicGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPublicGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPublicGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_public_gateway", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpc"},
			{Name: "zone"},
			{Name: "resource_group"},
			{Name: "floating_ip"},
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
	}

	return nil
}
