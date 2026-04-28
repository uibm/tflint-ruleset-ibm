package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMReservationCapacityRule checks the configuration of ibm_is_reservation_capacity resource
type IBMReservationCapacityRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMReservationCapacityRule returns a new rule
func NewIBMReservationCapacityRule() *IBMReservationCapacityRule {
	return &IBMReservationCapacityRule{
		resourceType: "ibm_is_reservation_capacity",
	}
}

// Name returns the rule name
func (r *IBMReservationCapacityRule) Name() string {
	return "reservation_capacity"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMReservationCapacityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMReservationCapacityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMReservationCapacityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMReservationCapacityRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "total"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["total"]; !exists {
			if err := runner.EmitIssue(
				r,
				"total attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if attr, exists := resource.Body.Attributes["total"]; exists {
			var val float64
			if err := runner.EvaluateExpr(attr.Expr, &val, nil); err == nil {
				if val < 1 {
					if err := runner.EmitIssue(
						r,
						"'total' value must be at least 1",
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
