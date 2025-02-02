package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetPublicGatewayAttachmentRule checks subnet public gateway attachment configuration
type IBMIsSubnetPublicGatewayAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetPublicGatewayAttachmentRule() *IBMIsSubnetPublicGatewayAttachmentRule {
	return &IBMIsSubnetPublicGatewayAttachmentRule{}
}

func (r *IBMIsSubnetPublicGatewayAttachmentRule) Name() string {
	return "ibm_is_subnet_public_gateway_attachment"
}

func (r *IBMIsSubnetPublicGatewayAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetPublicGatewayAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetPublicGatewayAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetPublicGatewayAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet_public_gateway_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "subnet"},
			{Name: "public_gateway"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"subnet", "public_gateway"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}
	}

	return nil
}
