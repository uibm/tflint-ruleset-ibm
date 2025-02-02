# `ibm_is_lb_pool`

This rule checks the configuration of load balancer pools.

## Example

```hcl
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 5
  health_retries = 2
  health_timeout = 10
  health_type    = "http"
}
```

```console
$ tflint
1 issue(s) found:
Error: `algorithm` must be one of: round_robin, weighted_round_robin, least_connections (ibm_is_lb_pool)
  on main.tf line 4:
   4:   algorithm      = "invalid"
```

## Why

The load balancer pool configuration requires specific attributes and valid values to function properly. For example, the `name`, `lb`, `algorithm`, and `protocol` attributes are required. Additionally, the `algorithm` must be one of `round_robin`, `weighted_round_robin`, or `least_connections`, and the `protocol` must be one of `http`, `https`, `tcp`, or `udp`. If health check attributes are specified, they must all be present and valid.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"  # Must be one of: round_robin, weighted_round_robin, least_connections
  protocol       = "http"  # Must be one of: http, https, tcp, udp
  health_delay   = 5  # Must be between 1 and 300 seconds
  health_retries = 2  # Must be between 1 and 10
  health_timeout = 10  # Must be between 1 and 120 seconds
  health_type    = "http"  # Must be one of: http, https, tcp
}
```