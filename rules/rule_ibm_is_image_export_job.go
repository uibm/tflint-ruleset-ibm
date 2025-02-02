package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsImageExportJobRule checks image export job configuration
type IBMIsImageExportJobRule struct {
	tflint.DefaultRule
}

func NewIBMIsImageExportJobRule() *IBMIsImageExportJobRule {
	return &IBMIsImageExportJobRule{}
}

func (r *IBMIsImageExportJobRule) Name() string {
	return "ibm_is_image_export_job"
}

func (r *IBMIsImageExportJobRule) Enabled() bool {
	return true
}

func (r *IBMIsImageExportJobRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsImageExportJobRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsImageExportJobRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_image_export_job", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "image"},
			{Name: "name"},
			{Name: "format"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "storage_bucket",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
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
		requiredAttrs := []string{"image", "name"}
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

		// Validate format if specified
		if attr, exists := resource.Body.Attributes["format"]; exists {
			var format string
			err := runner.EvaluateExpr(attr.Expr, &format, nil)
			if err != nil {
				return err
			}

			validFormats := map[string]bool{
				"qcow2": true,
				"raw":   true,
				"vhd":   true,
				"vmdk":  true,
			}

			if !validFormats[format] {
				runner.EmitIssue(
					r,
					"format must be one of: qcow2, raw, vhd, vmdk",
					attr.Expr.Range(),
				)
			}
		}

		// Check storage_bucket block
		hasStorageBucket := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "storage_bucket" {
				hasStorageBucket = true
				if _, exists := block.Body.Attributes["name"]; !exists {
					runner.EmitIssue(
						r,
						"name attribute must be specified in storage_bucket block",
						block.DefRange,
					)
				}
			}
		}

		if !hasStorageBucket {
			runner.EmitIssue(
				r,
				"storage_bucket block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
