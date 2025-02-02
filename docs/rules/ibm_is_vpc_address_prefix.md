# `ibm_is_vpc_address_prefix`

This rule validates VPC address prefix configuration.

## Example

```hcl
resource "ibm_is_vpc_address_prefix" "example" {
  name = "" # Empty name
  cidr = "192.168.1.0/33" # Invalid CIDR
}
```

```console
$ tflint
1 issue(s) found:
Error: name cannot be empty (ibm_is_vpc_address_prefix)
  on main.tf line 1:
   1: resource "ibm_is_vpc_address_prefix" "example" {
```

## Why

Address prefixes require `name`, `zone`, `vpc`, and valid `cidr`. Names must be 1-63 characters. CIDR must be valid IPv4 format.

## How To Fix

```hcl
resource "ibm_is_vpc_address_prefix" "example" {
  name = "example-prefix"
  zone = "us-south-1"
  vpc  = ibm_is_vpc.example.id
  cidr = "192.168.1.0/24" # Valid CIDR
}
```