package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVirtualEndpointGatewayRule checks virtual endpoint gateway configuration
type IBMIsVirtualEndpointGatewayRule struct {
	tflint.DefaultRule
}

func NewIBMIsVirtualEndpointGatewayRule() *IBMIsVirtualEndpointGatewayRule {
	return &IBMIsVirtualEndpointGatewayRule{}
}

func (r *IBMIsVirtualEndpointGatewayRule) Name() string {
	return "ibm_is_virtual_endpoint_gateway"
}

func (r *IBMIsVirtualEndpointGatewayRule) Enabled() bool {
	return true
}

func (r *IBMIsVirtualEndpointGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVirtualEndpointGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVirtualEndpointGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_virtual_endpoint_gateway", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "vpc"},
			{Name: "resource_group"},
			{Name: "security_groups"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "target",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
						{Name: "resource_type"},
					},
				},
			},
			{
				Type: "ips",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "subnet"},
						{Name: "name"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "vpc"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate name format
		if attr, exists := resource.Body.Attributes["name"]; exists {
			var name string
			err := runner.EvaluateExpr(attr.Expr, &name, nil)
			if err != nil {
				return err
			}

			if len(name) == 0 {
				runner.EmitIssue(
					r,
					"name cannot be empty",
					attr.Expr.Range(),
				)
			}

			if len(name) > 63 {
				runner.EmitIssue(
					r,
					"name cannot be longer than 63 characters",
					attr.Expr.Range(),
				)
			}
		}

		// Check target block
		hasTarget := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "target" {
				hasTarget = true

				// Check required target attributes
				targetAttrs := []string{"name", "resource_type"}
				for _, attr := range targetAttrs {
					if _, exists := block.Body.Attributes[attr]; !exists {
						runner.EmitIssue(
							r,
							fmt.Sprintf("`%s` attribute must be specified in target block", attr),
							block.DefRange,
						)
					}
				}

				// Validate resource_type
				if attr, exists := block.Body.Attributes["resource_type"]; exists {
					var resourceType string
					err := runner.EvaluateExpr(attr.Expr, &resourceType, nil)
					if err != nil {
						return err
					}

					validTypes := map[string]bool{
						"provider_cloud_service":          true,
						"provider_infrastructure_service": true,
					}

					if !validTypes[resourceType] {
						runner.EmitIssue(
							r,
							"resource_type must be either 'provider_cloud_service' or 'provider_infrastructure_service'",
							attr.Expr.Range(),
						)
					}
				}
			}
		}

		if !hasTarget {
			runner.EmitIssue(
				r,
				"target block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
