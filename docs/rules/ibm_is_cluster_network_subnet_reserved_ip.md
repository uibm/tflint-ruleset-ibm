# `ibm_is_cluster_network_subnet_reserved_ip`

This rule checks the configuration of IBM Cloud cluster network subnet reserved IPs.

## Example

```hcl
resource "ibm_is_cluster_network" "example" {
  name = "example-cluster-network"
}

resource "ibm_is_cluster_network_subnet" "example" {
  cluster_network = ibm_is_cluster_network.example.id
  name            = "example-cluster-network-subnet"
}

resource "ibm_is_cluster_network_subnet_reserved_ip" "example" {
  address                  = "10.0.0.10"
  cluster_network_id       = ibm_is_cluster_network.example.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.example.id
  name                     = "example-reserved-ip"
}
```

```console
$ tflint
1 issue(s) found:

Error: `address` attribute must be specified (ibm_is_cluster_network_subnet_reserved_ip)

  on main.tf line 1:
   1: resource "ibm_is_cluster_network_subnet_reserved_ip" "example" {
```

## Why

Cluster network subnet reserved IP configuration requires the `address`, `cluster_network_id`, `cluster_network_subnet_id`, and `name` attributes. The `address` must be a valid IPv4 address. The `name` must be between 1 and 63 characters.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_cluster_network" "example" {
  name = "example-cluster-network"
}

resource "ibm_is_cluster_network_subnet" "example" {
  cluster_network = ibm_is_cluster_network.example.id
  name            = "example-cluster-network-subnet"
}

resource "ibm_is_cluster_network_subnet_reserved_ip" "example" {
  address                  = "10.0.0.10" # Valid IPv4 address
  cluster_network_id       = ibm_is_cluster_network.example.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.example.id
  name                     = "example-reserved-ip" # Between 1 and 63 characters
}
```