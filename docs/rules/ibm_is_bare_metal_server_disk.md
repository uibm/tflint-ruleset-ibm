# `ibm_is_bare_metal_server_disk`

This rule ensures that the required attributes are specified for the `ibm_is_bare_metal_server_disk` resource.

## Example

```hcl
resource "ibm_is_bare_metal_server_disk" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  disk              = "disk-1"
  name              = "example-disk"
}
```

```console
$ tflint
1 issue(s) found:
Error: `disk` attribute must be specified (ibm_is_bare_metal_server_disk)
  on main.tf line 1:
   1: resource "ibm_is_bare_metal_server_disk" "example" {
```

## Why

The `bare_metal_server` and `disk` attributes are required to define a valid bare metal server disk. Additionally, the `name` attribute, if specified, cannot be empty.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_disk" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  disk              = "disk-1"  # Must be specified
  name              = "example-disk"  # Cannot be empty if specified
}
```