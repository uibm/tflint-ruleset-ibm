package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsShareDeleteAccessorBindingRule checks share delete accessor binding configuration
type IBMIsShareDeleteAccessorBindingRule struct {
	tflint.DefaultRule
}

func NewIBMIsShareDeleteAccessorBindingRule() *IBMIsShareDeleteAccessorBindingRule {
	return &IBMIsShareDeleteAccessorBindingRule{}
}

func (r *IBMIsShareDeleteAccessorBindingRule) Name() string {
	return "ibm_is_share_delete_accessor_binding"
}

func (r *IBMIsShareDeleteAccessorBindingRule) Enabled() bool {
	return true
}

func (r *IBMIsShareDeleteAccessorBindingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsShareDeleteAccessorBindingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsShareDeleteAccessorBindingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_share_delete_accessor_binding", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "share"},
			{Name: "accessor_binding"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"share", "accessor_binding"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate accessor_binding format
		if attr, exists := resource.Body.Attributes["accessor_binding"]; exists {
			var binding string
			err := runner.EvaluateExpr(attr.Expr, &binding, nil)
			if err != nil {
				return err
			}

			if len(binding) == 0 {
				runner.EmitIssue(
					r,
					"accessor_binding cannot be empty",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
