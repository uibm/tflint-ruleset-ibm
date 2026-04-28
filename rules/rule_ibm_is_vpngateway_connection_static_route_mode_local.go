package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionStaticRouteModeLocalRule checks the configuration of ibm_is_vpngateway_connection_static_route_mode_local resource
type IBMVpngatewayConnectionStaticRouteModeLocalRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionStaticRouteModeLocalRule returns a new rule
func NewIBMVpngatewayConnectionStaticRouteModeLocalRule() *IBMVpngatewayConnectionStaticRouteModeLocalRule {
	return &IBMVpngatewayConnectionStaticRouteModeLocalRule{
		resourceType: "ibm_is_vpngateway_connection_static_route_mode_local",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionStaticRouteModeLocalRule) Name() string {
	return "vpngateway_connection_static_route_mode_local"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionStaticRouteModeLocalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionStaticRouteModeLocalRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionStaticRouteModeLocalRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionStaticRouteModeLocalRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "ike_identities"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
