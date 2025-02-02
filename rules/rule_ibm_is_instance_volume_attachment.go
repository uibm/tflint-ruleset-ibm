package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceVolumeAttachmentRule checks instance volume attachment configuration
type IBMIsInstanceVolumeAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsInstanceVolumeAttachmentRule() *IBMIsInstanceVolumeAttachmentRule {
	return &IBMIsInstanceVolumeAttachmentRule{}
}

func (r *IBMIsInstanceVolumeAttachmentRule) Name() string {
	return "ibm_is_instance_volume_attachment"
}

func (r *IBMIsInstanceVolumeAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsInstanceVolumeAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsInstanceVolumeAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsInstanceVolumeAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_instance_volume_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "instance"},
			{Name: "name"},
			{Name: "volume"},
			{Name: "delete_volume_on_attachment_delete"},
			{Name: "delete_volume_on_instance_delete"},
			{Name: "profile"},
			{Name: "capacity"},
			{Name: "iops"},
			{Name: "snapshot"},
			{Name: "volume_name"},
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

		// Check that either volume or volume_name is specified
		hasVolume := false
		hasVolumeName := false
		if _, exists := resource.Body.Attributes["volume"]; exists {
			hasVolume = true
		}
		if _, exists := resource.Body.Attributes["volume_name"]; exists {
			hasVolumeName = true
		}

		if !hasVolume && !hasVolumeName {
			runner.EmitIssue(
				r,
				"either volume or volume_name must be specified",
				resource.DefRange,
			)
		}

		if hasVolume && hasVolumeName {
			runner.EmitIssue(
				r,
				"cannot specify both volume and volume_name",
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

		// Validate profile if specified
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			validProfiles := map[string]bool{
				"general-purpose": true,
				"custom":          true,
				"5iops-tier":      true,
				"10iops-tier":     true,
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid volume profile specified",
					attr.Expr.Range(),
				)
			}
		}

		// Validate capacity if specified
		if attr, exists := resource.Body.Attributes["capacity"]; exists {
			var capacity int
			err := runner.EvaluateExpr(attr.Expr, &capacity, nil)
			if err != nil {
				return err
			}

			if capacity < 10 || capacity > 2000 {
				runner.EmitIssue(
					r,
					"capacity must be between 10 and 2000 GB",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
