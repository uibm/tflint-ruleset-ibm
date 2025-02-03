package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsPrivatePathServiceGatewayOperationsRule checks private path service gateway operations configuration
type IBMIsPrivatePathServiceGatewayOperationsRule struct {
	tflint.DefaultRule
}

func NewIBMIsPrivatePathServiceGatewayOperationsRule() *IBMIsPrivatePathServiceGatewayOperationsRule {
	return &IBMIsPrivatePathServiceGatewayOperationsRule{}
}

func (r *IBMIsPrivatePathServiceGatewayOperationsRule) Name() string {
	return "ibm_is_private_path_service_gateway_operations"
}

func (r *IBMIsPrivatePathServiceGatewayOperationsRule) Enabled() bool {
	return true
}

func (r *IBMIsPrivatePathServiceGatewayOperationsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsPrivatePathServiceGatewayOperationsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsPrivatePathServiceGatewayOperationsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_private_path_service_gateway_operations", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "published"},
			{Name: "private_path_service_gateway"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"published", "private_path_service_gateway"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate published flag
		if attr, exists := resource.Body.Attributes["published"]; exists {
			var published bool
			err := runner.EvaluateExpr(attr.Expr, &published, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
