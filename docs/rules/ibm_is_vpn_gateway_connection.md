# `ibm_is_vpn_gateway_connection`

This rule checks VPN gateway connection configuration.

## Example

```hcl
resource "ibm_is_vpn_gateway_connection" "example" {
  name          = "conn-1"
  preshared_key = "secret"
  # Missing vpn_gateway and peer block
}
```

```console
$ tflint
1 issue(s) found:
Error: `vpn_gateway` attribute must be specified (ibm_is_vpn_gateway_connection)
  on main.tf line 1:
   1: resource "ibm_is_vpn_gateway_connection" "example" {
```

## Why

VPN connections require `name`, `vpn_gateway`, and `preshared_key`. Peer configurations need `address` and `cidrs` in a `peer` block. Dead peer detection requires `interval`, `timeout`, and `action` when enabled.

## How To Fix

```hcl
resource "ibm_is_vpn_gateway_connection" "example" {
  name          = "conn-1"  # <= 63 chars
  vpn_gateway   = ibm_is_vpn_gateway.example.id
  preshared_key = "secret"
  
  peer {
    address = "192.168.1.1"  # Valid IPv4
    cidrs   = ["10.0.0.0/24"]
  }
  
  dead_peer_detection {
    action   = "restart"  # restart/clear/hold/none
    interval = 30
    timeout  = 120
  }
}
```