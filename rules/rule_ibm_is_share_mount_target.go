package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsShareMountTargetRule checks share mount target configuration
type IBMIsShareMountTargetRule struct {
	tflint.DefaultRule
}

func NewIBMIsShareMountTargetRule() *IBMIsShareMountTargetRule {
	return &IBMIsShareMountTargetRule{}
}

func (r *IBMIsShareMountTargetRule) Name() string {
	return "ibm_is_share_mount_target"
}

func (r *IBMIsShareMountTargetRule) Enabled() bool {
	return true
}

func (r *IBMIsShareMountTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsShareMountTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsShareMountTargetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_share_mount_target", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "share"},
			{Name: "name"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "virtual_network_interface",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "subnet"},
						{Name: "name"},
					},
					Blocks: []hclext.BlockSchema{
						{
							Type: "primary_ip",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{
									{Name: "name"},
									{Name: "address"},
									{Name: "auto_delete"},
									{Name: "reserved_ip"},
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
		requiredAttrs := []string{"share", "name"}
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

				// Check required VNI attributes
				if _, exists := block.Body.Attributes["subnet"]; !exists {
					runner.EmitIssue(
						r,
						"subnet must be specified in virtual_network_interface block",
						block.DefRange,
					)
				}

				// Check primary_ip block
				hasPrimaryIP := false
				for _, ipBlock := range block.Body.Blocks {
					if ipBlock.Type == "primary_ip" {
						hasPrimaryIP = true

						// Validate primary IP configuration
						if err := r.validatePrimaryIP(runner, ipBlock); err != nil {
							return err
						}
					}
				}

				if !hasPrimaryIP {
					runner.EmitIssue(
						r,
						"primary_ip block must be specified in virtual_network_interface",
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

func (r *IBMIsShareMountTargetRule) validatePrimaryIP(runner tflint.Runner, block *hclext.Block) error {
	if attr, exists := block.Body.Attributes["name"]; exists {
		var name string
		err := runner.EvaluateExpr(attr.Expr, &name, nil)
		if err != nil {
			return err
		}

		if len(name) == 0 {
			runner.EmitIssue(
				r,
				"primary_ip name cannot be empty",
				attr.Expr.Range(),
			)
		}
	}

	// Add additional primary IP validations if needed
	return nil
}
