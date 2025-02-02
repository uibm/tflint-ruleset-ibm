# `ibm_is_vpn_gateway`

This rule validates VPN gateway configuration.

## Example

```hcl
resource "ibm_is_vpn_gateway" "example" {
  name   = "gateway-1"
  mode   = "mixed"  # Invalid mode
}
```

```console
$ tflint
1 issue(s) found:
Error: `subnet` attribute must be specified (ibm_is_vpn_gateway)
  on main.tf line 1:
   1: resource "ibm_is_vpn_gateway" "example" {
```

## Why

VPN gateways require `name` and `subnet`. The `mode` must be either 'policy' or 'route' if specified. Names must be 1-63 characters.

## How To Fix

```hcl
resource "ibm_is_vpn_gateway" "example" {
  name   = "gateway-1"
  subnet = ibm_is_subnet.example.id
  mode   = "policy"  # Valid mode
}
```