package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsImageObsoleteRule checks image obsolete configuration
type IBMIsImageObsoleteRule struct {
	tflint.DefaultRule
}

func NewIBMIsImageObsoleteRule() *IBMIsImageObsoleteRule {
	return &IBMIsImageObsoleteRule{}
}

func (r *IBMIsImageObsoleteRule) Name() string {
	return "ibm_is_image_obsolete"
}

func (r *IBMIsImageObsoleteRule) Enabled() bool {
	return true
}

func (r *IBMIsImageObsoleteRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsImageObsoleteRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsImageObsoleteRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_image_obsolete", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "image"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required image attribute
		if _, exists := resource.Body.Attributes["image"]; !exists {
			runner.EmitIssue(
				r,
				"image attribute must be specified",
				resource.DefRange,
			)
		}

		// Additional validation could be added here
		// For example:
		// - Verifying that the image exists
		// - Checking if the image is already obsolete
		// - Validating that the image is in a valid state for becoming obsolete
	}

	return nil
}
