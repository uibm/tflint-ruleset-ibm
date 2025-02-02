package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsIPSecPolicyRule checks IPSec policy configuration
type IBMIsIPSecPolicyRule struct {
	tflint.DefaultRule
}

func NewIBMIsIPSecPolicyRule() *IBMIsIPSecPolicyRule {
	return &IBMIsIPSecPolicyRule{}
}

func (r *IBMIsIPSecPolicyRule) Name() string {
	return "ibm_is_ipsec_policy"
}

func (r *IBMIsIPSecPolicyRule) Enabled() bool {
	return true
}

func (r *IBMIsIPSecPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsIPSecPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsIPSecPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_ipsec_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "authentication_algorithm"},
			{Name: "encryption_algorithm"},
			{Name: "pfs"},
			{Name: "key_lifetime"},
			{Name: "resource_group"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{
			"name",
			"authentication_algorithm",
			"encryption_algorithm",
			"pfs",
		}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate authentication_algorithm
		if attr, exists := resource.Body.Attributes["authentication_algorithm"]; exists {
			var authAlg string
			err := runner.EvaluateExpr(attr.Expr, &authAlg, nil)
			if err != nil {
				return err
			}

			validAuthAlgs := map[string]bool{
				"md5":    true,
				"sha1":   true,
				"sha256": true,
				"sha384": true,
				"sha512": true,
			}

			if !validAuthAlgs[authAlg] {
				runner.EmitIssue(
					r,
					"authentication_algorithm must be one of: md5, sha1, sha256, sha384, sha512",
					attr.Expr.Range(),
				)
			}
		}

		// Validate encryption_algorithm
		if attr, exists := resource.Body.Attributes["encryption_algorithm"]; exists {
			var encAlg string
			err := runner.EvaluateExpr(attr.Expr, &encAlg, nil)
			if err != nil {
				return err
			}

			validEncAlgs := map[string]bool{
				"triple_des": true,
				"aes128":     true,
				"aes192":     true,
				"aes256":     true,
			}

			if !validEncAlgs[encAlg] {
				runner.EmitIssue(
					r,
					"encryption_algorithm must be one of: triple_des, aes128, aes192, aes256",
					attr.Expr.Range(),
				)
			}
		}

		// Validate pfs
		if attr, exists := resource.Body.Attributes["pfs"]; exists {
			var pfs string
			err := runner.EvaluateExpr(attr.Expr, &pfs, nil)
			if err != nil {
				return err
			}

			validPFS := map[string]bool{
				"disabled": true,
				"group_2":  true,
				"group_5":  true,
				"group_14": true,
			}

			if !validPFS[pfs] {
				runner.EmitIssue(
					r,
					"pfs must be one of: disabled, group_2, group_5, group_14",
					attr.Expr.Range(),
				)
			}
		}

		// Validate key_lifetime if specified
		if attr, exists := resource.Body.Attributes["key_lifetime"]; exists {
			var lifetime int
			err := runner.EvaluateExpr(attr.Expr, &lifetime, nil)
			if err != nil {
				return err
			}

			if lifetime < 300 || lifetime > 86400 {
				runner.EmitIssue(
					r,
					"key_lifetime must be between 300 and 86400 seconds",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
