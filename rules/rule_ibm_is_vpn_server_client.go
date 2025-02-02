package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsVPNServerClientRule checks VPN server client configuration
type IBMIsVPNServerClientRule struct {
	tflint.DefaultRule
}

func NewIBMIsVPNServerClientRule() *IBMIsVPNServerClientRule {
	return &IBMIsVPNServerClientRule{}
}

func (r *IBMIsVPNServerClientRule) Name() string {
	return "ibm_is_vpn_server_client"
}

func (r *IBMIsVPNServerClientRule) Enabled() bool {
	return true
}

func (r *IBMIsVPNServerClientRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsVPNServerClientRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsVPNServerClientRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_vpn_server_client", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "vpn_server"},
			{Name: "vpn_client"},
			{Name: "delete"},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"vpn_server", "vpn_client"}
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
