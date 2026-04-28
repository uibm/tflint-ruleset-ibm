package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMClusterNetworkSubnetByTotalCountRule checks the configuration of ibm_is_cluster_network_subnet_by_total_count resource
type IBMClusterNetworkSubnetByTotalCountRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMClusterNetworkSubnetByTotalCountRule returns a new rule
func NewIBMClusterNetworkSubnetByTotalCountRule() *IBMClusterNetworkSubnetByTotalCountRule {
	return &IBMClusterNetworkSubnetByTotalCountRule{
		resourceType: "ibm_is_cluster_network_subnet_by_total_count",
	}
}

// Name returns the rule name
func (r *IBMClusterNetworkSubnetByTotalCountRule) Name() string {
	return "cluster_network_subnet_by_total_count"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMClusterNetworkSubnetByTotalCountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMClusterNetworkSubnetByTotalCountRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMClusterNetworkSubnetByTotalCountRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMClusterNetworkSubnetByTotalCountRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "total_ipv4_address_count"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["total_ipv4_address_count"]; !exists {
			if err := runner.EmitIssue(
				r,
				"total_ipv4_address_count attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
