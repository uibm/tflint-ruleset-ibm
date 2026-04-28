package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceClusterNetworkAttachmentRule checks the configuration of ibm_is_instance_cluster_network_attachment resource
type IBMInstanceClusterNetworkAttachmentRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceClusterNetworkAttachmentRule returns a new rule
func NewIBMInstanceClusterNetworkAttachmentRule() *IBMInstanceClusterNetworkAttachmentRule {
	return &IBMInstanceClusterNetworkAttachmentRule{
		resourceType: "ibm_is_instance_cluster_network_attachment",
	}
}

// Name returns the rule name
func (r *IBMInstanceClusterNetworkAttachmentRule) Name() string {
	return "instance_cluster_network_attachment"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceClusterNetworkAttachmentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceClusterNetworkAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceClusterNetworkAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceClusterNetworkAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "before"},
			{Name: "cluster_network_interface"},
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["cluster_network_interface"]; !exists {
			if err := runner.EmitIssue(
				r,
				"cluster_network_interface attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
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
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for name: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'name' value must match pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$",
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
