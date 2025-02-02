# `ibm_is_cluster_network`

This rule checks the configuration of cluster networks.

## Example

```hcl
resource "ibm_is_cluster_network" "example" {
  name    = "example-cluster-network"
  profile = "h100"
  zone    = "us-south-1"

  subnet_prefixes {
    cidr = "10.240.0.0/24"
  }

  vpc {
    id = ibm_is_vpc.example.id
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `subnet_prefixes` block must be specified (ibm_is_cluster_network)
  on main.tf line 1:
   1: resource "ibm_is_cluster_network" "example" {
```

## Why

The cluster network configuration requires specific attributes and valid values to function properly. For example, the `name`, `profile`, and `zone` attributes are required, and the `subnet_prefixes` and `vpc` blocks must be specified with valid CIDR and VPC IDs.

## How To Fix

Ensure all required attributes and blocks are specified with valid values:

```hcl
resource "ibm_is_cluster_network" "example" {
  name    = "example-cluster-network"  # Cannot be empty
  profile = "h100"  # Must be a valid profile
  zone    = "us-south-1"  # Must be a valid zone

  subnet_prefixes {
    cidr = "10.240.0.0/24"  # Must be a valid CIDR
  }

  vpc {
    id = ibm_is_vpc.example.id  # Must be specified
  }
}
```