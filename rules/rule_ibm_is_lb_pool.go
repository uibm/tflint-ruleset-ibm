package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsLBPoolRule checks load balancer pool configuration
type IBMIsLBPoolRule struct {
	tflint.DefaultRule
}

func NewIBMIsLBPoolRule() *IBMIsLBPoolRule {
	return &IBMIsLBPoolRule{}
}

func (r *IBMIsLBPoolRule) Name() string {
	return "ibm_is_lb_pool"
}

func (r *IBMIsLBPoolRule) Enabled() bool {
	return true
}

func (r *IBMIsLBPoolRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsLBPoolRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsLBPoolRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_lb_pool", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "lb"},
			{Name: "algorithm"},
			{Name: "protocol"},
			{Name: "health_delay"},
			{Name: "health_retries"},
			{Name: "health_timeout"},
			{Name: "health_type"},
			{Name: "proxy_protocol"},
			{Name: "session_persistence_type"},
			{Name: "session_persistence_app_cookie_name"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "lb", "algorithm", "protocol"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate algorithm
		if attr, exists := resource.Body.Attributes["algorithm"]; exists {
			var algorithm string
			err := runner.EvaluateExpr(attr.Expr, &algorithm, nil)
			if err != nil {
				return err
			}

			validAlgorithms := map[string]bool{
				"round_robin":          true,
				"weighted_round_robin": true,
				"least_connections":    true,
			}

			if !validAlgorithms[algorithm] {
				runner.EmitIssue(
					r,
					"algorithm must be one of: round_robin, weighted_round_robin, least_connections",
					attr.Expr.Range(),
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

		// Validate session persistence
		if attr, exists := resource.Body.Attributes["session_persistence_type"]; exists {
			var persistenceType string
			err := runner.EvaluateExpr(attr.Expr, &persistenceType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"source_ip":   true,
				"http_cookie": true,
				"app_cookie":  true,
			}

			if !validTypes[persistenceType] {
				runner.EmitIssue(
					r,
					"session_persistence_type must be one of: source_ip, http_cookie, app_cookie",
					attr.Expr.Range(),
				)
			}

			// Check if app_cookie_name is specified when type is app_cookie
			if persistenceType == "app_cookie" {
				if _, exists := resource.Body.Attributes["session_persistence_app_cookie_name"]; !exists {
					runner.EmitIssue(
						r,
						"session_persistence_app_cookie_name must be specified when session_persistence_type is app_cookie",
						attr.Expr.Range(),
					)
				}
			}
		}

		// Validate health check parameters
		if err := r.validateHealthCheck(runner, resource); err != nil {
			return err
		}
	}

	return nil
}

func (r *IBMIsLBPoolRule) validateHealthCheck(runner tflint.Runner, resource *hclext.Block) error {
	healthAttrs := []string{"health_delay", "health_retries", "health_timeout", "health_type"}

	// Check if any health check attribute is specified
	hasHealthCheck := false
	for _, attr := range healthAttrs {
		if _, exists := resource.Body.Attributes[attr]; exists {
			hasHealthCheck = true
			break
		}
	}

	// If health check is configured, all attributes must be specified
	if hasHealthCheck {
		for _, attr := range healthAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` must be specified when configuring health check", attr),
					resource.DefRange,
				)
			}
		}

		// Validate health_type
		if attr, exists := resource.Body.Attributes["health_type"]; exists {
			var healthType string
			err := runner.EvaluateExpr(attr.Expr, &healthType, nil)
			if err != nil {
				return err
			}

			validTypes := map[string]bool{
				"http":  true,
				"https": true,
				"tcp":   true,
			}

			if !validTypes[healthType] {
				runner.EmitIssue(
					r,
					"health_type must be one of: http, https, tcp",
					attr.Expr.Range(),
				)
			}
		}

		// Validate health check numeric values
		if attr, exists := resource.Body.Attributes["health_delay"]; exists {
			var delay int
			if err := runner.EvaluateExpr(attr.Expr, &delay, nil); err == nil {
				if delay < 1 || delay > 300 {
					runner.EmitIssue(
						r,
						"health_delay must be between 1 and 300 seconds",
						attr.Expr.Range(),
					)
				}
			}
		}

		if attr, exists := resource.Body.Attributes["health_retries"]; exists {
			var retries int
			if err := runner.EvaluateExpr(attr.Expr, &retries, nil); err == nil {
				if retries < 1 || retries > 10 {
					runner.EmitIssue(
						r,
						"health_retries must be between 1 and 10",
						attr.Expr.Range(),
					)
				}
			}
		}

		if attr, exists := resource.Body.Attributes["health_timeout"]; exists {
			var timeout int
			if err := runner.EvaluateExpr(attr.Expr, &timeout, nil); err == nil {
				if timeout < 1 || timeout > 120 {
					runner.EmitIssue(
						r,
						"health_timeout must be between 1 and 120 seconds",
						attr.Expr.Range(),
					)
				}
			}
		}
	}

	return nil
}
