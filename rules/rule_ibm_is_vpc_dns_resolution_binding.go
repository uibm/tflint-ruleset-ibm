package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPCDNSResolutionBindingRule checks VPC DNS resolution binding configuration
type IBMIsVPCDNSResolutionBindingRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPCDNSResolutionBindingRule() *IBMIsVPCDNSResolutionBindingRule {
	return &IBMIsVPCDNSResolutionBindingRule{}
}

func (r *IBMIsVPCDNSResolutionBindingRule) Name() string {
	return "ibm_is_vpc_dns_resolution_binding"
}

func (r *IBMIsVPCDNSResolutionBindingRule) Enabled() bool {
	return true
}

func (r *IBMIsVPCDNSResolutionBindingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPCDNSResolutionBindingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPCDNSResolutionBindingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpc_dns_resolution_binding", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpc_id"},
		},
		Blocks: []hclext.BlockSchema{
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
		if _, exists := resource.Body.Attributes["name"]; !exists {
			runner.EmitIssue(
				r,
				"name attribute must be specified",
				resource.DefRange,
			)
		}

		// Check that either vpc_id or vpc block is specified
		hasVPCID := false
		hasVPCBlock := false

		if _, exists := resource.Body.Attributes["vpc_id"]; exists {
			hasVPCID = true
		}

		for _, block := range resource.Body.Blocks {
			if block.Type == "vpc" {
				hasVPCBlock = true
				if _, exists := block.Body.Attributes["id"]; !exists {
					runner.EmitIssue(
						r,
						"id attribute must be specified in vpc block",
						block.DefRange,
					)
				}
			}
		}

		if !hasVPCID && !hasVPCBlock {
			runner.EmitIssue(
				r,
				"either vpc_id attribute or vpc block must be specified",
				resource.DefRange,
			)
		}

		if hasVPCID && hasVPCBlock {
			runner.EmitIssue(
				r,
				"cannot specify both vpc_id attribute and vpc block",
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
