# `ibm_is_lb_pool_member`

This rule ensures that the required attributes are specified for the `ibm_is_lb_pool_member` resource.

## Example

```hcl
resource "ibm_is_lb_pool_member" "example" {
  lb             = ibm_is_lb.example.id
  pool           = ibm_is_lb_pool.example.id
  port           = 80
  target_address = "192.168.1.1"
}
```

```console
$ tflint
1 issue(s) found:
Error: `port` attribute must be specified (ibm_is_lb_pool_member)
  on main.tf line 1:
   1: resource "ibm_is_lb_pool_member" "example" {
```

## Why

The `lb`, `pool`, and `port` attributes are required to define a valid load balancer pool member. Additionally, either `target_address` or `target_id` must be specified. The `port` must be between 1 and 65535, and the `weight` (if specified) must be between 0 and 100.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_lb_pool_member" "example" {
  lb             = ibm_is_lb.example.id
  pool           = ibm_is_lb_pool.example.id
  port           = 80  # Must be between 1 and 65535
  target_address = "192.168.1.1"  # Either target_address or target_id must be specified
  weight         = 50  # Must be between 0 and 100 if specified
}
```