package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVolumeRule checks volume configuration
type IBMIsVolumeRule struct {
	tflint.DefaultRule
}

func NewIBMIsVolumeRule() *IBMIsVolumeRule {
	return &IBMIsVolumeRule{}
}

func (r *IBMIsVolumeRule) Name() string {
	return "ibm_is_volume"
}

func (r *IBMIsVolumeRule) Enabled() bool {
	return true
}

func (r *IBMIsVolumeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVolumeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVolumeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_volume", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "profile"},
			{Name: "zone"},
			{Name: "iops"},
			{Name: "capacity"},
			{Name: "encryption_key"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "profile", "zone", "capacity"}
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

		// Validate profile
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			validProfiles := map[string]bool{
				"general-purpose": true,
				"custom":          true,
				"5iops-tier":      true,
				"10iops-tier":     true,
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"profile must be one of: general-purpose, custom, 5iops-tier, 10iops-tier",
					attr.Expr.Range(),
				)
			}

			// Check IOPS for custom profile
			if profile == "custom" {
				if _, exists := resource.Body.Attributes["iops"]; !exists {
					runner.EmitIssue(
						r,
						"iops must be specified when profile is 'custom'",
						attr.Expr.Range(),
					)
				}
			}
		}

		// Validate capacity
		if attr, exists := resource.Body.Attributes["capacity"]; exists {
			var capacity int
			err := runner.EvaluateExpr(attr.Expr, &capacity, nil)
			if err != nil {
				return err
			}

			if capacity < 10 || capacity > 16000 {
				runner.EmitIssue(
					r,
					"capacity must be between 10 and 16000 GB",
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

		// Validate IOPS if specified
		if attr, exists := resource.Body.Attributes["iops"]; exists {
			var iops int
			err := runner.EvaluateExpr(attr.Expr, &iops, nil)
			if err != nil {
				return err
			}

			if iops < 100 || iops > 48000 {
				runner.EmitIssue(
					r,
					"iops must be between 100 and 48000",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
