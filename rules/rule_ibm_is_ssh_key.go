package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSSHKeyRule checks SSH key configuration
type IBMIsSSHKeyRule struct {
	tflint.DefaultRule
}

func NewIBMIsSSHKeyRule() *IBMIsSSHKeyRule {
	return &IBMIsSSHKeyRule{}
}

func (r *IBMIsSSHKeyRule) Name() string {
	return "ibm_is_ssh_key"
}

func (r *IBMIsSSHKeyRule) Enabled() bool {
	return true
}

func (r *IBMIsSSHKeyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSSHKeyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSSHKeyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_ssh_key", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "public_key"},
			{Name: "type"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "public_key"}
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

		// Validate public_key format
		if attr, exists := resource.Body.Attributes["public_key"]; exists {
			var publicKey string
			err := runner.EvaluateExpr(attr.Expr, &publicKey, nil)
			if err != nil {
				return err
			}

			if !isValidPublicKey(publicKey) {
				runner.EmitIssue(
					r,
					"invalid SSH public key format",
					attr.Expr.Range(),
				)
			}
		}

		// Validate key type
		if attr, exists := resource.Body.Attributes["type"]; exists {
			var keyType string
			err := runner.EvaluateExpr(attr.Expr, &keyType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"rsa":     true,
				"ed25519": true,
			}

			if !validTypes[keyType] {
				runner.EmitIssue(
					r,
					"type must be either 'rsa' or 'ed25519'",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}

func isValidPublicKey(key string) bool {
	// Basic validation of SSH public key format
	// This is a simplified check - implement more thorough validation if needed
	key = strings.TrimSpace(key)
	parts := strings.Fields(key)
	return len(parts) >= 2 && (strings.HasPrefix(parts[0], "ssh-rsa") || strings.HasPrefix(parts[0], "ssh-ed25519"))
}
