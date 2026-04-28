package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMSnapshotCloneRule checks the configuration of ibm_is_snapshot_clone resource
type IBMSnapshotCloneRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMSnapshotCloneRule returns a new rule
func NewIBMSnapshotCloneRule() *IBMSnapshotCloneRule {
	return &IBMSnapshotCloneRule{
		resourceType: "ibm_is_snapshot_clone",
	}
}

// Name returns the rule name
func (r *IBMSnapshotCloneRule) Name() string {
	return "snapshot_clone"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMSnapshotCloneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMSnapshotCloneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMSnapshotCloneRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMSnapshotCloneRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "zone"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["zone"]; !exists {
			if err := runner.EmitIssue(
				r,
				"zone attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
