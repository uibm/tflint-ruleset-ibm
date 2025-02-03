# `ibm_is_private_path_service_gateway`

This rule checks the configuration of IBM Cloud private path service gateways.

## Example

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name                  = "example-gateway"
  default_access_policy = "permit"
  load_balancer          = ibm_is_lb.example.id
  zonal_affinity         = false
  service_endpoints      = ["10.0.0.10", "10.0.0.11"]
}
```

```console
$ tflint
1 issue(s) found:

Error: `name` attribute must be specified (ibm_is_private_path_service_gateway)

  on main.tf line 1:
   1: resource "ibm_is_private_path_service_gateway" "example" {
```

## Why

Private path service gateway configuration requires the `name`, `load_balancer`, and `default_access_policy` attributes. The `name` must be between 1 and 63 characters. The `default_access_policy` must be either `permit` or `deny`. The `service_endpoints` attribute is optional but cannot be empty if specified.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name                  = "example-gateway" # Between 1 and 63 characters
  default_access_policy = "permit" # Must be "permit" or "deny"
  load_balancer          = ibm_is_lb.example.id
  zonal_affinity         = false
  service_endpoints      = ["10.0.0.10", "10.0.0.11"] # Optional, but not empty if specified
}
```