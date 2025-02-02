# `ibm_is_subnet_reserved_ip_patch`

This rule checks the configuration of IBM Cloud subnet reserved IP patches.

## Example

```hcl
resource "ibm_is_subnet_reserved_ip_patch" "example" {
  subnet      = ibm_is_subnet.example.id
  reserved_ip = ibm_is_reserved_ip.example.id
  name        = "example-patch"
}
```

```console
$ tflint
1 issue(s) found:
Error: `subnet` attribute must be specified (ibm_is_subnet_reserved_ip_patch)
  on main.tf line 1:
   1: resource "ibm_is_subnet_reserved_ip_patch" "example" {
```

## Why

Subnet reserved IP patch configuration requires the `subnet` and `reserved_ip` attributes. The `name` attribute is optional but, if provided, must be no longer than 63 characters.

## How To Fix

Ensure the required attributes are specified and the `name` (if provided) adheres to the length restriction:

```hcl
resource "ibm_is_subnet_reserved_ip_patch" "example" {
  subnet      = ibm_is_subnet.example.id
  reserved_ip = ibm_is_reserved_ip.example.id
  name        = "example-patch" # Optional, but if provided, must be <= 63 characters
}
```
