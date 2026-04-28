package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMFlowLogCollectorTargetRule checks the configuration of ibm_is_flow_log_collector_target resource
type IBMFlowLogCollectorTargetRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMFlowLogCollectorTargetRule returns a new rule
func NewIBMFlowLogCollectorTargetRule() *IBMFlowLogCollectorTargetRule {
	return &IBMFlowLogCollectorTargetRule{
		resourceType: "ibm_is_flow_log_collector_target",
	}
}

// Name returns the rule name
func (r *IBMFlowLogCollectorTargetRule) Name() string {
	return "flow_log_collector_target"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMFlowLogCollectorTargetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMFlowLogCollectorTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMFlowLogCollectorTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMFlowLogCollectorTargetRule) Check(runner tflint.Runner) error {
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
