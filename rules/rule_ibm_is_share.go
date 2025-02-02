package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsShareRule checks file share configuration
type IBMIsShareRule struct {
	tflint.DefaultRule
}

func NewIBMIsShareRule() *IBMIsShareRule {
	return &IBMIsShareRule{}
}

func (r *IBMIsShareRule) Name() string {
	return "ibm_is_share"
}

func (r *IBMIsShareRule) Enabled() bool {
	return true
}

func (r *IBMIsShareRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsShareRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsShareRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_share", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "size"},
			{Name: "profile"},
			{Name: "zone"},
			{Name: "resource_group"},
			{Name: "encryption_key"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "size", "profile", "zone"}
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

		// Validate size
		if attr, exists := resource.Body.Attributes["size"]; exists {
			var size int
			err := runner.EvaluateExpr(attr.Expr, &size, nil)
			if err != nil {
				return err
			}

			if size < 10 || size > 16000 {
				runner.EmitIssue(
					r,
					"size must be between 10 and 16000 GB",
					attr.Expr.Range(),
				)
			}
		}

		// Validate profile
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			validProfiles := map[string]bool{
				"dp2": true,
				"dp4": true,
				"dp8": true,
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"profile must be one of: dp2, dp4, dp8",
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
