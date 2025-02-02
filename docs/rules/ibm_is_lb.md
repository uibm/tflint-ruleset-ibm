# `ibm_is_lb`

This rule checks the configuration of load balancers.

## Example

```hcl
resource "ibm_is_lb" "example" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
  type    = "public"
  profile = "network-small"
}
```

```console
$ tflint
1 issue(s) found:
Error: `type` must be either 'public' or 'private' (ibm_is_lb)
  on main.tf line 4:
   4:   type    = "invalid"
```

## Why

The load balancer configuration requires specific attributes and valid values to function properly. For example, the `name` and `subnets` attributes are required. Additionally, the `type` must be either `public` or `private`, and the `profile` must be a valid load balancer profile.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_lb" "example" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
  type    = "public"  # Must be either 'public' or 'private'
  profile = "network-small"  # Must be a valid load balancer profile
}
```