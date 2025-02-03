package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPrivatePathServiceGatewayRevokeAccountRule checks private path service gateway revoke account configuration
type IBMIsPrivatePathServiceGatewayRevokeAccountRule struct {
	tflint.DefaultRule
}

func NewIBMIsPrivatePathServiceGatewayRevokeAccountRule() *IBMIsPrivatePathServiceGatewayRevokeAccountRule {
	return &IBMIsPrivatePathServiceGatewayRevokeAccountRule{}
}

func (r *IBMIsPrivatePathServiceGatewayRevokeAccountRule) Name() string {
	return "ibm_is_private_path_service_gateway_revoke_account"
}

func (r *IBMIsPrivatePathServiceGatewayRevokeAccountRule) Enabled() bool {
	return true
}

func (r *IBMIsPrivatePathServiceGatewayRevokeAccountRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPrivatePathServiceGatewayRevokeAccountRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPrivatePathServiceGatewayRevokeAccountRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_private_path_service_gateway_revoke_account", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "account"},
			{Name: "private_path_service_gateway"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"account", "private_path_service_gateway"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate account format
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
