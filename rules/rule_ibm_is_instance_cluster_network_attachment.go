package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceClusterNetworkAttachmentRule checks instance cluster network attachment configuration
type IBMIsInstanceClusterNetworkAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceClusterNetworkAttachmentRule() *IBMIsInstanceClusterNetworkAttachmentRule {
	return &IBMIsInstanceClusterNetworkAttachmentRule{}
}

func (r *IBMIsInstanceClusterNetworkAttachmentRule) Name() string {
	return "ibm_is_instance_cluster_network_attachment"
}

func (r *IBMIsInstanceClusterNetworkAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceClusterNetworkAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceClusterNetworkAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceClusterNetworkAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_cluster_network_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "instance_id"},
			{Name: "name"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "before",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "id"},
					},
				},
			},
			{
				Type: "cluster_network_interface",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
					},
					Blocks: []hclext.BlockSchema{
						{
							Type: "subnet",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{
									{Name: "id"},
								},
							},
						},
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
		requiredAttrs := []string{"instance_id", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
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

		// Validate cluster_network_interface block
		hasInterface := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "cluster_network_interface" {
				hasInterface = true

				// Check required subnet block
				hasSubnet := false
				for _, subBlock := range block.Body.Blocks {
					if subBlock.Type == "subnet" {
						hasSubnet = true
						if _, exists := subBlock.Body.Attributes["id"]; !exists {
							runner.EmitIssue(
								r,
								"id attribute must be specified in subnet block",
								subBlock.DefRange,
							)
						}
					}
				}

				if !hasSubnet {
					runner.EmitIssue(
						r,
						"subnet block must be specified in cluster_network_interface",
						block.DefRange,
					)
				}
			}
		}

		if !hasInterface {
			runner.EmitIssue(
				r,
				"cluster_network_interface block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
