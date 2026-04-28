package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpnserverAuthenticationByUsernameRule checks the configuration of ibm_is_vpnserver_authentication_by_username resource
type IBMVpnserverAuthenticationByUsernameRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpnserverAuthenticationByUsernameRule returns a new rule
func NewIBMVpnserverAuthenticationByUsernameRule() *IBMVpnserverAuthenticationByUsernameRule {
	return &IBMVpnserverAuthenticationByUsernameRule{
		resourceType: "ibm_is_vpnserver_authentication_by_username",
	}
}

// Name returns the rule name
func (r *IBMVpnserverAuthenticationByUsernameRule) Name() string {
	return "vpnserver_authentication_by_username"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpnserverAuthenticationByUsernameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpnserverAuthenticationByUsernameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpnserverAuthenticationByUsernameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpnserverAuthenticationByUsernameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "identity_provider"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["identity_provider"]; !exists {
			if err := runner.EmitIssue(
				r,
				"identity_provider attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
