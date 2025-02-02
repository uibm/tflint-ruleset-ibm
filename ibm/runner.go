package ibm

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Runner is a wrapper of the RPC client for IBM Cloud.
type Runner struct {
	tflint.Runner
	PluginConfig *Config
	ibmClient    Client
}

// NewRunner returns a custom IBM Cloud runner.
func NewRunner(runner tflint.Runner, config *Config) (*Runner, error) {
	var client Client
	var err error

	if config != nil {
		// Create a Credentials object from the config
		creds := Credentials{
			APIKey: config.IBMCloudApiKey,
			Region: config.Region,
		}
		client, err = NewClient(creds)
		if err != nil {
			return nil, err
		}
	}

	return &Runner{
		Runner:       runner,
		PluginConfig: config,
		ibmClient:    client,
	}, nil
}

// IBMClient returns the IBM Cloud client.
func (r *Runner) NewIBMClient() Client { // Use IBMClient() as the method name
	return r.ibmClient
}

// EachStringSliceExprs iterates an evaluated value and the corresponding expression
// If the given expression is a static list, get an expression for each value
// If not, the given expression is used as it is
func (r *Runner) EachStringSliceExprs(expr hcl.Expression, proc func(val string, expr hcl.Expression)) error {
	var vals []string
	err := r.EvaluateExpr(expr, func(v []string) error {
		vals = v
		return nil
	}, nil)
	if err != nil {
		return err
	}

	exprs, diags := hcl.ExprList(expr)
	if diags.HasErrors() {
		logger.Debug("Expr is not static list: %s", diags)
		for range vals {
			exprs = append(exprs, expr)
		}
	}

	for idx, val := range vals {
		proc(val, exprs[idx])
	}
	return nil
}
