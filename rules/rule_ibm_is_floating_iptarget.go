package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMFloatingIptargetRule checks the configuration of ibm_is_floating_iptarget resource
type IBMFloatingIptargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMFloatingIptargetRule returns a new rule
func NewIBMFloatingIptargetRule() *IBMFloatingIptargetRule {
	return &IBMFloatingIptargetRule{
		resourceType: "ibm_is_floating_iptarget",
	}
}

// Name returns the rule name
func (r *IBMFloatingIptargetRule) Name() string {
	return "floating_iptarget"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMFloatingIptargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMFloatingIptargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMFloatingIptargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMFloatingIptargetRule) Check(runner tflint.Runner) error {
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
