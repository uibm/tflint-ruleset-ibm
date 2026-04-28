package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpnserverAuthenticationByCertificateRule checks the configuration of ibm_is_vpnserver_authentication_by_certificate resource
type IBMVpnserverAuthenticationByCertificateRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpnserverAuthenticationByCertificateRule returns a new rule
func NewIBMVpnserverAuthenticationByCertificateRule() *IBMVpnserverAuthenticationByCertificateRule {
	return &IBMVpnserverAuthenticationByCertificateRule{
		resourceType: "ibm_is_vpnserver_authentication_by_certificate",
	}
}

// Name returns the rule name
func (r *IBMVpnserverAuthenticationByCertificateRule) Name() string {
	return "vpnserver_authentication_by_certificate"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpnserverAuthenticationByCertificateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpnserverAuthenticationByCertificateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpnserverAuthenticationByCertificateRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpnserverAuthenticationByCertificateRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "client_ca"},
			{Name: "crl"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
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
