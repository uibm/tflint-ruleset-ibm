package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsIKEPolicyRule checks IKE policy configuration
type IBMIsIKEPolicyRule struct {
	tflint.DefaultRule
}

func NewIBMIsIKEPolicyRule() *IBMIsIKEPolicyRule {
	return &IBMIsIKEPolicyRule{}
}

func (r *IBMIsIKEPolicyRule) Name() string {
	return "ibm_is_ike_policy"
}

func (r *IBMIsIKEPolicyRule) Enabled() bool {
	return true
}

func (r *IBMIsIKEPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsIKEPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsIKEPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_ike_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "authentication_algorithm"},
			{Name: "encryption_algorithm"},
			{Name: "dh_group"},
			{Name: "ike_version"},
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
			"dh_group",
			"ike_version",
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
			}

			if !validAuthAlgs[authAlg] {
				runner.EmitIssue(
					r,
					"authentication_algorithm must be one of: md5, sha1, sha256",
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

		// Validate dh_group
		if attr, exists := resource.Body.Attributes["dh_group"]; exists {
			var dhGroup int
			err := runner.EvaluateExpr(attr.Expr, &dhGroup, nil)
			if err != nil {
				return err
			}

			validDHGroups := map[int]bool{
				2:  true,
				5:  true,
				14: true,
				19: true,
			}

			if !validDHGroups[dhGroup] {
				runner.EmitIssue(
					r,
					"dh_group must be one of: 2, 5, 14, 19",
					attr.Expr.Range(),
				)
			}
		}

		// Validate ike_version
		if attr, exists := resource.Body.Attributes["ike_version"]; exists {
			var ikeVersion int
			err := runner.EvaluateExpr(attr.Expr, &ikeVersion, nil)
			if err != nil {
				return err
			}

			if ikeVersion != 1 && ikeVersion != 2 {
				runner.EmitIssue(
					r,
					"ike_version must be either 1 or 2",
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
