package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionPolicyModeLocalRule checks the configuration of ibm_is_vpngateway_connection_policy_mode_local resource
type IBMVpngatewayConnectionPolicyModeLocalRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionPolicyModeLocalRule returns a new rule
func NewIBMVpngatewayConnectionPolicyModeLocalRule() *IBMVpngatewayConnectionPolicyModeLocalRule {
	return &IBMVpngatewayConnectionPolicyModeLocalRule{
		resourceType: "ibm_is_vpngateway_connection_policy_mode_local",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionPolicyModeLocalRule) Name() string {
	return "vpngateway_connection_policy_mode_local"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionPolicyModeLocalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionPolicyModeLocalRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionPolicyModeLocalRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionPolicyModeLocalRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "cidrs"},
			{Name: "ike_identities"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["cidrs"]; !exists {
			if err := runner.EmitIssue(
				r,
				"cidrs attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
