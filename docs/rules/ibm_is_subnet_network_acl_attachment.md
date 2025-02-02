# `ibm_is_subnet_network_acl_attachment`

This rule checks the configuration of subnet Network ACL attachments.

## Example

```hcl
resource "ibm_is_subnet_network_acl_attachment" "example" {
  subnet = ibm_is_subnet.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `network_acl` attribute must be specified (ibm_is_subnet_network_acl_attachment)
  on main.tf line 1:
   1: resource "ibm_is_subnet_network_acl_attachment" "example" {
```

## Why

The Network ACL attachment requires both `subnet` and `network_acl` attributes to properly associate a subnet with a network ACL. Missing either attribute will prevent proper network security configuration.

## How To Fix

Ensure both required attributes are specified:

```hcl
resource "ibm_is_subnet_network_acl_attachment" "example" {
  subnet      = ibm_is_subnet.example.id
  network_acl = ibm_is_network_acl.example.id
}
```