package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/uibm/tflint-ruleset-ibm/ibm"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: ibm.NewRuleSet(),
	})
}
