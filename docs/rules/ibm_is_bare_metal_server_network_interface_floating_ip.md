# `ibm_is_bare_metal_server_network_interface_floating_ip`

This rule checks the configuration of floating IPs for bare metal server network interfaces.

## Example

```hcl
resource "ibm_is_bare_metal_server_network_interface_floating_ip" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  network_interface = ibm_is_bare_metal_server_network_interface.example.id
  floating_ip       = ibm_is_floating_ip.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `floating_ip` attribute must be specified (ibm_is_bare_metal_server_network_interface_floating_ip)
  on main.tf line 1:
   1: resource "ibm_is_bare_metal_server_network_interface_floating_ip" "example" {
```

## Why

The floating IP configuration requires specific attributes and valid values to function properly. For example, the `bare_metal_server`, `network_interface`, and `floating_ip` attributes are required.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_network_interface_floating_ip" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  network_interface = ibm_is_bare_metal_server_network_interface.example.id
  floating_ip       = ibm_is_floating_ip.example.id  # Must be specified
}
```