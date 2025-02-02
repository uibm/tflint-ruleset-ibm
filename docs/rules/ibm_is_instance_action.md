# `ibm_is_instance_action`

## Example
```hcl
resource "ibm_is_instance_action" "example" {
  action  = "start"
  instance = ibm_is_instance.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: action must be one of: start, stop, restart (ibm_is_instance_action)
  on main.tf line 2:
   2:   action  = "invalid-action"
```

## Why
The `action` attribute must be one of the supported actions (`start`, `stop`, `restart`). Invalid or missing actions will result in errors.

## How To Fix
Ensure the `action` attribute is specified with a valid value:
```hcl
resource "ibm_is_instance_action" "example" {
  action  = "start"
  instance = ibm_is_instance.example.id
}
```