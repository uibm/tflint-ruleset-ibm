package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetReservedIPPatchRule checks subnet reserved IP patch configuration
type IBMIsSubnetReservedIPPatchRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetReservedIPPatchRule() *IBMIsSubnetReservedIPPatchRule {
	return &IBMIsSubnetReservedIPPatchRule{}
}

func (r *IBMIsSubnetReservedIPPatchRule) Name() string {
	return "ibm_is_subnet_reserved_ip_patch"
}

func (r *IBMIsSubnetReservedIPPatchRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetReservedIPPatchRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetReservedIPPatchRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetReservedIPPatchRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet_reserved_ip_patch", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "subnet"},
			{Name: "reserved_ip"},
			{Name: "name"},
			{Name: "auto_delete"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"subnet", "reserved_ip"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format if specified
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
