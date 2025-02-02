# `ibm_is_subnet_routing_table_attachment`

This rule checks the configuration of IBM Cloud subnet routing table attachments.

## Example

```hcl
resource "ibm_is_subnet_routing_table_attachment" "example" {
  subnet        = ibm_is_subnet.example.id
  routing_table = ibm_is_routing_table.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `subnet` attribute must be specified (ibm_is_subnet_routing_table_attachment)
  on main.tf line 1:
   1: resource "ibm_is_subnet_routing_table_attachment" "example" {
```

## Why

Subnet routing table attachments require both the `subnet` and `routing_table` attributes to be specified.

## How To Fix

Ensure both the `subnet` and `routing_table` attributes are specified:

```hcl
resource "ibm_is_subnet_routing_table_attachment" "example" {
  subnet        = ibm_is_subnet.example.id
  routing_table = ibm_is_routing_table.example.id
}
```