package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMClusterNetworkSubnetByIpv4CidrblockRule checks the configuration of ibm_is_cluster_network_subnet_by_ipv4_cidrblock resource
type IBMClusterNetworkSubnetByIpv4CidrblockRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMClusterNetworkSubnetByIpv4CidrblockRule returns a new rule
func NewIBMClusterNetworkSubnetByIpv4CidrblockRule() *IBMClusterNetworkSubnetByIpv4CidrblockRule {
	return &IBMClusterNetworkSubnetByIpv4CidrblockRule{
		resourceType: "ibm_is_cluster_network_subnet_by_ipv4_cidrblock",
	}
}

// Name returns the rule name
func (r *IBMClusterNetworkSubnetByIpv4CidrblockRule) Name() string {
	return "cluster_network_subnet_by_ipv4_cidrblock"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMClusterNetworkSubnetByIpv4CidrblockRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMClusterNetworkSubnetByIpv4CidrblockRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMClusterNetworkSubnetByIpv4CidrblockRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMClusterNetworkSubnetByIpv4CidrblockRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "ipv4_cidr_block"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["ipv4_cidr_block"]; !exists {
			if err := runner.EmitIssue(
				r,
				"ipv4_cidr_block attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
