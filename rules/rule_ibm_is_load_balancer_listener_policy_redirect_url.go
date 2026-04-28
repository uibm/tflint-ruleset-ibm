package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerListenerPolicyRedirectUrlRule checks the configuration of ibm_is_load_balancer_listener_policy_redirect_url resource
type IBMLoadBalancerListenerPolicyRedirectUrlRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerListenerPolicyRedirectUrlRule returns a new rule
func NewIBMLoadBalancerListenerPolicyRedirectUrlRule() *IBMLoadBalancerListenerPolicyRedirectUrlRule {
	return &IBMLoadBalancerListenerPolicyRedirectUrlRule{
		resourceType: "ibm_is_load_balancer_listener_policy_redirect_url",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerListenerPolicyRedirectUrlRule) Name() string {
	return "load_balancer_listener_policy_redirect_url"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerListenerPolicyRedirectUrlRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerListenerPolicyRedirectUrlRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerListenerPolicyRedirectUrlRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerListenerPolicyRedirectUrlRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "http_status_code"},
			{Name: "url"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["http_status_code"]; !exists {
			if err := runner.EmitIssue(
				r,
				"http_status_code attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["url"]; !exists {
			if err := runner.EmitIssue(
				r,
				"url attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["url"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 10 {
					if err := runner.EmitIssue(
						r,
						"'url' value must be at least 10 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["url"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 8000 {
					if err := runner.EmitIssue(
						r,
						"'url' value must not exceed 8000 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["url"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^(http(s)?|\\{protocol\\}):\\/\\/([^\\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for url: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'url' value must match pattern: ^(http(s)?|\\{protocol\\}):\\/\\/([^\\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$",
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
