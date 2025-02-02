# `ibm_is_virtual_endpoint_gateway_ip`

This rule checks the configuration of IBM Cloud virtual endpoint gateway IPs.

## Example

```hcl
resource "ibm_is_virtual_endpoint_gateway_ip" "example" {
  gateway     = ibm_is_virtual_endpoint_gateway.example.id
  reserved_ip = ibm_is_reserved_ip.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `gateway` attribute must be specified (ibm_is_virtual_endpoint_gateway_ip)
  on main.tf line 1:
   1: resource "ibm_is_virtual_endpoint_gateway_ip" "example" {
```

## Why

Virtual endpoint gateway IP configuration requires both the `gateway` and `reserved_ip` attributes.

## How To Fix

Ensure both the `gateway` and `reserved_ip` attributes are specified:

```hcl
resource "ibm_is_virtual_endpoint_gateway_ip" "example" {
  gateway     = ibm_is_virtual_endpoint_gateway.example.id
  reserved_ip = ibm_is_reserved_ip.example.id
}
```