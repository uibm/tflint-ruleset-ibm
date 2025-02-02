# `ibm_is_placement_group`

## Example
```hcl
resource "ibm_is_placement_group" "example" {
  strategy = "host_spread"
  name     = "example-group"
}
```

```console
$ tflint
1 issue(s) found:
Error: strategy must be one of: host_spread, power_spread, rack_spread (ibm_is_placement_group)
  on main.tf line 2:
   2:   strategy = "invalid-strategy"
```

## Why
The `strategy` must be one of the supported types (`host_spread`, `power_spread`, `rack_spread`). Invalid or missing strategies will result in errors.

## How To Fix
Ensure the `strategy` attribute is specified with a valid value:
```hcl
resource "ibm_is_placement_group" "example" {
  strategy = "host_spread"
  name     = "example-group"
}
```