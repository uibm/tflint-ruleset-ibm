package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsDedicatedHostDiskManagementRule checks dedicated host disk management configuration
type IBMIsDedicatedHostDiskManagementRule struct {
	tflint.DefaultRule
}

func NewIBMIsDedicatedHostDiskManagementRule() *IBMIsDedicatedHostDiskManagementRule {
	return &IBMIsDedicatedHostDiskManagementRule{}
}

func (r *IBMIsDedicatedHostDiskManagementRule) Name() string {
	return "ibm_is_dedicated_host_disk_management"
}

func (r *IBMIsDedicatedHostDiskManagementRule) Enabled() bool {
	return true
}

func (r *IBMIsDedicatedHostDiskManagementRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsDedicatedHostDiskManagementRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsDedicatedHostDiskManagementRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_dedicated_host_disk_management", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "dedicated_host"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "disks",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
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
		// Check required attribute
		if _, exists := resource.Body.Attributes["dedicated_host"]; !exists {
			runner.EmitIssue(
				r,
				"dedicated_host attribute must be specified",
				resource.DefRange,
			)
		}

		// Check disks blocks
		if len(resource.Body.Blocks) == 0 {
			runner.EmitIssue(
				r,
				"at least one disks block must be specified",
				resource.DefRange,
			)
		}

		for _, block := range resource.Body.Blocks {
			if block.Type == "disks" {
				// Check required disk attributes
				requiredDiskAttrs := []string{"name", "id"}
				for _, attr := range requiredDiskAttrs {
					if _, exists := block.Body.Attributes[attr]; !exists {
						runner.EmitIssue(
							r,
							fmt.Sprintf("`%s` attribute must be specified in disks block", attr),
							block.DefRange,
						)
					}
				}

				// Validate disk name
				if attr, exists := block.Body.Attributes["name"]; exists {
					var name string
					err := runner.EvaluateExpr(attr.Expr, &name, nil)
					if err != nil {
						return err
					}

					if len(name) == 0 {
						runner.EmitIssue(
							r,
							"disk name cannot be empty",
							attr.Expr.Range(),
						)
					}
				}
			}
		}
	}

	return nil
}
