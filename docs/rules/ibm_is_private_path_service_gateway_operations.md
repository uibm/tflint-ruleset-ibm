# `ibm_is_private_path_service_gateway_operations`

This rule checks the configuration of IBM Cloud private path service gateway operations.

## Example

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_operations" "example" {
  published                     = true
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `published` attribute must be specified (ibm_is_private_path_service_gateway_operations)

  on main.tf line 1:
   1: resource "ibm_is_private_path_service_gateway_operations" "example" {
```

## Why

Private path service gateway operations configuration requires the `published` and `private_path_service_gateway` attributes.

## How To Fix

Ensure both required attributes are specified:

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_operations" "example" {
  published                     = true
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```