package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMInstanceCatalogOfferingRule checks the configuration of ibm_is_instance_catalog_offering resource
type IBMInstanceCatalogOfferingRule struct {
	tflint.DefaultRule
	resourceType string
}

// NewIBMInstanceCatalogOfferingRule returns a new rule
func NewIBMInstanceCatalogOfferingRule() *IBMInstanceCatalogOfferingRule {
	return &IBMInstanceCatalogOfferingRule{
		resourceType: "ibm_is_instance_catalog_offering",
	}
}

// Name returns the rule name
func (r *IBMInstanceCatalogOfferingRule) Name() string {
	return "instance_catalog_offering"
}

// Enabled returns whether the rule is enabled by default
func (r *IBMInstanceCatalogOfferingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *IBMInstanceCatalogOfferingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *IBMInstanceCatalogOfferingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the resource configuration
func (r *IBMInstanceCatalogOfferingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "version"},
			{Name: "plan"},
			{Name: "offering"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes["offering"]; !exists {
			if err := runner.EmitIssue(
				r,
				"offering attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
		if _, exists := resource.Body.Attributes["version"]; !exists {
			if err := runner.EmitIssue(
				r,
				"version attribute must be specified",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
