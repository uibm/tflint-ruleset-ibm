package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerPoolMemberTargetRule checks the configuration of ibm_is_load_balancer_pool_member_target resource
type IBMLoadBalancerPoolMemberTargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerPoolMemberTargetRule returns a new rule
func NewIBMLoadBalancerPoolMemberTargetRule() *IBMLoadBalancerPoolMemberTargetRule {
	return &IBMLoadBalancerPoolMemberTargetRule{
		resourceType: "ibm_is_load_balancer_pool_member_target",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerPoolMemberTargetRule) Name() string {
	return "load_balancer_pool_member_target"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerPoolMemberTargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerPoolMemberTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerPoolMemberTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerPoolMemberTargetRule) Check(runner tflint.Runner) error {
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
