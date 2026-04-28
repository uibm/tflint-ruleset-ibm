package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceTemplateRule checks the configuration of ibm_is_instance_template resource
type IBMInstanceTemplateRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceTemplateRule returns a new rule
func NewIBMInstanceTemplateRule() *IBMInstanceTemplateRule {
	return &IBMInstanceTemplateRule{
		resourceType: "ibm_is_instance_template",
	}
}

// Name returns the rule name
func (r *IBMInstanceTemplateRule) Name() string {
	return "instance_template"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceTemplateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceTemplateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceTemplateRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceTemplateRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "enable_secure_boot"},
			{Name: "zone"},
			{Name: "default_trusted_profile"},
			{Name: "total_volume_bandwidth"},
			{Name: "catalog_offering"},
			{Name: "image"},
			{Name: "boot_volume_attachment"},
			{Name: "primary_network_interface"},
			{Name: "reservation_affinity"},
			{Name: "source_template"},
			{Name: "cluster_network_attachments"},
			{Name: "metadata_service"},
			{Name: "network_attachments"},
			{Name: "placement_target"},
			{Name: "resource_group"},
			{Name: "keys"},
			{Name: "primary_network_attachment"},
			{Name: "user_data"},
			{Name: "network_interfaces"},
			{Name: "availability_policy"},
			{Name: "name"},
			{Name: "profile"},
			{Name: "vpc"},
			{Name: "confidential_compute_mode"},
			{Name: "volume_attachments"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["zone"]; !exists {
			if err := runner.EmitIssue(
				r,
				"zone attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["catalog_offering"]; !exists {
			if err := runner.EmitIssue(
				r,
				"catalog_offering attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["image"]; !exists {
			if err := runner.EmitIssue(
				r,
				"image attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["primary_network_interface"]; !exists {
			if err := runner.EmitIssue(
				r,
				"primary_network_interface attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["source_template"]; !exists {
			if err := runner.EmitIssue(
				r,
				"source_template attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["boot_volume_attachment"]; !exists {
			if err := runner.EmitIssue(
				r,
				"boot_volume_attachment attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["primary_network_attachment"]; !exists {
			if err := runner.EmitIssue(
				r,
				"primary_network_attachment attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["confidential_compute_mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"disabled",
					"sgx",
					"tdx",
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
						fmt.Sprintf("'confidential_compute_mode' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["user_data"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 65535 {
					if err := runner.EmitIssue(
						r,
						"'user_data' value must not exceed 65535 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["confidential_compute_mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'confidential_compute_mode' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["confidential_compute_mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'confidential_compute_mode' value must not exceed 128 characters",
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
		if attr, exists := resource.Body.Attributes["user_data"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[\\s\\S]*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for user_data: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'user_data' value must match pattern: ^[\\s\\S]*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["confidential_compute_mode"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for confidential_compute_mode: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'confidential_compute_mode' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
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
	}

	return nil
}
