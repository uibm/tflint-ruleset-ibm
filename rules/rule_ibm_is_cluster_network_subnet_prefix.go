package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMClusterNetworkSubnetPrefixRule checks the configuration of ibm_is_cluster_network_subnet_prefix resource
type IBMClusterNetworkSubnetPrefixRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMClusterNetworkSubnetPrefixRule returns a new rule
func NewIBMClusterNetworkSubnetPrefixRule() *IBMClusterNetworkSubnetPrefixRule {
	return &IBMClusterNetworkSubnetPrefixRule{
		resourceType: "ibm_is_cluster_network_subnet_prefix",
	}
}

// Name returns the rule name
func (r *IBMClusterNetworkSubnetPrefixRule) Name() string {
	return "cluster_network_subnet_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMClusterNetworkSubnetPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMClusterNetworkSubnetPrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMClusterNetworkSubnetPrefixRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMClusterNetworkSubnetPrefixRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "cidr"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
