package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsShareReplicaOperationsRule checks share replica operations configuration
type IBMIsShareReplicaOperationsRule struct {
	tflint.DefaultRule
}

func NewIBMIsShareReplicaOperationsRule() *IBMIsShareReplicaOperationsRule {
	return &IBMIsShareReplicaOperationsRule{}
}

func (r *IBMIsShareReplicaOperationsRule) Name() string {
	return "ibm_is_share_replica_operations"
}

func (r *IBMIsShareReplicaOperationsRule) Enabled() bool {
	return true
}

func (r *IBMIsShareReplicaOperationsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsShareReplicaOperationsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsShareReplicaOperationsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_share_replica_operations", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "share_replica"},
			{Name: "split_share"},
			{Name: "fallback_policy"},
			{Name: "timeout"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		if _, exists := resource.Body.Attributes["share_replica"]; !exists {
			runner.EmitIssue(
				r,
				"share_replica attribute must be specified",
				resource.DefRange,
			)
		}

		// Validate operation type - either split_share or fallback_policy must be specified
		hasSplitShare := false
		hasFallbackPolicy := false

		if attr, exists := resource.Body.Attributes["split_share"]; exists {
			hasSplitShare = true
			var splitShare bool
			err := runner.EvaluateExpr(attr.Expr, &splitShare, nil)
			if err != nil {
				return err
			}
		}

		if attr, exists := resource.Body.Attributes["fallback_policy"]; exists {
			hasFallbackPolicy = true
			var fallbackPolicy string
			err := runner.EvaluateExpr(attr.Expr, &fallbackPolicy, nil)
			if err != nil {
				return err
			}

			validPolicies := map[string]bool{
				"split":    true,
				"failover": true,
			}

			if !validPolicies[fallbackPolicy] {
				runner.EmitIssue(
					r,
					"fallback_policy must be either 'split' or 'failover'",
					attr.Expr.Range(),
				)
			}
		}

		if !hasSplitShare && !hasFallbackPolicy {
			runner.EmitIssue(
				r,
				"either split_share or fallback_policy must be specified",
				resource.DefRange,
			)
		}

		if hasSplitShare && hasFallbackPolicy {
			runner.EmitIssue(
				r,
				"cannot specify both split_share and fallback_policy",
				resource.DefRange,
			)
		}

		// Validate timeout if specified
		if attr, exists := resource.Body.Attributes["timeout"]; exists {
			var timeout int
			err := runner.EvaluateExpr(attr.Expr, &timeout, nil)
			if err != nil {
				return err
			}

			if timeout < 0 {
				runner.EmitIssue(
					r,
					"timeout cannot be negative",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
