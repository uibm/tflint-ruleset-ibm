package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsDedicatedHostRule checks dedicated host configuration
type IBMIsDedicatedHostRule struct {
	tflint.DefaultRule
}

func NewIBMIsDedicatedHostRule() *IBMIsDedicatedHostRule {
	return &IBMIsDedicatedHostRule{}
}

func (r *IBMIsDedicatedHostRule) Name() string {
	return "ibm_is_dedicated_host"
}

func (r *IBMIsDedicatedHostRule) Enabled() bool {
	return true
}

func (r *IBMIsDedicatedHostRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsDedicatedHostRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsDedicatedHostRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_dedicated_host", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "profile"},
			{Name: "name"},
			{Name: "host_group"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"profile", "name", "host_group"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
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
				"bx2d-host-152x608":  true,
				"mx2d-host-304x1216": true,
				"cx2d-host-76x304":   true,
				// Add other valid profiles
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid dedicated host profile specified",
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

			// Add additional name validation rules if needed
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
