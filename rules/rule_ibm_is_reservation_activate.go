package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsReservationActivateRule checks reservation activation
type IBMIsReservationActivateRule struct {
	tflint.DefaultRule
}

func NewIBMIsReservationActivateRule() *IBMIsReservationActivateRule {
	return &IBMIsReservationActivateRule{}
}

func (r *IBMIsReservationActivateRule) Name() string {
	return "ibm_is_reservation_activate"
}

func (r *IBMIsReservationActivateRule) Enabled() bool {
	return true
}

func (r *IBMIsReservationActivateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsReservationActivateRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsReservationActivateRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_reservation_activate", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "reservation"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["reservation"]; !exists {
			runner.EmitIssue(
				r,
				"reservation attribute must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
