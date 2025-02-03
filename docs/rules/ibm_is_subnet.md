# `ibm_is_subnet`

This rule checks the configuration of IBM Cloud subnets.

## Example

```hcl
resource "ibm_is_subnet" "example" {
  name                  = "example-subnet"
  vpc                   = ibm_is_vpc.example.id
  zone                  = "us-south-1"
  ipv4_cidr_block       = "10.0.0.0/24"
  resource_group        = ibm_resource_group.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_subnet)
  on main.tf line 1:
   1: resource "ibm_is_subnet" "example" {
```

## Why

Subnet configuration requires specific attributes to be set. The `name`, `vpc`, and `zone` attributes are required.  Either `ipv4_cidr_block` or `total_ipv4_address_count` must be specified, but not both. The `name` must be between 1 and 63 characters. The `ipv4_cidr_block` must be a valid CIDR block. The `total_ipv4_address_count` must be a power of 2 between 8 and 8192. The `zone` must be a valid zone format (e.g., `region-number`, like `us-south-1`).

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_subnet" "example" {
  name                  = "example-subnet"  # Must be between 1 and 63 characters
  vpc                   = ibm_is_vpc.example.id
  zone                  = "us-south-1"      # Must be in format: region-number
  ipv4_cidr_block       = "10.0.0.0/24"   # Must be a valid CIDR block
  resource_group        = ibm_resource_group.example.id
}

# OR

resource "ibm_is_subnet" "example" {
  name                  = "example-subnet"  # Must be between 1 and 63 characters
  vpc                   = ibm_is_vpc.example.id
  zone                  = "us-south-1"      # Must be in format: region-number
  total_ipv4_address_count = 256            # Must be a power of 2 between 8 and 8192
  resource_group        = ibm_resource_group.example.id
}

```