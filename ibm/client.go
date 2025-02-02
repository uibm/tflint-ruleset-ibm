// ibm/client.go
package ibm

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	DefaultTimeout = 30 * time.Second
	RetryDelay     = 5 * time.Second
)

type IBMClient struct {
	VPC *vpcv1.VpcV1
}

type Credentials struct {
	APIKey         string
	Region         string
	Zone           string
	IAMAccessToken string
	Timeout        int
}

func NewClient(creds Credentials) (*IBMClient, error) {
	if creds.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	if creds.Region == "" {
		return nil, fmt.Errorf("region is required")
	}
	authenticator := &core.IamAuthenticator{
		ApiKey: creds.APIKey,
	}

	vpcOptions := &vpcv1.VpcV1Options{
		Authenticator: authenticator,
	}

	if creds.Region != "" {
		vpcOptions.URL = fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", creds.Region)
	}

	vpcClient, err := vpcv1.NewVpcV1(vpcOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create VPC client: %w", err)
	}

	if creds.Timeout == 0 {
		creds.Timeout = int(DefaultTimeout.Seconds())
	}

	vpcClient.Service.EnableRetries(creds.Timeout, RetryDelay)
	vpcClient.SetDefaultHeaders(http.Header{
		"User-Agent": []string{fmt.Sprintf("tflint-ruleset-ibm/%s", "0.1.0")},
	})

	return &IBMClient{
		VPC: vpcClient,
	}, nil
}

func (c *IBMClient) ValidateVPC(ctx context.Context, vpcID string) (bool, error) {
	options := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}

	_, response, err := c.VPC.GetVPCWithContext(ctx, options)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			return false, nil
		}
		return false, fmt.Errorf("failed to get VPC: %w", err)
	}

	return true, nil
}

func (c *IBMClient) ValidateInstanceProfile(ctx context.Context, profileName string) (bool, error) {
	options := &vpcv1.ListInstanceProfilesOptions{}
	result, _, err := c.VPC.ListInstanceProfilesWithContext(ctx, options)
	if err != nil {
		return false, fmt.Errorf("failed to list instance profiles: %w", err)
	}

	for _, profile := range result.Profiles {
		if *profile.Name == profileName {
			return true, nil
		}
	}

	return false, nil
}
func (c *IBMClient) ValidateBackupPolicy(ctx context.Context, policyID string) (bool, error) {
	options := &vpcv1.GetBackupPolicyOptions{
		ID: &policyID,
	}
	_, response, err := c.VPC.GetBackupPolicyWithContext(ctx, options)
	if err != nil {
		if response != nil && response.StatusCode == http.StatusNotFound {
			return false, nil
		}
		return false, fmt.Errorf("failed to get backup policy: %w", err)
	}
	return true, nil
}
