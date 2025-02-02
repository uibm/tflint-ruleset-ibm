# `ibm_is_vpc`

This rule ensures that the required attributes are specified for the `ibm_is_vpc` resource.

## Example

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_vpc)
  on main.tf line 1:
   1: resource "ibm_is_vpc" "example" {
```

## Why

The `name` attribute is required to define a valid VPC. Missing this attribute will result in errors during Terraform apply.

## How To Fix

Ensure the `name` attribute is specified:

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
```