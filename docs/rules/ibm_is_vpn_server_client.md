# `ibm_is_vpn_server_client`

This rule checks VPN server client associations.

## Example

```hcl
resource "ibm_is_vpn_server_client" "example" {
  vpn_server = ibm_is_vpn_server.example.id
  # Missing vpn_client
}
```

```console
$ tflint
1 issue(s) found:
Error: `vpn_client` attribute must be specified (ibm_is_vpn_server_client)
  on main.tf line 1:
   1: resource "ibm_is_vpn_server_client" "example" {
```

## Why

Client associations require both `vpn_server` and `vpn_client` attributes to link resources properly.

## How To Fix

```hcl
resource "ibm_is_vpn_server_client" "example" {
  vpn_server = ibm_is_vpn_server.example.id
  vpn_client = ibm_is_vpn_gateway_connection.example.id
}
```