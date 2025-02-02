package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsDedicatedHostGroupRule checks dedicated host group configuration
type IBMIsDedicatedHostGroupRule struct {
	tflint.DefaultRule
}

func NewIBMIsDedicatedHostGroupRule() *IBMIsDedicatedHostGroupRule {
	return &IBMIsDedicatedHostGroupRule{}
}

func (r *IBMIsDedicatedHostGroupRule) Name() string {
	return "ibm_is_dedicated_host_group"
}

func (r *IBMIsDedicatedHostGroupRule) Enabled() bool {
	return true
}

func (r *IBMIsDedicatedHostGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsDedicatedHostGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsDedicatedHostGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_dedicated_host_group", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "family"},
			{Name: "class"},
			{Name: "zone"},
			{Name: "name"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"family", "class", "zone", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate family
		if attr, exists := resource.Body.Attributes["family"]; exists {
			var family string
			err := runner.EvaluateExpr(attr.Expr, &family, nil)
			if err != nil {
				return err
			}

			validFamilies := map[string]bool{
				"memory":   true,
				"balanced": true,
				"compute":  true,
			}

			if !validFamilies[family] {
				runner.EmitIssue(
					r,
					"family must be one of: memory, balanced, compute",
					attr.Expr.Range(),
				)
			}
		}

		// Validate class
		if attr, exists := resource.Body.Attributes["class"]; exists {
			var class string
			err := runner.EvaluateExpr(attr.Expr, &class, nil)
			if err != nil {
				return err
			}

			validClasses := map[string]bool{
				"bx2d": true,
				"mx2d": true,
				"cx2d": true,
			}

			if !validClasses[class] {
				runner.EmitIssue(
					r,
					"class must be one of: bx2d, mx2d, cx2d",
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

// isValidZone checks if the zone format is valid
func isValidZone(zone string) bool {
	// Add proper zone validation logic here
	// This is a placeholder - implement proper zone format validation
	return len(zone) > 0
}
