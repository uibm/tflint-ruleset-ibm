package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceNetworkAttachmentRule checks instance network attachment configuration
type IBMIsInstanceNetworkAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceNetworkAttachmentRule() *IBMIsInstanceNetworkAttachmentRule {
	return &IBMIsInstanceNetworkAttachmentRule{}
}

func (r *IBMIsInstanceNetworkAttachmentRule) Name() string {
	return "ibm_is_instance_network_attachment"
}

func (r *IBMIsInstanceNetworkAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceNetworkAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceNetworkAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceNetworkAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_network_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "instance"},
			{Name: "name"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "virtual_network_interface",
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
		requiredAttrs := []string{"instance", "name"}
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

		// Check virtual_network_interface block
		hasVNI := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "virtual_network_interface" {
				hasVNI = true
				if _, exists := block.Body.Attributes["id"]; !exists {
					runner.EmitIssue(
						r,
						"id attribute must be specified in virtual_network_interface block",
						block.DefRange,
					)
				}
			}
		}

		if !hasVNI {
			runner.EmitIssue(
				r,
				"virtual_network_interface block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
