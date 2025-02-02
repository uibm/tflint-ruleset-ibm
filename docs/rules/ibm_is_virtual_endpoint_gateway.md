# `ibm_is_virtual_endpoint_gateway`

This rule checks the configuration of IBM Cloud virtual endpoint gateways.

## Example

```hcl
resource "ibm_is_virtual_endpoint_gateway" "example" {
  name           = "example-vgw"
  vpc            = ibm_is_vpc.example.id
  resource_group = ibm_resource_group.example.id
  target {
    name          = "example-target"
    resource_type = "provider_cloud_service"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_virtual_endpoint_gateway)
  on main.tf line 1:
   1: resource "ibm_is_virtual_endpoint_gateway" "example" {
```

## Why

Virtual endpoint gateway configuration requires the `name` and `vpc` attributes. The `name` must be between 1 and 63 characters. A `target` block must be specified, containing the `name` and `resource_type` attributes. The `resource_type` within the `target` block must be either `provider_cloud_service` or `provider_infrastructure_service`.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_virtual_endpoint_gateway" "example" {
  name           = "example-vgw"  # Must be between 1 and 63 characters
  vpc            = ibm_is_vpc.example.id
  resource_group = ibm_resource_group.example.id
  target {
    name          = "example-target"
    resource_type = "provider_cloud_service" # Must be either provider_cloud_service or provider_infrastructure_service
  }
}

resource "ibm_is_virtual_endpoint_gateway" "example_ips" {
  name           = "example-vgw-ips"  # Must be between 1 and 63 characters
  vpc            = ibm_is_vpc.example.id
  resource_group = ibm_resource_group.example.id
  ips {
    name   = "example-ip-1"
    subnet = ibm_is_subnet.example.id
  }
  ips {
    name   = "example-ip-2"
    subnet = ibm_is_subnet.example2.id
  }
  target {
    name          = "example-target"
    resource_type = "provider_cloud_service" # Must be either provider_cloud_service or provider_infrastructure_service
  }
}
```