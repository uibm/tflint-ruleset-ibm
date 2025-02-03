package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBackupPolicyPlanRule checks backup policy plan configuration
type IBMIsBackupPolicyPlanRule struct {
	tflint.DefaultRule
}

func NewIBMIsBackupPolicyPlanRule() *IBMIsBackupPolicyPlanRule {
	return &IBMIsBackupPolicyPlanRule{}
}

func (r *IBMIsBackupPolicyPlanRule) Name() string {
	return "ibm_is_backup_policy_plan"
}

func (r *IBMIsBackupPolicyPlanRule) Enabled() bool {
	return true
}

func (r *IBMIsBackupPolicyPlanRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBackupPolicyPlanRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBackupPolicyPlanRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_backup_policy_plan", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "backup_policy_id"},
			{Name: "cron_spec"},
			{Name: "name"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "clone_policy",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "zones"},
						{Name: "max_snapshots"},
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
		requiredAttrs := []string{"backup_policy_id", "cron_spec", "name"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate cron_spec format
		if attr, exists := resource.Body.Attributes["cron_spec"]; exists {
			var cronSpec string
			err := runner.EvaluateExpr(attr.Expr, &cronSpec, nil)
			if err != nil {
				return err
			}

			if !isValidCronSpec(cronSpec) {
				runner.EmitIssue(
					r,
					"invalid cron_spec format",
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

		// Validate clone_policy block if present
		for _, block := range resource.Body.Blocks {
			if block.Type == "clone_policy" {
				// Validate zones
				if attr, exists := block.Body.Attributes["zones"]; exists {
					var zones []string
					err := runner.EvaluateExpr(attr.Expr, &zones, nil)
					if err != nil {
						return err
					}

					// Check if zones array is empty
					if len(zones) == 0 {
						runner.EmitIssue(
							r,
							"zones cannot be empty in clone_policy",
							attr.Expr.Range(),
						)
					}

					// Validate each zone format
					for _, zone := range zones {
						if !isValidZone(zone) {
							runner.EmitIssue(
								r,
								"invalid zone format in clone_policy. Must be in format: region-number (e.g., us-south-1)",
								attr.Expr.Range(),
							)
						}
					}
				} else {
					runner.EmitIssue(
						r,
						"zones must be specified in clone_policy block",
						block.DefRange,
					)
				}

				// Validate max_snapshots
				if attr, exists := block.Body.Attributes["max_snapshots"]; exists {
					var maxSnapshots int
					err := runner.EvaluateExpr(attr.Expr, &maxSnapshots, nil)
					if err != nil {
						return err
					}

					if maxSnapshots < 1 {
						runner.EmitIssue(
							r,
							"max_snapshots must be greater than 0",
							attr.Expr.Range(),
						)
					}
				} else {
					runner.EmitIssue(
						r,
						"max_snapshots must be specified in clone_policy block",
						block.DefRange,
					)
				}
			}
		}
	}

	return nil
}
