package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

type IBMIsVPCRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPCRule() *IBMIsVPCRule {
	return &IBMIsVPCRule{}
}

func (r *IBMIsVPCRule) Name() string {
	return "ibm_is_vpc"
}

func (r *IBMIsVPCRule) Enabled() bool {
	return true
}

func (r *IBMIsVPCRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPCRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPCRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpc", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "resource_group"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["name"]; !exists {
			runner.EmitIssue(
				r,
				"`name` attribute must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
