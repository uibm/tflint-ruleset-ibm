# `ibm_is_bare_metal_server_initialization`

This rule checks the configuration of IBM Cloud bare metal server initialization.

## Example

```hcl
resource "ibm_is_bare_metal_server" "example" {
  name = "example-bare-metal-server"
}

resource "ibm_is_bare_metal_server_initialization" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  image             = ibm_is_image.example.id
  keys              = [ibm_is_ssh_key.example.id]
  user_data         = "echo 'Hello, world!' > /tmp/hello.txt"
}
```

```console
$ tflint
1 issue(s) found:

Error: `bare_metal_server` attribute must be specified (ibm_is_bare_metal_server_initialization)

  on main.tf line 1:
   1: resource "ibm_is_bare_metal_server_initialization" "example" {
```

## Why

Bare metal server initialization requires the `bare_metal_server`, `image`, and `keys` attributes. At least one SSH key must be specified in the `keys` attribute. The `user_data` attribute is optional but cannot exceed 16KB in size.

## How To Fix

Ensure all required attributes are specified and that `user_data` (if provided) adheres to the size limit:

```hcl
resource "ibm_is_bare_metal_server" "example" {
  name = "example-bare-metal-server"
}

resource "ibm_is_bare_metal_server_initialization" "example" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
  image             = ibm_is_image.example.id
  keys              = [ibm_is_ssh_key.example.id] # At least one key
  user_data         = "echo 'Hello, world!' > /tmp/hello.txt" # Optional, but <= 16KB
}
```