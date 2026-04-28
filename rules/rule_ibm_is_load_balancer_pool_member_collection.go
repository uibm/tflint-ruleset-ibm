package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerPoolMemberCollectionRule checks the configuration of ibm_is_load_balancer_pool_member_collection resource
type IBMLoadBalancerPoolMemberCollectionRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerPoolMemberCollectionRule returns a new rule
func NewIBMLoadBalancerPoolMemberCollectionRule() *IBMLoadBalancerPoolMemberCollectionRule {
	return &IBMLoadBalancerPoolMemberCollectionRule{
		resourceType: "ibm_is_load_balancer_pool_member_collection",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerPoolMemberCollectionRule) Name() string {
	return "load_balancer_pool_member_collection"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerPoolMemberCollectionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerPoolMemberCollectionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerPoolMemberCollectionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerPoolMemberCollectionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "members"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["members"]; !exists {
			if err := runner.EmitIssue(
				r,
				"members attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
