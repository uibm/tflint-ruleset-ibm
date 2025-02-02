# `ibm_is_dedicated_host_disk_management`

This rule checks the configuration of dedicated host disk management.

## Example

```hcl
resource "ibm_is_dedicated_host_disk_management" "example" {
  dedicated_host = ibm_is_dedicated_host.example.id

  disks {
    name = "disk-1"
    id   = "disk-id-1"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `dedicated_host` attribute must be specified (ibm_is_dedicated_host_disk_management)
  on main.tf line 1:
   1: resource "ibm_is_dedicated_host_disk_management" "example" {
```

## Why

The dedicated host disk management configuration requires specific attributes and valid values to function properly. For example, the `dedicated_host` attribute is required, and at least one `disks` block must be specified with valid `name` and `id` attributes.

## How To Fix

Ensure all required attributes and blocks are specified with valid values:

```hcl
resource "ibm_is_dedicated_host_disk_management" "example" {
  dedicated_host = ibm_is_dedicated_host.example.id  # Must be specified

  disks {
    name = "disk-1"  # Cannot be empty
    id   = "disk-id-1"  # Must be specified
  }
}
```