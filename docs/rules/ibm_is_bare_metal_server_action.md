# `ibm_is_bare_metal_server_action`

This rule checks the configuration of bare metal server actions.

## Example

```hcl
resource "ibm_is_bare_metal_server_action" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  action            = "stop"
  stop_type         = "soft"
}
```

```console
$ tflint
1 issue(s) found:
Error: `action` must be one of: start, stop, reboot (ibm_is_bare_metal_server_action)
  on main.tf line 4:
   4:   action            = "invalid"
```

## Why

The action configuration requires specific attributes and valid values to function properly. For example, the `action` must be one of `start`, `stop`, or `reboot`. Additionally, the `stop_type` must be specified when the action is `stop` and must be either `hard` or `soft`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_action" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  action            = "stop"  # Must be one of: start, stop, reboot
  stop_type         = "soft"  # Must be either 'hard' or 'soft' when action is 'stop'
}
```
