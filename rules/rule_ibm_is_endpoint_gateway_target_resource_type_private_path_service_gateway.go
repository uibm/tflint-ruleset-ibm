package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule checks the configuration of ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway resource
type IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule returns a new rule
func NewIBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule() *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule {
	return &IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule{
		resourceType: "ibm_is_endpoint_gateway_target_resource_type_private_path_service_gateway",
	}
}

// Name returns the rule name
func (r *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule) Name() string {
	return "endpoint_gateway_target_resource_type_private_path_service_gateway"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMEndpointGatewayTargetResourceTypePrivatePathServiceGatewayRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "crn"},
			{Name: "resource_type"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["crn"]; !exists {
			if err := runner.EmitIssue(
				r,
				"crn attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["resource_type"]; !exists {
			if err := runner.EmitIssue(
				r,
				"resource_type attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["resource_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				validValues := []string{
					"private_path_service_gateway",
				}

				valid := false
				for _, v := range validValues {
					if v == val {
						valid = true
						break
					}
				}

				if !valid {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("'resource_type' value must be one of [%s]", strings.Join(validValues, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 17 {
					if err := runner.EmitIssue(
						r,
						"'crn' value must be at least 17 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 512 {
					if err := runner.EmitIssue(
						r,
						"'crn' value must not exceed 512 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["resource_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'resource_type' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["resource_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 128 {
					if err := runner.EmitIssue(
						r,
						"'resource_type' value must not exceed 128 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["crn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\\/]*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for crn: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'crn' value must match pattern: ^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\\/]*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["resource_type"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^[a-z][a-z0-9]*(_[a-z0-9]+)*$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for resource_type: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'resource_type' value must match pattern: ^[a-z][a-z0-9]*(_[a-z0-9]+)*$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
