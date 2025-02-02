package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsLBRule checks load balancer configuration
type IBMIsLBRule struct {
	tflint.DefaultRule
}

func NewIBMIsLBRule() *IBMIsLBRule {
	return &IBMIsLBRule{}
}

func (r *IBMIsLBRule) Name() string {
	return "ibm_is_lb"
}

func (r *IBMIsLBRule) Enabled() bool {
	return true
}

func (r *IBMIsLBRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsLBRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsLBRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_lb", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "subnets"},
			{Name: "type"},
			{Name: "security_groups"},
			{Name: "resource_group"},
			{Name: "profile"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "subnets"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate type if specified
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var lbType string
			err := runner.EvaluateExpr(attr.Expr, &lbType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"public":  true,
				"private": true,
			}

			if !validTypes[lbType] {
				runner.EmitIssue(
					r,
					"type must be either 'public' or 'private'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate profile if specified
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			validProfiles := map[string]bool{
				"network-fixed":    true,
				"network-small":    true,
				"network-medium":   true,
				"network-large":    true,
				"network-xlarge":   true,
				"network-xxlarge":  true,
				"network-xxxlarge": true,
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid load balancer profile specified",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
