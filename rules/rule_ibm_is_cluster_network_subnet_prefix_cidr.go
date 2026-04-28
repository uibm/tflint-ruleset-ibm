package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMClusterNetworkSubnetPrefixCidrRule checks the configuration of ibm_is_cluster_network_subnet_prefix_cidr resource
type IBMClusterNetworkSubnetPrefixCidrRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMClusterNetworkSubnetPrefixCidrRule returns a new rule
func NewIBMClusterNetworkSubnetPrefixCidrRule() *IBMClusterNetworkSubnetPrefixCidrRule {
	return &IBMClusterNetworkSubnetPrefixCidrRule{
		resourceType: "ibm_is_cluster_network_subnet_prefix_cidr",
	}
}

// Name returns the rule name
func (r *IBMClusterNetworkSubnetPrefixCidrRule) Name() string {
	return "cluster_network_subnet_prefix_cidr"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMClusterNetworkSubnetPrefixCidrRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMClusterNetworkSubnetPrefixCidrRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMClusterNetworkSubnetPrefixCidrRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMClusterNetworkSubnetPrefixCidrRule) Check(runner tflint.Runner) error {
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
