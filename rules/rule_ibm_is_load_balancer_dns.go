package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerDnsRule checks the configuration of ibm_is_load_balancer_dns resource
type IBMLoadBalancerDnsRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerDnsRule returns a new rule
func NewIBMLoadBalancerDnsRule() *IBMLoadBalancerDnsRule {
	return &IBMLoadBalancerDnsRule{
		resourceType: "ibm_is_load_balancer_dns",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerDnsRule) Name() string {
	return "load_balancer_dns"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerDnsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerDnsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerDnsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerDnsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "zone"},
			{Name: "instance"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["zone"]; !exists {
			if err := runner.EmitIssue(
				r,
				"zone attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["instance"]; !exists {
			if err := runner.EmitIssue(
				r,
				"instance attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
