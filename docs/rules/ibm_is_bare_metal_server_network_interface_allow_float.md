# `ibm_is_bare_metal_server_network_interface_allow_float`

This rule checks the configuration of bare metal server network interfaces that allow floating IPs.

## Example

```hcl
resource "ibm_is_bare_metal_server_network_interface_allow_float" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  subnet            = ibm_is_subnet.example.id
  name              = "example-nic"
  vlan              = 100
}
```

```console
$ tflint
1 issue(s) found:
Error: `vlan` must be between 1 and 4094 (ibm_is_bare_metal_server_network_interface_allow_float)
  on main.tf line 4:
   4:   vlan              = 5000
```

## Why

The network interface configuration requires specific attributes and valid values to function properly. For example, the `bare_metal_server`, `subnet`, `name`, and `vlan` attributes are required. Additionally, the `vlan` must be between 1 and 4094.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server_network_interface_allow_float" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  subnet            = ibm_is_subnet.example.id
  name              = "example-nic"  # Cannot be empty
  vlan              = 100  # Must be between 1 and 4094
}
```