package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsClusterNetworkInterfaceRule checks cluster network interface configuration
type IBMIsClusterNetworkInterfaceRule struct {
	tflint.DefaultRule
}

func NewIBMIsClusterNetworkInterfaceRule() *IBMIsClusterNetworkInterfaceRule {
	return &IBMIsClusterNetworkInterfaceRule{}
}

func (r *IBMIsClusterNetworkInterfaceRule) Name() string {
	return "ibm_is_cluster_network_interface"
}

func (r *IBMIsClusterNetworkInterfaceRule) Enabled() bool {
	return true
}

func (r *IBMIsClusterNetworkInterfaceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsClusterNetworkInterfaceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsClusterNetworkInterfaceRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_cluster_network_interface", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "cluster_network_id"},
			{Name: "name"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "primary_ip",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "id"},
					},
				},
			},
			{
				Type: "subnet",
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
		requiredAttrs := []string{"cluster_network_id", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Check primary_ip block
		hasPrimaryIP := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "primary_ip" {
				hasPrimaryIP = true
				if _, exists := block.Body.Attributes["id"]; !exists {
					runner.EmitIssue(
						r,
						"primary_ip block must contain an id attribute",
						block.DefRange,
					)
				}
			}
		}

		if !hasPrimaryIP {
			runner.EmitIssue(
				r,
				"primary_ip block must be specified",
				resource.DefRange,
			)
		}

		// Check subnet block
		hasSubnet := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "subnet" {
				hasSubnet = true
				if _, exists := block.Body.Attributes["id"]; !exists {
					runner.EmitIssue(
						r,
						"subnet block must contain an id attribute",
						block.DefRange,
					)
				}
			}
		}

		if !hasSubnet {
			runner.EmitIssue(
				r,
				"subnet block must be specified",
				resource.DefRange,
			)
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
