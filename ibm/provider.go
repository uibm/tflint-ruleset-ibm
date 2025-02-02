package ibm

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// IBMProviderBlockSchema is a schema of the `ibm` provider block
var IBMProviderBlockSchema = &hclext.BodySchema{
	Attributes: []hclext.AttributeSchema{
		{Name: "ibmcloud_api_key"},
		{Name: "region"},
		// Add other attributes as needed (e.g., zone, iam_token, etc.)
	},
}

// GetCredentialsFromProvider retrieves credentials from the "provider" block.
func GetCredentialsFromProvider(runner tflint.Runner) (map[string]Config, error) {
	providers, err := runner.GetModuleContent(
		&hclext.BodySchema{
			Blocks: []hclext.BlockSchema{
				{
					Type:       "provider",
					LabelNames: []string{"name"},
					Body:       IBMProviderBlockSchema,
				},
			},
		},
		&tflint.GetModuleContentOption{ModuleCtx: tflint.RootModuleCtxType},
	)
	if err != nil {
		return nil, err
	}

	configs := map[string]Config{}

	for _, provider := range providers.Blocks {
		if provider.Labels[0] != "ibm" {
			continue
		}

		var config Config
		opts := &tflint.EvaluateExprOption{ModuleCtx: tflint.RootModuleCtxType}

		if attr, exists := provider.Body.Attributes["ibmcloud_api_key"]; exists {
			if err := runner.EvaluateExpr(attr.Expr, func(apiKey string) error {
				config.IBMCloudApiKey = apiKey
				return nil
			}, opts); err != nil {
				return nil, err
			}
		}

		if attr, exists := provider.Body.Attributes["region"]; exists {
			if err := runner.EvaluateExpr(attr.Expr, func(region string) error {
				config.Region = region
				return nil
			}, opts); err != nil {
				return nil, err
			}
		}

		// Add logic to handle aliases if needed

		configs["ibm"] = config // Or use alias if available
	}

	return configs, nil
}
