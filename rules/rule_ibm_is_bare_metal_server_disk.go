// rules/rule_ibm_is_bare_metal_server_disk.go
package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerDiskRule checks bare metal server disk configuration
type IBMIsBareMetalServerDiskRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerDiskRule() *IBMIsBareMetalServerDiskRule {
	return &IBMIsBareMetalServerDiskRule{}
}

func (r *IBMIsBareMetalServerDiskRule) Name() string {
	return "ibm_is_bare_metal_server_disk"
}

func (r *IBMIsBareMetalServerDiskRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerDiskRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerDiskRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerDiskRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server_disk", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "bare_metal_server"},
			{Name: "disk"},
			{Name: "name"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"bare_metal_server", "disk"}
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

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty when specified",
					attr.Expr.Range(),
				)
			}

			// Add additional name validation rules if needed
		}
	}

	return nil
}
