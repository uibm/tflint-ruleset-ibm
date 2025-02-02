# `ibm_is_lb_listener_policy`

## Example
```hcl
resource "ibm_is_lb_listener_policy" "example" {
  lb               = ibm_is_lb.example.id
  listener         = ibm_is_lb_listener.example.listener_id
  action           = "forward"
  priority         = 5
  name             = "example-policy"

  rules {
    condition = "equals"
    type      = "host"
    field     = "Host"
    value     = "example.com"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: action must be one of: forward, redirect, reject (ibm_is_lb_listener_policy)
  on main.tf line 4:
   4:   action           = "invalid-action"
```

## Why
The `action` must be one of the supported types (`forward`, `redirect`, `reject`). Additionally, the `rules` block requires specific attributes like `condition`, `type`, `field`, and `value`.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_lb_listener_policy" "example" {
  lb               = ibm_is_lb.example.id
  listener         = ibm_is_lb_listener.example.listener_id
  action           = "forward"
  priority         = 5
  name             = "example-policy"

  rules {
    condition = "equals"
    type      = "host"
    field     = "Host"
    value     = "example.com"
  }
}
```