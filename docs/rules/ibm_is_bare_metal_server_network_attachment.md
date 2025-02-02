# `ibm_is_bare_metal_server_network_attachment`

This rule checks the configuration of bare metal server network attachments.

## Example

```hcl
resource "ibm_is_bare_metal_server_network_attachment" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  interface_type    = "vlan"
  vlan              = 100
}
```

```console
$ tflint
1 issue(s) found:
Error: `interface_type` must be either 'vlan' or 'pci' (ibm_is_bare_metal_server_network_attachment)
  on main.tf line 4:
   4:   interface_type    = "invalid"
```

## Why

The network attachment configuration requires specific attributes and valid values to function properly. For example, the `bare_metal_server` attribute is required, and the `interface_type` must be either `vlan` or `pci`. If the `interface_type` is `vlan`, the `vlan` attribute must be specified, and if it is `pci`, the `allowed_vlans` attribute must be specified.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_network_attachment" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  interface_type    = "vlan"  # Must be either 'vlan' or 'pci'
  vlan              = 100  # Must be specified if interface_type is 'vlan'
}
```