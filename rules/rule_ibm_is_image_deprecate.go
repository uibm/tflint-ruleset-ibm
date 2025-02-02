package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsImageDeprecateRule checks image deprecation configuration
type IBMIsImageDeprecateRule struct {
	tflint.DefaultRule
}

func NewIBMIsImageDeprecateRule() *IBMIsImageDeprecateRule {
	return &IBMIsImageDeprecateRule{}
}

func (r *IBMIsImageDeprecateRule) Name() string {
	return "ibm_is_image_deprecate"
}

func (r *IBMIsImageDeprecateRule) Enabled() bool {
	return true
}

func (r *IBMIsImageDeprecateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsImageDeprecateRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsImageDeprecateRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_image_deprecate", &hclext.BodySchema{
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
		// For example, checking if the image exists
	}

	return nil
}
