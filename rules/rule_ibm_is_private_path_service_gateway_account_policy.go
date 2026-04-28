package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMPrivatePathServiceGatewayAccountPolicyRule checks the configuration of ibm_is_private_path_service_gateway_account_policy resource
type IBMPrivatePathServiceGatewayAccountPolicyRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMPrivatePathServiceGatewayAccountPolicyRule returns a new rule
func NewIBMPrivatePathServiceGatewayAccountPolicyRule() *IBMPrivatePathServiceGatewayAccountPolicyRule {
	return &IBMPrivatePathServiceGatewayAccountPolicyRule{
		resourceType: "ibm_is_private_path_service_gateway_account_policy",
	}
}

// Name returns the rule name
func (r *IBMPrivatePathServiceGatewayAccountPolicyRule) Name() string {
	return "private_path_service_gateway_account_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMPrivatePathServiceGatewayAccountPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMPrivatePathServiceGatewayAccountPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMPrivatePathServiceGatewayAccountPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMPrivatePathServiceGatewayAccountPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "access_policy"},
			{Name: "account"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["access_policy"]; !exists {
			if err := runner.EmitIssue(
				r,
				"access_policy attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["account"]; !exists {
			if err := runner.EmitIssue(
				r,
				"account attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
