# `ibm_is_virtual_network_interface_floating_ip`

This rule checks the configuration of IBM Cloud virtual network interface floating IPs.

## Example

```hcl
resource "ibm_is_virtual_network_interface_floating_ip" "example" {
  virtual_network_interface = ibm_is_virtual_network_interface.example.id
  floating_ip               = ibm_is_floating_ip.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `virtual_network_interface` attribute must be specified (ibm_is_virtual_network_interface_floating_ip)
  on main.tf line 1:
   1: resource "ibm_is_virtual_network_interface_floating_ip" "example" {
```

## Why

Virtual network interface floating IP configuration requires both the `virtual_network_interface` and `floating_ip` attributes.

## How To Fix

Ensure both the `virtual_network_interface` and `floating_ip` attributes are specified:

```hcl
resource "ibm_is_virtual_network_interface_floating_ip" "example" {
  virtual_network_interface = ibm_is_virtual_network_interface.example.id
  floating_ip               = ibm_is_floating_ip.example.id
}
```