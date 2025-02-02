# `ibm_is_vpc_routing_table_route`

This rule validates VPC routing table routes.

## Example

```hcl
resource "ibm_is_vpc_routing_table_route" "example" {
  action      = "forward" # Invalid action
  destination = "192.168.1.0/33"
}
```

```console
$ tflint
1 issue(s) found:
Error: action must be either 'deliver' or 'drop' (ibm_is_vpc_routing_table_route)
  on main.tf line 4:
   4:   action      = "forward"
```

## Why

Routes require `vpc`, `routing_table`, `destination`, `action`, `zone`, and `name`. Action must be deliver/drop. Deliver actions require next_hop. Destination must be valid CIDR.

## How To Fix

```hcl
resource "ibm_is_vpc_routing_table_route" "example" {
  name           = "example-route"
  vpc            = ibm_is_vpc.example.id
  routing_table  = ibm_is_vpc_routing_table.example.id
  zone           = "us-south-1"
  destination    = "192.168.1.0/24"
  action         = "deliver"
  next_hop       = "10.240.0.5" # Required for deliver
}
```