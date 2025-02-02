package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSecurityGroupTargetRule checks security group target configuration
type IBMIsSecurityGroupTargetRule struct {
	tflint.DefaultRule
}

func NewIBMIsSecurityGroupTargetRule() *IBMIsSecurityGroupTargetRule {
	return &IBMIsSecurityGroupTargetRule{}
}

func (r *IBMIsSecurityGroupTargetRule) Name() string {
	return "ibm_is_security_group_target"
}

func (r *IBMIsSecurityGroupTargetRule) Enabled() bool {
	return true
}

func (r *IBMIsSecurityGroupTargetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSecurityGroupTargetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSecurityGroupTargetRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_security_group_target", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "security_group"},
			{Name: "target"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"security_group", "target"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}
	}

	return nil
}
