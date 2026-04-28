package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIkepolicyRule checks the configuration of ibm_is_ikepolicy resource
type IBMIkepolicyRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMIkepolicyRule returns a new rule
func NewIBMIkepolicyRule() *IBMIkepolicyRule {
	return &IBMIkepolicyRule{
		resourceType: "ibm_is_ikepolicy",
	}
}

// Name returns the rule name
func (r *IBMIkepolicyRule) Name() string {
	return "ikepolicy"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIkepolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIkepolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIkepolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMIkepolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "encryption_algorithm"},
			{Name: "ike_version"},
			{Name: "key_lifetime"},
			{Name: "name"},
			{Name: "resource_group"},
			{Name: "authentication_algorithm"},
			{Name: "dh_group"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["authentication_algorithm"]; !exists {
			if err := runner.EmitIssue(
				r,
				"authentication_algorithm attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["dh_group"]; !exists {
			if err := runner.EmitIssue(
				r,
				"dh_group attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["encryption_algorithm"]; !exists {
			if err := runner.EmitIssue(
				r,
				"encryption_algorithm attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["ike_version"]; !exists {
			if err := runner.EmitIssue(
				r,
				"ike_version attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["authentication_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"sha256",
					"sha384",
					"sha512",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'authentication_algorithm' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["dh_group"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"14",
					"15",
					"16",
					"17",
					"18",
					"19",
					"20",
					"21",
					"22",
					"23",
					"24",
					"31",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'dh_group' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["encryption_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"aes128",
					"aes192",
					"aes256",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'encryption_algorithm' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["ike_version"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"1",
					"2",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'ike_version' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["encryption_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'encryption_algorithm' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["encryption_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'encryption_algorithm' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'name' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 63 {
					if err := runner.EmitIssue(
						r,
						"'name' value must not exceed 63 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["authentication_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'authentication_algorithm' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["authentication_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'authentication_algorithm' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["authentication_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for authentication_algorithm: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'authentication_algorithm' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["encryption_algorithm"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for encryption_algorithm: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'encryption_algorithm' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["key_lifetime"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1800 {
					if err := runner.EmitIssue(
						r,
						"'key_lifetime' value must be at least 1800",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["key_lifetime"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val > 86400 {
					if err := runner.EmitIssue(
						r,
						"'key_lifetime' value must not exceed 86400",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
