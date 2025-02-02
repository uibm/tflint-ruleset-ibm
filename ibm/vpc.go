package ibm

import (
	"context"
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// Client is an interface for the IBM Cloud API client.
type Client interface {
	GetInstanceProfiles() (map[string]bool, error)
	GetImages(region string) (map[string]bool, error)
	GetVPC(id string) (*vpcv1.VPC, error)
}

// GetInstanceProfiles is a wrapper to fetch instance profiles.
func (c *IBMClient) GetInstanceProfiles() (map[string]bool, error) {
	profiles := map[string]bool{}
	options := &vpcv1.ListInstanceProfilesOptions{}
	result, _, err := c.VPC.ListInstanceProfilesWithContext(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to list instance profiles: %w", err)
	}
	for _, profile := range result.Profiles {
		profiles[*profile.Name] = true
	}
	return profiles, nil
}

// GetImages fetches images available in a specific region.
func (c *IBMClient) GetImages(region string) (map[string]bool, error) {
	images := map[string]bool{}
	options := &vpcv1.ListImagesOptions{}
	result, _, err := c.VPC.ListImagesWithContext(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to list images in region %s: %w", region, err)
	}
	for _, image := range result.Images {
		images[*image.Name] = true
	}
	return images, nil
}

// GetVPC is a wrapper to fetch a VPC by ID.
func (c *IBMClient) GetVPC(id string) (*vpcv1.VPC, error) {
	options := &vpcv1.GetVPCOptions{
		ID: &id,
	}
	vpc, _, err := c.VPC.GetVPCWithContext(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC with ID %s: %w", id, err)
	}
	return vpc, nil
}

// GetBackupPolicies is a wrapper to list Backup policies
func (c *IBMClient) GetBackupPolicies() (map[string]bool, error) {
	policies := map[string]bool{}
	options := &vpcv1.ListBackupPoliciesOptions{}
	result, _, err := c.VPC.ListBackupPoliciesWithContext(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to list backup policies: %w", err)
	}
	for _, policyintf := range result.BackupPolicies {
		policy := policyintf.(*vpcv1.BackupPolicy)
		policies[*policy.ID] = true
	}
	return policies, nil
}
