# `ibm_is_dedicated_host_group`

This rule checks the configuration of dedicated host groups.

## Example

```hcl
resource "ibm_is_dedicated_host_group" "example" {
  family = "memory"
  class  = "bx2d"
  zone   = "us-south-1"
  name   = "example-host-group"
}
```

```console
$ tflint
1 issue(s) found:
Error: `family` must be one of: memory, balanced, compute (ibm_is_dedicated_host_group)
  on main.tf line 4:
   4:   family = "invalid"
```

## Why

The dedicated host group configuration requires specific attributes and valid values to function properly. For example, the `family`, `class`, `zone`, and `name` attributes are required. Additionally, the `family` must be one of `memory`, `balanced`, or `compute`, and the `class` must be one of `bx2d`, `mx2d`, or `cx2d`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_dedicated_host_group" "example" {
  family = "memory"  # Must be one of: memory, balanced, compute
  class  = "bx2d"  # Must be one of: bx2d, mx2d, cx2d
  zone   = "us-south-1"  # Must be a valid zone
  name   = "example-host-group"  # Cannot be empty
}
```