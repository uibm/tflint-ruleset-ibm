# `ibm_is_cluster_network_interface`

This rule checks the configuration of cluster network interfaces.

## Example

```hcl
resource "ibm_is_cluster_network_interface" "example" {
  cluster_network_id = ibm_is_cluster_network.example.id
  name               = "example-nic"

  primary_ip {
    id = "primary-ip-1"
  }

  subnet {
    id = "subnet-1"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `primary_ip` block must be specified (ibm_is_cluster_network_interface)
  on main.tf line 1:
   1: resource "ibm_is_cluster_network_interface" "example" {
```

## Why

The cluster network interface configuration requires specific attributes and valid values to function properly. For example, the `cluster_network_id` and `name` attributes are required, and the `primary_ip` and `subnet` blocks must be specified with valid IDs.

## How To Fix

Ensure all required attributes and blocks are specified with valid values:

```hcl
resource "ibm_is_cluster_network_interface" "example" {
  cluster_network_id = ibm_is_cluster_network.example.id
  name               = "example-nic"  # Cannot be empty

  primary_ip {
    id = "primary-ip-1"  # Must be specified
  }

  subnet {
    id = "subnet-1"  # Must be specified
  }
}
```