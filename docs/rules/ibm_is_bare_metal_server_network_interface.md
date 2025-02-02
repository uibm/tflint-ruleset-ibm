# `ibm_is_bare_metal_server_network_interface`

This rule validates the configuration of bare metal server network interfaces.

## Example

```hcl
resource "ibm_is_bare_metal_server_network_interface" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  subnet            = ibm_is_subnet.example.id
  name              = "example-nic"
  allow_ip_spoofing = false
  vlan              = 100
}
```

```console
$ tflint
1 issue(s) found:
Error: cannot specify both allowed_vlans and vlan (ibm_is_bare_metal_server_network_interface)
  on main.tf line 1:
   1: resource "ibm_is_bare_metal_server_network_interface" "example" {
```

## Why

The network interface configuration requires specific attributes and valid values to function properly. For example, the `bare_metal_server`, `subnet`, and `name` attributes are required. Additionally, the `vlan` must be between 1 and 4094, and you cannot specify both `allowed_vlans` and `vlan`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_network_interface" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  subnet            = ibm_is_subnet.example.id
  name              = "example-nic"
  allow_ip_spoofing = false
  vlan              = 100  # Must be between 1 and 4094
}
```