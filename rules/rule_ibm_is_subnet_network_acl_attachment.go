package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetNetworkACLAttachmentRule checks subnet Network ACL attachment configuration
type IBMIsSubnetNetworkACLAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetNetworkACLAttachmentRule() *IBMIsSubnetNetworkACLAttachmentRule {
	return &IBMIsSubnetNetworkACLAttachmentRule{}
}

func (r *IBMIsSubnetNetworkACLAttachmentRule) Name() string {
	return "ibm_is_subnet_network_acl_attachment"
}

func (r *IBMIsSubnetNetworkACLAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetNetworkACLAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetNetworkACLAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetNetworkACLAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet_network_acl_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "subnet"},
			{Name: "network_acl"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"subnet", "network_acl"}
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
