package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceDiskManagementRule checks instance disk management configuration
type IBMIsInstanceDiskManagementRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceDiskManagementRule() *IBMIsInstanceDiskManagementRule {
	return &IBMIsInstanceDiskManagementRule{}
}

func (r *IBMIsInstanceDiskManagementRule) Name() string {
	return "ibm_is_instance_disk_management"
}

func (r *IBMIsInstanceDiskManagementRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceDiskManagementRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceDiskManagementRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceDiskManagementRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_disk_management", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "instance"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "disks",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "id"},
						{Name: "name"},
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
		if _, exists := resource.Body.Attributes["instance"]; !exists {
			runner.EmitIssue(
				r,
				"instance attribute must be specified",
				resource.DefRange,
			)
		}

		// Check for at least one disks block
		hasDisks := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "disks" {
				hasDisks = true

				// Check required disk attributes
				requiredDiskAttrs := []string{"id"}
				for _, attr := range requiredDiskAttrs {
					if _, exists := block.Body.Attributes[attr]; !exists {
						runner.EmitIssue(
							r,
							fmt.Sprintf("`%s` attribute must be specified in disks block", attr),
							block.DefRange,
						)
					}
				}

				// Validate disk name if specified
				if attr, exists := block.Body.Attributes["name"]; exists {
					var name string
					err := runner.EvaluateExpr(attr.Expr, &name, nil)
					if err != nil {
						return err
					}

					if len(name) == 0 {
						runner.EmitIssue(
							r,
							"disk name cannot be empty when specified",
							attr.Expr.Range(),
						)
					}

					if len(name) > 63 {
						runner.EmitIssue(
							r,
							"disk name cannot be longer than 63 characters",
							attr.Expr.Range(),
						)
					}
				}
			}
		}

		if !hasDisks {
			runner.EmitIssue(
				r,
				"at least one disks block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
