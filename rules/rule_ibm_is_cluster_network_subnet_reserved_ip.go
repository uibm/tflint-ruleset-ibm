package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsClusterNetworkSubnetReservedIPRule checks cluster network subnet reserved IP configuration
type IBMIsClusterNetworkSubnetReservedIPRule struct {
	tflint.DefaultRule
}

func NewIBMIsClusterNetworkSubnetReservedIPRule() *IBMIsClusterNetworkSubnetReservedIPRule {
	return &IBMIsClusterNetworkSubnetReservedIPRule{}
}

func (r *IBMIsClusterNetworkSubnetReservedIPRule) Name() string {
	return "ibm_is_cluster_network_subnet_reserved_ip"
}

func (r *IBMIsClusterNetworkSubnetReservedIPRule) Enabled() bool {
	return true
}

func (r *IBMIsClusterNetworkSubnetReservedIPRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsClusterNetworkSubnetReservedIPRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsClusterNetworkSubnetReservedIPRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_cluster_network_subnet_reserved_ip", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "address"},
			{Name: "cluster_network_id"},
			{Name: "cluster_network_subnet_id"},
			{Name: "name"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{
			"address",
			"cluster_network_id",
			"cluster_network_subnet_id",
			"name",
		}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate address format
		if attr, exists := resource.Body.Attributes["address"]; exists {
			var address string
			err := runner.EvaluateExpr(attr.Expr, &address, nil)
			if err != nil {
				return err
			}

			if !isValidIPv4(address) {
				runner.EmitIssue(
					r,
					"invalid IPv4 address format",
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

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
