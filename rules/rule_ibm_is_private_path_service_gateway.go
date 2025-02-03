package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPrivatePathServiceGatewayRule checks private path service gateway configuration
type IBMIsPrivatePathServiceGatewayRule struct {
	tflint.DefaultRule
}

func NewIBMIsPrivatePathServiceGatewayRule() *IBMIsPrivatePathServiceGatewayRule {
	return &IBMIsPrivatePathServiceGatewayRule{}
}

func (r *IBMIsPrivatePathServiceGatewayRule) Name() string {
	return "ibm_is_private_path_service_gateway"
}

func (r *IBMIsPrivatePathServiceGatewayRule) Enabled() bool {
	return true
}

func (r *IBMIsPrivatePathServiceGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPrivatePathServiceGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPrivatePathServiceGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_private_path_service_gateway", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "default_access_policy"},
			{Name: "load_balancer"},
			{Name: "zonal_affinity"},
			{Name: "service_endpoints"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "load_balancer", "default_access_policy"}
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

		// Validate default_access_policy
		if attr, exists := resource.Body.Attributes["default_access_policy"]; exists {
			var policy string
			err := runner.EvaluateExpr(attr.Expr, &policy, nil)
			if err != nil {
				return err
			}

			validPolicies := map[string]bool{
				"permit": true,
				"deny":   true,
			}

			if !validPolicies[policy] {
				runner.EmitIssue(
					r,
					"default_access_policy must be either 'permit' or 'deny'",
					attr.Expr.Range(),
				)
			}
		}

		// Validate service_endpoints if specified
		if attr, exists := resource.Body.Attributes["service_endpoints"]; exists {
			var endpoints []string
			err := runner.EvaluateExpr(attr.Expr, &endpoints, nil)
			if err != nil {
				return err
			}

			if len(endpoints) == 0 {
				runner.EmitIssue(
					r,
					"service_endpoints cannot be empty when specified",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
