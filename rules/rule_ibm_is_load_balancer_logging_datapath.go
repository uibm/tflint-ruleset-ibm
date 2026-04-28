package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerLoggingDatapathRule checks the configuration of ibm_is_load_balancer_logging_datapath resource
type IBMLoadBalancerLoggingDatapathRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerLoggingDatapathRule returns a new rule
func NewIBMLoadBalancerLoggingDatapathRule() *IBMLoadBalancerLoggingDatapathRule {
	return &IBMLoadBalancerLoggingDatapathRule{
		resourceType: "ibm_is_load_balancer_logging_datapath",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerLoggingDatapathRule) Name() string {
	return "load_balancer_logging_datapath"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerLoggingDatapathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerLoggingDatapathRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerLoggingDatapathRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerLoggingDatapathRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "active"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
