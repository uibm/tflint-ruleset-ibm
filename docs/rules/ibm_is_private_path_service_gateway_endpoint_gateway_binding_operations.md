# `ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations`

This rule checks the configuration of IBM Cloud private path service gateway endpoint gateway binding operations.

## Example

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_endpoint_gateway" "example" {
  name = "example-endpoint-gateway"
}

resource "ibm_is_endpoint_gateway_ip_configuration" "example" {
  endpoint_gateway = ibm_is_endpoint_gateway.example.id
  ip_address       = "192.168.1.1"
}

resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding" "example" {
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
  endpoint_gateway            = ibm_is_endpoint_gateway.example.id
  ip_configuration             = ibm_is_endpoint_gateway_ip_configuration.example.id
}

resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" "example" {
  access_policy                = "permit"
  endpoint_gateway_binding     = ibm_is_private_path_service_gateway_endpoint_gateway_binding.example.id
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `access_policy` attribute must be specified (ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations)

  on main.tf line 1:
   1: resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" "example" {
```

## Why

Private path service gateway endpoint gateway binding operations configuration requires the `access_policy`, `endpoint_gateway_binding`, and `private_path_service_gateway` attributes. The `access_policy` must be either `permit` or `deny`.

## How To Fix

Ensure all required attributes are specified and that the `access_policy` is valid:

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_endpoint_gateway" "example" {
  name = "example-endpoint-gateway"
}

resource "ibm_is_endpoint_gateway_ip_configuration" "example" {
  endpoint_gateway = ibm_is_endpoint_gateway.example.id
  ip_address       = "192.168.1.1"
}

resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding" "example" {
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
  endpoint_gateway            = ibm_is_endpoint_gateway.example.id
  ip_configuration             = ibm_is_endpoint_gateway_ip_configuration.example.id
}

resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" "example" {
  access_policy                = "permit" # Must be "permit" or "deny"
  endpoint_gateway_binding     = ibm_is_private_path_service_gateway_endpoint_gateway_binding.example.id
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```