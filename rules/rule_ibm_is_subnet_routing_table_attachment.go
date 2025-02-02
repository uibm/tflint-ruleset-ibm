package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsSubnetRoutingTableAttachmentRule checks subnet routing table attachment configuration
type IBMIsSubnetRoutingTableAttachmentRule struct {
	tflint.DefaultRule
}

func NewIBMIsSubnetRoutingTableAttachmentRule() *IBMIsSubnetRoutingTableAttachmentRule {
	return &IBMIsSubnetRoutingTableAttachmentRule{}
}

func (r *IBMIsSubnetRoutingTableAttachmentRule) Name() string {
	return "ibm_is_subnet_routing_table_attachment"
}

func (r *IBMIsSubnetRoutingTableAttachmentRule) Enabled() bool {
	return true
}

func (r *IBMIsSubnetRoutingTableAttachmentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsSubnetRoutingTableAttachmentRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsSubnetRoutingTableAttachmentRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_subnet_routing_table_attachment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "subnet"},
			{Name: "routing_table"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"subnet", "routing_table"}
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
