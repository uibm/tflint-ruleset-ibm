package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule checks endpoint gateway binding operations configuration
type IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule struct {
	tflint.DefaultRule
}

func NewIBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule() *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule {
	return &IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule{}
}

func (r *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule) Name() string {
	return "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations"
}

func (r *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule) Enabled() bool {
	return true
}

func (r *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPrivatePathServiceGatewayEndpointGatewayBindingOperationsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "access_policy"},
			{Name: "endpoint_gateway_binding"},
			{Name: "private_path_service_gateway"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{
			"access_policy",
			"endpoint_gateway_binding",
			"private_path_service_gateway",
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

		// Validate access_policy
		if attr, exists := resource.Body.Attributes["access_policy"]; exists {
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
					"access_policy must be either 'permit' or 'deny'",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
