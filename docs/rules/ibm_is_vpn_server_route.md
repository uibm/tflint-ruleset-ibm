# `ibm_is_vpn_server_route`

This rule validates VPN server route configuration.

## Example

```hcl
resource "ibm_is_vpn_server_route" "example" {
  vpn_server  = ibm_is_vpn_server.example.id
  destination = "192.168.1.0/33"  # Invalid CIDR
  action      = "block"           # Invalid action
}
```

```console
$ tflint
1 issue(s) found:
Error: action must be either 'translate' or 'drop' (ibm_is_vpn_server_route)
  on main.tf line 4:
   4:   action      = "block"
```

## Why

Routes require `vpn_server`, valid `destination` CIDR, `action` (translate/drop), and `name`. Names must be 1-63 characters.

## How To Fix

```hcl
resource "ibm_is_vpn_server_route" "example" {
  name       = "route-1"
  vpn_server  = ibm_is_vpn_server.example.id
  destination = "192.168.1.0/24"  # Valid CIDR
  action      = "translate"
}
```