package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsInstanceRule checks the configuration of IBM Cloud VPC Instance resources
type IBMIsInstanceRule struct {
	tflint.DefaultRule
	resourceType    string
	attributeSchema []hclext.AttributeSchema
}

// NewIBMIsInstanceRule returns a new rule
func NewIBMIsInstanceRule() *IBMIsInstanceRule {
	return &IBMIsInstanceRule{
		resourceType: "ibm_is_instance",
		attributeSchema: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "profile"},
			{Name: "image"},
			{Name: "vpc"},
			{Name: "zone"},
			{Name: "keys"},
			{Name: "primary_network_interface"},
		},
	}
}

// Name returns the rule name
func (r *IBMIsInstanceRule) Name() string {
	return r.resourceType
}

// Enabled returns whether the rule is enabled by default
func (r *IBMIsInstanceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMIsInstanceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMIsInstanceRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check performs the check for this rule
func (r *IBMIsInstanceRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: r.attributeSchema,
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		if err := r.checkRequiredAttributes(runner, resource); err != nil {
			return err
		}

		// Validate profile if specified
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			if err := r.validateProfile(runner, attr); err != nil {
				return err
			}
		}

		// Validate image if specified
		if attr, exists := resource.Body.Attributes["image"]; exists {
			if err := r.validateImage(runner, attr); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *IBMIsInstanceRule) checkRequiredAttributes(runner tflint.Runner, resource *hclext.Block) error {
	requiredAttrs := []string{"name", "profile", "image", "vpc", "zone"}

	for _, attr := range requiredAttrs {
		if _, exists := resource.Body.Attributes[attr]; !exists {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` attribute must be specified", attr),
				resource.DefRange,
			)
		}
	}
	return nil
}

func (r *IBMIsInstanceRule) validateProfile(runner tflint.Runner, attr *hclext.Attribute) error {
	var profile string
	err := runner.EvaluateExpr(attr.Expr, &profile, nil)
	if err != nil {
		return err
	}

	if profile == "" {
		runner.EmitIssue(
			r,
			"`profile` attribute cannot be empty",
			attr.Expr.Range(),
		)
	}

	// Add more profile validation logic here
	// For example, check if the profile exists in IBM Cloud

	return nil
}

func (r *IBMIsInstanceRule) validateImage(runner tflint.Runner, attr *hclext.Attribute) error {
	var image string
	err := runner.EvaluateExpr(attr.Expr, &image, nil)
	if err != nil {
		return err
	}

	if image == "" {
		runner.EmitIssue(
			r,
			"`image` attribute cannot be empty",
			attr.Expr.Range(),
		)
	}

	// Add more image validation logic here
	// For example, check if the image exists in IBM Cloud

	return nil
}
