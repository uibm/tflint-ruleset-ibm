package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMReservedIptargetRule checks the configuration of ibm_is_reserved_iptarget resource
type IBMReservedIptargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMReservedIptargetRule returns a new rule
func NewIBMReservedIptargetRule() *IBMReservedIptargetRule {
	return &IBMReservedIptargetRule{
		resourceType: "ibm_is_reserved_iptarget",
	}
}

// Name returns the rule name
func (r *IBMReservedIptargetRule) Name() string {
	return "reserved_iptarget"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMReservedIptargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMReservedIptargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMReservedIptargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMReservedIptargetRule) Check(runner tflint.Runner) error {
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
