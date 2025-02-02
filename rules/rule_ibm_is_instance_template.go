package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceTemplateRule checks instance template configuration
type IBMIsInstanceTemplateRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceTemplateRule() *IBMIsInstanceTemplateRule {
	return &IBMIsInstanceTemplateRule{}
}

func (r *IBMIsInstanceTemplateRule) Name() string {
	return "ibm_is_instance_template"
}

func (r *IBMIsInstanceTemplateRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceTemplateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceTemplateRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceTemplateRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_template", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "profile"},
			{Name: "image"},
			{Name: "vpc"},
			{Name: "zone"},
			{Name: "keys"},
			{Name: "resource_group"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "primary_network_attachment",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
					},
					Blocks: []hclext.BlockSchema{
						{
							Type: "virtual_network_interface",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{
									{Name: "id"},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "profile", "image", "vpc", "zone", "keys"}
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
				"bx2-2x8":  true,
				"bx2-4x16": true,
				"cx2-2x4":  true,
				"mx2-2x16": true,
				// Add other valid profiles
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid instance profile specified",
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

		// Check primary_network_attachment block
		for _, block := range resource.Body.Blocks {
			if block.Type == "primary_network_attachment" {
				if _, exists := block.Body.Attributes["name"]; !exists {
					runner.EmitIssue(
						r,
						"name attribute must be specified in primary_network_attachment block",
						block.DefRange,
					)
				}

				hasVNI := false
				for _, vniBlock := range block.Body.Blocks {
					if vniBlock.Type == "virtual_network_interface" {
						hasVNI = true
						if _, exists := vniBlock.Body.Attributes["id"]; !exists {
							runner.EmitIssue(
								r,
								"id attribute must be specified in virtual_network_interface block",
								vniBlock.DefRange,
							)
						}
					}
				}

				if !hasVNI {
					runner.EmitIssue(
						r,
						"virtual_network_interface block must be specified in primary_network_attachment",
						block.DefRange,
					)
				}
			}
		}
	}

	return nil
}
