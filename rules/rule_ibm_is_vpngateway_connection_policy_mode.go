package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionPolicyModeRule checks the configuration of ibm_is_vpngateway_connection_policy_mode resource
type IBMVpngatewayConnectionPolicyModeRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionPolicyModeRule returns a new rule
func NewIBMVpngatewayConnectionPolicyModeRule() *IBMVpngatewayConnectionPolicyModeRule {
	return &IBMVpngatewayConnectionPolicyModeRule{
		resourceType: "ibm_is_vpngateway_connection_policy_mode",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionPolicyModeRule) Name() string {
	return "vpngateway_connection_policy_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionPolicyModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionPolicyModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionPolicyModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionPolicyModeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "local"},
			{Name: "peer"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["local"]; !exists {
			if err := runner.EmitIssue(
				r,
				"local attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["peer"]; !exists {
			if err := runner.EmitIssue(
				r,
				"peer attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
