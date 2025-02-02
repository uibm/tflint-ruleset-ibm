# `ibm_is_dedicated_host`

This rule checks the configuration of dedicated hosts.

## Example

```hcl
resource "ibm_is_dedicated_host" "example" {
  profile  = "bx2d-host-152x608"
  name     = "example-host"
  host_group = ibm_is_dedicated_host_group.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `profile` attribute must be specified (ibm_is_dedicated_host)
  on main.tf line 1:
   1: resource "ibm_is_dedicated_host" "example" {
```

## Why

The dedicated host configuration requires specific attributes and valid values to function properly. For example, the `profile`, `name`, and `host_group` attributes are required. Additionally, the `profile` must be a valid dedicated host profile, and the `name` cannot be longer than 63 characters.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_dedicated_host" "example" {
  profile  = "bx2d-host-152x608"  # Must be a valid profile
  name     = "example-host"  # Cannot be empty or longer than 63 characters
  host_group = ibm_is_dedicated_host_group.example.id  # Must be specified
}
```