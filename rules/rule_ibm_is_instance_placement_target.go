package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstancePlacementTargetRule checks the configuration of ibm_is_instance_placement_target resource
type IBMInstancePlacementTargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstancePlacementTargetRule returns a new rule
func NewIBMInstancePlacementTargetRule() *IBMInstancePlacementTargetRule {
	return &IBMInstancePlacementTargetRule{
		resourceType: "ibm_is_instance_placement_target",
	}
}

// Name returns the rule name
func (r *IBMInstancePlacementTargetRule) Name() string {
	return "instance_placement_target"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstancePlacementTargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstancePlacementTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstancePlacementTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstancePlacementTargetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
	}

	return nil
}
