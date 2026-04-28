package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMVpngatewayConnectionStaticRouteModePeerRule checks the configuration of ibm_is_vpngateway_connection_static_route_mode_peer resource
type IBMVpngatewayConnectionStaticRouteModePeerRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMVpngatewayConnectionStaticRouteModePeerRule returns a new rule
func NewIBMVpngatewayConnectionStaticRouteModePeerRule() *IBMVpngatewayConnectionStaticRouteModePeerRule {
	return &IBMVpngatewayConnectionStaticRouteModePeerRule{
		resourceType: "ibm_is_vpngateway_connection_static_route_mode_peer",
	}
}

// Name returns the rule name
func (r *IBMVpngatewayConnectionStaticRouteModePeerRule) Name() string {
	return "vpngateway_connection_static_route_mode_peer"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMVpngatewayConnectionStaticRouteModePeerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMVpngatewayConnectionStaticRouteModePeerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMVpngatewayConnectionStaticRouteModePeerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMVpngatewayConnectionStaticRouteModePeerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "fqdn"},
			{Name: "ike_identity"},
			{Name: "address"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["fqdn"]; !exists {
			if err := runner.EmitIssue(
				r,
				"fqdn attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["address"]; !exists {
			if err := runner.EmitIssue(
				r,
				"address attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["address"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 7 {
					if err := runner.EmitIssue(
						r,
						"'address' value must be at least 7 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["address"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 15 {
					if err := runner.EmitIssue(
						r,
						"'address' value must not exceed 15 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["fqdn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 4 {
					if err := runner.EmitIssue(
						r,
						"'fqdn' value must be at least 4 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["fqdn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 255 {
					if err := runner.EmitIssue(
						r,
						"'fqdn' value must not exceed 255 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["address"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^(?!(0(\\.0){3})|(255(\\.255){3})|(22[4-9]|23[0-9](\\.[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]){3}))([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\.([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])){3}$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for address: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'address' value must match pattern: ^(?!(0(\\.0){3})|(255(\\.255){3})|(22[4-9]|23[0-9](\\.[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]){3}))([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\.([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])){3}$",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["fqdn"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^((?=[A-Za-z0-9-]{1,63}\\.)[A-Za-z0-9-]*\\.)+[A-Za-z]{2,63}\\.?$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for fqdn: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'fqdn' value must match pattern: ^((?=[A-Za-z0-9-]{1,63}\\.)[A-Za-z0-9-]*\\.)+[A-Za-z]{2,63}\\.?$",
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
