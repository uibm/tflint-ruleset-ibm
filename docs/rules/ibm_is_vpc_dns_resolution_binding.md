# `ibm_is_vpc_dns_resolution_binding`

This rule checks VPC DNS resolution binding configuration.

## Example

```hcl
resource "ibm_is_vpc_dns_resolution_binding" "example" {
  name = "binding-1"
  # Missing vpc_id and vpc block
}
```

```console
$ tflint
1 issue(s) found:
Error: either vpc_id attribute or vpc block must be specified (ibm_is_vpc_dns_resolution_binding)
  on main.tf line 1:
   1: resource "ibm_is_vpc_dns_resolution_binding" "example" {
```

## Why

DNS bindings require `name` and either `vpc_id` or `vpc { id }` block. Names must be 1-63 characters. Cannot specify both vpc_id and vpc block.

## How To Fix

```hcl
resource "ibm_is_vpc_dns_resolution_binding" "example" {
  name   = "binding-1"
  vpc_id = ibm_is_vpc.example.id
  # OR
  vpc {
    id = ibm_is_vpc.example.id
  }
}
```