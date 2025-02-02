package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsReservationRule checks reservation configuration
type IBMIsReservationRule struct {
	tflint.DefaultRule
}

func NewIBMIsReservationRule() *IBMIsReservationRule {
	return &IBMIsReservationRule{}
}

func (r *IBMIsReservationRule) Name() string {
	return "ibm_is_reservation"
}

func (r *IBMIsReservationRule) Enabled() bool {
	return true
}

func (r *IBMIsReservationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsReservationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsReservationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_reservation", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "zone"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "capacity",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "total"},
					},
				},
			},
			{
				Type: "committed_use",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "term"},
					},
				},
			},
			{
				Type: "profile",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "name"},
						{Name: "resource_type"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required blocks and attributes
		hasCapacity := false
		hasCommittedUse := false
		hasProfile := false

		for _, block := range resource.Body.Blocks {
			switch block.Type {
			case "capacity":
				hasCapacity = true
				if attr, exists := block.Body.Attributes["total"]; exists {
					var total int
					err := runner.EvaluateExpr(attr.Expr, &total, nil)
					if err != nil {
						return err
					}

					if total < 1 {
						runner.EmitIssue(
							r,
							"capacity total must be greater than 0",
							attr.Expr.Range(),
						)
					}
				}
			case "committed_use":
				hasCommittedUse = true
				if attr, exists := block.Body.Attributes["term"]; exists {
					var term string
					err := runner.EvaluateExpr(attr.Expr, &term, nil)
					if err != nil {
						return err
					}

					validTerms := map[string]bool{
						"one_year":   true,
						"three_year": true,
					}

					if !validTerms[term] {
						runner.EmitIssue(
							r,
							"term must be either 'one_year' or 'three_year'",
							attr.Expr.Range(),
						)
					}
				}
			case "profile":
				hasProfile = true
				if attr, exists := block.Body.Attributes["resource_type"]; exists {
					var resourceType string
					err := runner.EvaluateExpr(attr.Expr, &resourceType, nil)
					if err != nil {
						return err
					}

					validTypes := map[string]bool{
						"instance_profile": true,
						"volume_profile":   true,
					}

					if !validTypes[resourceType] {
						runner.EmitIssue(
							r,
							"resource_type must be either 'instance_profile' or 'volume_profile'",
							attr.Expr.Range(),
						)
					}
				}
			}
		}

		if !hasCapacity {
			runner.EmitIssue(
				r,
				"capacity block must be specified",
				resource.DefRange,
			)
		}

		if !hasCommittedUse {
			runner.EmitIssue(
				r,
				"committed_use block must be specified",
				resource.DefRange,
			)
		}

		if !hasProfile {
			runner.EmitIssue(
				r,
				"profile block must be specified",
				resource.DefRange,
			)
		}

		// Validate zone format
		if attr, exists := resource.Body.Attributes["zone"]; exists {
			var zone string
			err := runner.EvaluateExpr(attr.Expr, &zone, nil)
			if err != nil {
				return err
			}

			if !isValidZone(zone) {
				runner.EmitIssue(
					r,
					"invalid zone format. Must be in format: region-number (e.g., us-south-1)",
					attr.Expr.Range(),
				)
			}
		}
	}

	return nil
}
