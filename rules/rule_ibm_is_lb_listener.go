package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsLBListenerRule checks load balancer listener configuration
type IBMIsLBListenerRule struct {
	tflint.DefaultRule
}

func NewIBMIsLBListenerRule() *IBMIsLBListenerRule {
	return &IBMIsLBListenerRule{}
}

func (r *IBMIsLBListenerRule) Name() string {
	return "ibm_is_lb_listener"
}

func (r *IBMIsLBListenerRule) Enabled() bool {
	return true
}

func (r *IBMIsLBListenerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsLBListenerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsLBListenerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_lb_listener", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "lb"},
			{Name: "port"},
			{Name: "protocol"},
			{Name: "default_pool"},
			{Name: "certificate_instance"},
			{Name: "connection_limit"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"lb", "port", "protocol"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate protocol
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var protocol string
			err := runner.EvaluateExpr(attr.Expr, &protocol, nil)
			if err != nil {
				return err
			}

			validProtocols := map[string]bool{
				"http":  true,
				"https": true,
				"tcp":   true,
				"udp":   true,
			}

			if !validProtocols[protocol] {
				runner.EmitIssue(
					r,
					"protocol must be one of: http, https, tcp, udp",
					attr.Expr.Range(),
				)
			}
		}

		// Validate port
		if attr, exists := resource.Body.Attributes["port"]; exists {
			var port int
			err := runner.EvaluateExpr(attr.Expr, &port, nil)
			if err != nil {
				return err
			}

			if port < 1 || port > 65535 {
				runner.EmitIssue(
					r,
					"port must be between 1 and 65535",
					attr.Expr.Range(),
				)
			}
		}

		// Check certificate_instance for HTTPS protocol
		if attr, exists := resource.Body.Attributes["protocol"]; exists {
			var protocol string
			err := runner.EvaluateExpr(attr.Expr, &protocol, nil)
			if err != nil {
				return err
			}

			if protocol == "https" {
				if _, exists := resource.Body.Attributes["certificate_instance"]; !exists {
					runner.EmitIssue(
						r,
						"certificate_instance must be specified when protocol is https",
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	return nil
}
