package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerListenerPolicyTargetRule checks the configuration of ibm_is_load_balancer_listener_policy_target resource
type IBMLoadBalancerListenerPolicyTargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerListenerPolicyTargetRule returns a new rule
func NewIBMLoadBalancerListenerPolicyTargetRule() *IBMLoadBalancerListenerPolicyTargetRule {
	return &IBMLoadBalancerListenerPolicyTargetRule{
		resourceType: "ibm_is_load_balancer_listener_policy_target",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerListenerPolicyTargetRule) Name() string {
	return "load_balancer_listener_policy_target"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerListenerPolicyTargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerListenerPolicyTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerListenerPolicyTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerListenerPolicyTargetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
