// rules/rule_ibm_is_bare_metal_server.go
package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/uibm/tflint-ruleset-ibm/project"
)

// IBMIsBareMetalServerRule checks bare metal server configuration
type IBMIsBareMetalServerRule struct {
	tflint.DefaultRule
}

func NewIBMIsBareMetalServerRule() *IBMIsBareMetalServerRule {
	return &IBMIsBareMetalServerRule{}
}

func (r *IBMIsBareMetalServerRule) Name() string {
	return "ibm_is_bare_metal_server"
}

func (r *IBMIsBareMetalServerRule) Enabled() bool {
	return true
}

func (r *IBMIsBareMetalServerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *IBMIsBareMetalServerRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *IBMIsBareMetalServerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("ibm_is_bare_metal_server", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
			{Name: "profile"},
			{Name: "image"},
			{Name: "zone"},
			{Name: "vpc"},
			{Name: "keys"},
		},
		Blocks: []hclext.BlockSchema{
			{
				Type: "primary_network_interface",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "subnet"},
						{Name: "allow_ip_spoofing"},
						{Name: "security_groups"},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		// Check required attributes
		requiredAttrs := []string{"name", "profile", "image", "zone", "vpc", "keys"}
		for _, attr := range requiredAttrs {
			if _, exists := resource.Body.Attributes[attr]; !exists {
				runner.EmitIssue(
					r,
					fmt.Sprintf("`%s` attribute must be specified", attr),
					resource.DefRange,
				)
			}
		}

		// Validate profile format
		if attr, exists := resource.Body.Attributes["profile"]; exists {
			var profile string
			err := runner.EvaluateExpr(attr.Expr, &profile, nil)
			if err != nil {
				return err
			}

			// Example profiles: bx2-metal-192x768, cx2-metal-96x384
			validProfiles := map[string]bool{
				"bx2-metal-192x768": true,
				"cx2-metal-96x384":  true,
				// Add other valid profiles
			}

			if !validProfiles[profile] {
				runner.EmitIssue(
					r,
					"invalid bare metal server profile specified",
					attr.Expr.Range(),
				)
			}
		}

		// Check primary network interface block
		hasPrimaryInterface := false
		for _, block := range resource.Body.Blocks {
			if block.Type == "primary_network_interface" {
				hasPrimaryInterface = true
				if _, exists := block.Body.Attributes["subnet"]; !exists {
					runner.EmitIssue(
						r,
						"subnet must be specified in primary_network_interface block",
						block.DefRange,
					)
				}
				break
			}
		}

		if !hasPrimaryInterface {
			runner.EmitIssue(
				r,
				"primary_network_interface block must be specified",
				resource.DefRange,
			)
		}
	}

	return nil
}
