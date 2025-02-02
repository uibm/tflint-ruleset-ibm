package ibm

// Config is the configuration for the IBM ruleset.
type Config struct {
	IBMCloudApiKey string `hclext:"ibmcloud_api_key,optional"`
	Region         string `hclext:"region,optional"`
}
