package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMPrivatePathServiceGatewayRule checks the configuration of ibm_is_private_path_service_gateway resource
type IBMPrivatePathServiceGatewayRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMPrivatePathServiceGatewayRule returns a new rule
func NewIBMPrivatePathServiceGatewayRule() *IBMPrivatePathServiceGatewayRule {
	return &IBMPrivatePathServiceGatewayRule{
		resourceType: "ibm_is_private_path_service_gateway",
	}
}

// Name returns the rule name
func (r *IBMPrivatePathServiceGatewayRule) Name() string {
	return "private_path_service_gateway"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMPrivatePathServiceGatewayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMPrivatePathServiceGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMPrivatePathServiceGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMPrivatePathServiceGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "load_balancer"},
			{Name: "name"},
			{Name: "resource_group"},
			{Name: "service_endpoints"},
			{Name: "zonal_affinity"},
			{Name: "default_access_policy"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["load_balancer"]; !exists {
			if err := runner.EmitIssue(
				r,
				"load_balancer attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["service_endpoints"]; !exists {
			if err := runner.EmitIssue(
				r,
				"service_endpoints attribute must be specified",
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
