package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMDnsserverRule checks the configuration of ibm_is_dnsserver resource
type IBMDnsserverRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMDnsserverRule returns a new rule
func NewIBMDnsserverRule() *IBMDnsserverRule {
	return &IBMDnsserverRule{
		resourceType: "ibm_is_dnsserver",
	}
}

// Name returns the rule name
func (r *IBMDnsserverRule) Name() string {
	return "dnsserver"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMDnsserverRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMDnsserverRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMDnsserverRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMDnsserverRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "address"},
			{Name: "zone_affinity"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
