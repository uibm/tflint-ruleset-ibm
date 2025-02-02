# `ibm_is_bare_metal_server`

This rule checks the configuration of bare metal servers.

## Example

```hcl
resource "ibm_is_bare_metal_server" "example" {
  name   = "example-bare-metal-server"
  profile = "bx2-metal-192x768"
  image   = ibm_is_image.example.id
  zone    = "us-south-1"
  vpc     = ibm_is_vpc.example.id
  keys    = [ibm_is_ssh_key.example.id]

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `profile` attribute must be specified (ibm_is_bare_metal_server)
  on main.tf line 1:
   1: resource "ibm_is_bare_metal_server" "example" {
```

## Why

The bare metal server configuration requires specific attributes and valid values to function properly. For example, the `name`, `profile`, `image`, `zone`, `vpc`, and `keys` attributes are required. Additionally, the `primary_network_interface` block must be specified, and the `profile` must be a valid bare metal server profile.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_bare_metal_server" "example" {
  name   = "example-bare-metal-server"
  profile = "bx2-metal-192x768"  # Must be a valid bare metal server profile
  image   = ibm_is_image.example.id
  zone    = "us-south-1"
  vpc     = ibm_is_vpc.example.id
  keys    = [ibm_is_ssh_key.example.id]

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }
}
```