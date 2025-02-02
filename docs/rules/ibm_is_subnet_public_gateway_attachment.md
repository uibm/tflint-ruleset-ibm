# `ibm_is_subnet_public_gateway_attachment`

This rule checks the configuration of subnet public gateway attachments.

## Example

```hcl
resource "ibm_is_subnet_public_gateway_attachment" "example" {
  public_gateway = ibm_is_public_gateway.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `subnet` attribute must be specified (ibm_is_subnet_public_gateway_attachment)
  on main.tf line 1:
   1: resource "ibm_is_subnet_public_gateway_attachment" "example" {
```

## Why

The public gateway attachment requires both `subnet` and `public_gateway` attributes to properly enable internet access for a subnet. Missing either attribute will prevent proper network configuration.

## How To Fix

Ensure both required attributes are specified:

```hcl
resource "ibm_is_subnet_public_gateway_attachment" "example" {
  subnet         = ibm_is_subnet.example.id
  public_gateway = ibm_is_public_gateway.example.id
}
```