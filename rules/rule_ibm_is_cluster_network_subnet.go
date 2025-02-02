package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsClusterNetworkSubnetRule checks cluster network subnet configuration
type IBMIsClusterNetworkSubnetRule struct {
	tflint.DefaultRule
}

func NewIBMIsClusterNetworkSubnetRule() *IBMIsClusterNetworkSubnetRule {
	return &IBMIsClusterNetworkSubnetRule{}
}

func (r *IBMIsClusterNetworkSubnetRule) Name() string {
	return "ibm_is_cluster_network_subnet"
}

func (r *IBMIsClusterNetworkSubnetRule) Enabled() bool {
	return true
}

func (r *IBMIsClusterNetworkSubnetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsClusterNetworkSubnetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsClusterNetworkSubnetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_cluster_network_subnet", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "cluster_network_id"},
			{Name: "name"},
			{Name: "total_ipv4_address_count"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"cluster_network_id", "name", "total_ipv4_address_count"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate total_ipv4_address_count
		if attr, exists := resource.Body.Attributes["total_ipv4_address_count"]; exists {
			var count int
			err := runner.EvaluateExpr(attr.Expr, &count, nil)
			if err != nil {
				return err
			}

			// Check if count is a power of 2 and within valid range
			if count < 8 || count > 8192 || !isPowerOfTwo(count) {
				runner.EmitIssue(
					r,
					"total_ipv4_address_count must be a power of 2 between 8 and 8192",
					attr.Expr.Range(),
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}

// isPowerOfTwo checks if a number is a power of 2
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
