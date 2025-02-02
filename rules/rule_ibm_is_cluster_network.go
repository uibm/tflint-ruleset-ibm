package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsClusterNetworkRule checks cluster network configuration
type IBMIsClusterNetworkRule struct {
	tflint.DefaultRule
}

func NewIBMIsClusterNetworkRule() *IBMIsClusterNetworkRule {
	return &IBMIsClusterNetworkRule{}
}

func (r *IBMIsClusterNetworkRule) Name() string {
	return "ibm_is_cluster_network"
}

func (r *IBMIsClusterNetworkRule) Enabled() bool {
	return true
}

func (r *IBMIsClusterNetworkRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsClusterNetworkRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsClusterNetworkRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_cluster_network", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "profile"},
			{Name: "zone"},
			{Name: "resource_group"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "subnet_prefixes",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "cidr"},
					},
				},
			},
			{
				Type: "vpc",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "id"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "profile", "zone"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate profile
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			validProfiles := map[string]bool{
				"h100": true,
				// Add other valid profiles
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid cluster network profile specified",
					attr.Expr.Range(),
				)
			}
		}

		// Check subnet_prefixes block
		hasSubnetPrefixes := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "subnet_prefixes" {
				hasSubnetPrefixes = true
				if attr, exists := block.Body.Attributes["cidr"]; exists {
					var cidr string
					err := runner.EvaluateExpr(attr.Expr, &cidr, nil)
					if err != nil {
						return err
					}

					// Add CIDR validation if needed
					if !isValidCIDR(cidr) {
						runner.EmitIssue(
							r,
							"invalid CIDR format in subnet_prefixes",
							attr.Expr.Range(),
						)
					}
				}
			}
		}

		if !hasSubnetPrefixes {
			runner.EmitIssue(
				r,
				"subnet_prefixes block must be specified",
				resource.DefRange,
			)
		}

		// Check VPC block
		hasVPC := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "vpc" {
				hasVPC = true
				if _, exists := block.Body.Attributes["id"]; !exists {
					runner.EmitIssue(
						r,
						"vpc block must contain an id attribute",
						block.DefRange,
					)
				}
			}
		}

		if !hasVPC {
			runner.EmitIssue(
				r,
				"vpc block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}

func isValidCIDR(cidr string) bool {
	// Add CIDR validation logic here
	// This is a placeholder - implement proper CIDR validation
	return len(cidr) > 0
}
