# `ibm_is_vpn_server`

This rule checks VPN server configuration.

## Example

```hcl
resource "ibm_is_vpn_server" "example" {
  name            = "server-1"
  certificate_crn = "crn:v1:...:certificate:1234"
  client_ip_pool  = "192.168.1.0/24"
  subnets         = [ibm_is_subnet.example.id]
  protocol        = "icmp"  # Invalid protocol
}
```

```console
$ tflint
1 issue(s) found:
Error: client_authentication block must be specified (ibm_is_vpn_server)
  on main.tf line 1:
   1: resource "ibm_is_vpn_server" "example" {
```

## Why

VPN servers require `name`, `certificate_crn`, `client_ip_pool`, `subnets`, and `client_authentication` block. Protocol must be udp/tcp. Client IP pool must be valid CIDR.

## How To Fix

```hcl
resource "ibm_is_vpn_server" "example" {
  name            = "server-1"
  certificate_crn = "crn:v1:...:certificate:1234"
  client_ip_pool  = "192.168.1.0/24"  # Valid CIDR
  subnets         = [ibm_is_subnet.example.id]
  protocol        = "udp"  # Valid protocol
  
  client_authentication {
    method         = "certificate"
    client_ca_crn = "crn:v1:...:certificate:5678"
  }
}
```