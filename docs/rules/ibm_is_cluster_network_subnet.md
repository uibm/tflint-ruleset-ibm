# `ibm_is_cluster_network_subnet`

This rule checks the configuration of cluster network subnets.

## Example

```hcl
resource "ibm_is_cluster_network_subnet" "example" {
  cluster_network_id        = ibm_is_cluster_network.example.id
  name                      = "example-subnet"
  total_ipv4_address_count  = 16
}
```

```console
$ tflint
1 issue(s) found:
Error: `total_ipv4_address_count` must be a power of 2 between 8 and 8192 (ibm_is_cluster_network_subnet)
  on main.tf line 4:
   4:   total_ipv4_address_count  = 15
```

## Why

The cluster network subnet configuration requires specific attributes and valid values to function properly. For example, the `cluster_network_id`, `name`, and `total_ipv4_address_count` attributes are required. Additionally, the `total_ipv4_address_count` must be a power of 2 between 8 and 8192.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_cluster_network_subnet" "example" {
  cluster_network_id        = ibm_is_cluster_network.example.id
  name                      = "example-subnet"  # Cannot be empty
  total_ipv4_address_count  = 16  # Must be a power of 2 between 8 and 8192
}
```