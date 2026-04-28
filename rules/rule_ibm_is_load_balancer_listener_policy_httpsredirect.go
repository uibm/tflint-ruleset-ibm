package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMLoadBalancerListenerPolicyHttpsredirectRule checks the configuration of ibm_is_load_balancer_listener_policy_httpsredirect resource
type IBMLoadBalancerListenerPolicyHttpsredirectRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMLoadBalancerListenerPolicyHttpsredirectRule returns a new rule
func NewIBMLoadBalancerListenerPolicyHttpsredirectRule() *IBMLoadBalancerListenerPolicyHttpsredirectRule {
	return &IBMLoadBalancerListenerPolicyHttpsredirectRule{
		resourceType: "ibm_is_load_balancer_listener_policy_httpsredirect",
	}
}

// Name returns the rule name
func (r *IBMLoadBalancerListenerPolicyHttpsredirectRule) Name() string {
	return "load_balancer_listener_policy_httpsredirect"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMLoadBalancerListenerPolicyHttpsredirectRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMLoadBalancerListenerPolicyHttpsredirectRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMLoadBalancerListenerPolicyHttpsredirectRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMLoadBalancerListenerPolicyHttpsredirectRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "uri"},
			{Name: "http_status_code"},
			{Name: "listener"},
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
		if _, exists := resource.Body.Attributes["listener"]; !exists {
			if err := runner.EmitIssue(
				r,
				"listener attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["uri"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) < 1 {
					if err := runner.EmitIssue(
						r,
						"'uri' value must be at least 1 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["uri"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if len(val) > 8192 {
					if err := runner.EmitIssue(
						r,
						"'uri' value must not exceed 8192 characters",
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
		if attr, exists := resource.Body.Attributes["uri"]; exists {
			var val string
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				pattern, err := regexp.Compile("^\\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\\?([a-zA-Z0-9-._~!$&'()*+,;=:@\\/?]|%[a-fA-F0-9]{2})*)?$")
				if err != nil {
					return fmt.Errorf("invalid regex pattern for uri: %v", err)
				}
				if !pattern.MatchString(val) {
					if err := runner.EmitIssue(
						r,
						"'uri' value must match pattern: ^\\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\\?([a-zA-Z0-9-._~!$&'()*+,;=:@\\/?]|%[a-fA-F0-9]{2})*)?$",
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
