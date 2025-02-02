# `ibm_is_vpc_routing_table`

This rule checks VPC routing table configuration.

## Example

```hcl
resource "ibm_is_vpc_routing_table" "example" {
  accept_routes_from_resource_type = ["invalid-type"]
}
```

```console
$ tflint
1 issue(s) found:
Error: accept_routes_from_resource_type must be either 'vpn_server' or 'vpn_gateway' (ibm_is_vpc_routing_table)
  on main.tf line 3:
   3:   accept_routes_from_resource_type = ["invalid-type"]
```

## Why

Routing tables require `vpc` and `name`. The `accept_routes_from_resource_type` must contain only 'vpn_server' or 'vpn_gateway' if specified.

## How To Fix

```hcl
resource "ibm_is_vpc_routing_table" "example" {
  name = "example-rt"
  vpc  = ibm_is_vpc.example.id
  accept_routes_from_resource_type = ["vpn_gateway"]
}
```