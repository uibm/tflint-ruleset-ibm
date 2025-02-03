package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPrivatePathServiceGatewayAccountPolicyRule checks private path service gateway account policy configuration
type IBMIsPrivatePathServiceGatewayAccountPolicyRule struct {
	tflint.DefaultRule
}

func NewIBMIsPrivatePathServiceGatewayAccountPolicyRule() *IBMIsPrivatePathServiceGatewayAccountPolicyRule {
	return &IBMIsPrivatePathServiceGatewayAccountPolicyRule{}
}

func (r *IBMIsPrivatePathServiceGatewayAccountPolicyRule) Name() string {
	return "ibm_is_private_path_service_gateway_account_policy"
}

func (r *IBMIsPrivatePathServiceGatewayAccountPolicyRule) Enabled() bool {
	return true
}

func (r *IBMIsPrivatePathServiceGatewayAccountPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPrivatePathServiceGatewayAccountPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPrivatePathServiceGatewayAccountPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_private_path_service_gateway_account_policy", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "access_policy"},
			{Name: "account"},
			{Name: "private_path_service_gateway"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"access_policy", "account", "private_path_service_gateway"}
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

		// Validate account format (IBM Cloud account ID format)
		if attr, exists := resource.Body.Attributes["account"]; exists {
			var account string
			err := runner.EvaluateExpr(attr.Expr, &account, nil)
			if err != nil {
				return err
			}

			if len(account) != 32 { // IBM Cloud account IDs are typically 32 characters
				runner.EmitIssue(
					r,
					"invalid account ID format",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
