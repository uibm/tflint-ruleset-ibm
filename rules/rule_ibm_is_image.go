package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsImageRule checks image configuration
type IBMIsImageRule struct {
	tflint.DefaultRule
}

func NewIBMIsImageRule() *IBMIsImageRule {
	return &IBMIsImageRule{}
}

func (r *IBMIsImageRule) Name() string {
	return "ibm_is_image"
}

func (r *IBMIsImageRule) Enabled() bool {
	return true
}

func (r *IBMIsImageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsImageRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsImageRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_image", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "href"},
			{Name: "operating_system"},
			{Name: "encrypted_data_key"},
			{Name: "encryption_key"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "href", "operating_system"}
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

		// Validate href format
		if attr, exists := resource.Body.Attributes["href"]; exists {
			var href string
			err := runner.EvaluateExpr(attr.Expr, &href, nil)
			if err != nil {
				return err
			}

			if !isValidCOSURL(href) {
				runner.EmitIssue(
					r,
					"href must be a valid Cloud Object Storage URL",
					attr.Expr.Range(),
				)
			}
		}

		// Check encryption attributes
		hasEncryptedDataKey := false
		hasEncryptionKey := false

		if _, exists := resource.Body.Attributes["encrypted_data_key"]; exists {
			hasEncryptedDataKey = true
		}
		if _, exists := resource.Body.Attributes["encryption_key"]; exists {
			hasEncryptionKey = true
		}

		if hasEncryptedDataKey != hasEncryptionKey {
			runner.EmitIssue(
				r,
				"both encrypted_data_key and encryption_key must be specified together for encryption",
				resource.DefRange,
			)
		}
	}

	return nil
}

// isValidCOSURL validates if the URL is a valid Cloud Object Storage URL
func isValidCOSURL(url string) bool {
	// Add proper COS URL validation logic here
	// This is a placeholder - implement proper URL format validation
	return len(url) > 0
}
