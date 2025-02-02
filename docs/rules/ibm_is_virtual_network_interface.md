# `ibm_is_virtual_network_interface`

This rule checks the configuration of IBM Cloud virtual network interfaces.

## Example

```hcl
resource "ibm_is_virtual_network_interface" "example" {
  name    = "example-vni"
  subnet  = ibm_is_subnet.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_virtual_network_interface)
  on main.tf line 1:
   1: resource "ibm_is_virtual_network_interface" "example" {
```

## Why

Virtual network interface configuration requires the `name` and `subnet` attributes. The `name` must be between 1 and 63 characters. The optional attribute `protocol_state_filtering_mode`, if specified, must be either `enabled` or `disabled`.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_virtual_network_interface" "example" {
  name    = "example-vni"  # Must be between 1 and 63 characters
  subnet  = ibm_is_subnet.example.id
}

resource "ibm_is_virtual_network_interface" "example_filtering" {
  name    = "example-vni-filtering"  # Must be between 1 and 63 characters
  subnet  = ibm_is_subnet.example.id
  protocol_state_filtering_mode = "enabled" # Must be either enabled or disabled
}
```