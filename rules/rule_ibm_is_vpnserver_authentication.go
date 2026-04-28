package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpnserverAuthenticationRule checks the configuration of ibm_is_vpnserver_authentication resource
type IBMVpnserverAuthenticationRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpnserverAuthenticationRule returns a new rule
func NewIBMVpnserverAuthenticationRule() *IBMVpnserverAuthenticationRule {
	return &IBMVpnserverAuthenticationRule{
		resourceType: "ibm_is_vpnserver_authentication",
	}
}

// Name returns the rule name
func (r *IBMVpnserverAuthenticationRule) Name() string {
	return "vpnserver_authentication"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpnserverAuthenticationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpnserverAuthenticationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpnserverAuthenticationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpnserverAuthenticationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "crl"},
			{Name: "method"},
			{Name: "identity_provider"},
			{Name: "client_ca"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["method"]; !exists {
			if err := runner.EmitIssue(
				r,
				"method attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["identity_provider"]; !exists {
			if err := runner.EmitIssue(
				r,
				"identity_provider attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["client_ca"]; !exists {
			if err := runner.EmitIssue(
				r,
				"client_ca attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["crl"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'crl' value must be at least 1 characters",
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
