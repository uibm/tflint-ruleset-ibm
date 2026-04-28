package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceDefaultTrustedProfileRule checks the configuration of ibm_is_instance_default_trusted_profile resource
type IBMInstanceDefaultTrustedProfileRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceDefaultTrustedProfileRule returns a new rule
func NewIBMInstanceDefaultTrustedProfileRule() *IBMInstanceDefaultTrustedProfileRule {
	return &IBMInstanceDefaultTrustedProfileRule{
		resourceType: "ibm_is_instance_default_trusted_profile",
	}
}

// Name returns the rule name
func (r *IBMInstanceDefaultTrustedProfileRule) Name() string {
	return "instance_default_trusted_profile"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceDefaultTrustedProfileRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceDefaultTrustedProfileRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceDefaultTrustedProfileRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceDefaultTrustedProfileRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "auto_link"},
			{Name: "target"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["target"]; !exists {
			if err := runner.EmitIssue(
				r,
				"target attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
