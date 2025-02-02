# `ibm_is_lb_listener`

This rule checks the configuration of load balancer listeners.

## Example

```hcl
resource "ibm_is_lb_listener" "example" {
  lb       = ibm_is_lb.example.id
  port     = 443
  protocol = "https"
  default_pool = ibm_is_lb_pool.example.id
  certificate_instance = ibm_is_certificate.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `protocol` must be one of: http, https, tcp, udp (ibm_is_lb_listener)
  on main.tf line 4:
   4:   protocol = "invalid"
```

## Why

The load balancer listener configuration requires specific attributes and valid values to function properly. For example, the `lb`, `port`, and `protocol` attributes are required. Additionally, the `protocol` must be one of `http`, `https`, `tcp`, or `udp`, and the `port` must be between 1 and 65535. If the protocol is `https`, the `certificate_instance` must be specified.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_lb_listener" "example" {
  lb       = ibm_is_lb.example.id
  port     = 443  # Must be between 1 and 65535
  protocol = "https"  # Must be one of: http, https, tcp, udp
  default_pool = ibm_is_lb_pool.example.id
  certificate_instance = ibm_is_certificate.example.id  # Required when protocol is https
}
```