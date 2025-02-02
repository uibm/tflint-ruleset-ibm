# `ibm_is_virtual_network_interface_ip`

This rule checks the configuration of IBM Cloud virtual network interface IPs.

## Example

```hcl
resource "ibm_is_virtual_network_interface_ip" "example" {
  reserved_ip             = ibm_is_reserved_ip.example.id
  virtual_network_interface = ibm_is_virtual_network_interface.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `reserved_ip` attribute must be specified (ibm_is_virtual_network_interface_ip)
  on main.tf line 1:
   1: resource "ibm_is_virtual_network_interface_ip" "example" {
```

## Why

Virtual network interface IP configuration requires both the `reserved_ip` and `virtual_network_interface` attributes.

## How To Fix

Ensure both the `reserved_ip` and `virtual_network_interface` attributes are specified:

```hcl
resource "ibm_is_virtual_network_interface_ip" "example" {
  reserved_ip             = ibm_is_reserved_ip.example.id
  virtual_network_interface = ibm_is_virtual_network_interface.example.id
}
```