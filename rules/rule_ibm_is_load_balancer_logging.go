package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerLoggingRule checks the configuration of ibm_is_load_balancer_logging resource
type IBMLoadBalancerLoggingRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerLoggingRule returns a new rule
func NewIBMLoadBalancerLoggingRule() *IBMLoadBalancerLoggingRule {
	return &IBMLoadBalancerLoggingRule{
		resourceType: "ibm_is_load_balancer_logging",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerLoggingRule) Name() string {
	return "load_balancer_logging"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerLoggingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerLoggingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerLoggingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerLoggingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "datapath"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
