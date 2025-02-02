# `ibm_is_instance`

This rule checks the configuration of IBM Cloud VPC instances.

## Example

```hcl
resource "ibm_is_instance" "example" {
  name   = "example-instance"
  profile = "bx2-2x8"
  image   = ibm_is_image.example.id
  vpc     = ibm_is_vpc.example.id
  zone    = "us-south-1"
  keys    = [ibm_is_ssh_key.example.id]
}
```

```console
$ tflint
1 issue(s) found:
Error: `profile` attribute must be specified (ibm_is_instance)
  on main.tf line 1:
   1: resource "ibm_is_instance" "example" {
```

## Why

The instance configuration requires specific attributes and valid values to function properly. For example, the `name`, `profile`, `image`, `vpc`, and `zone` attributes are required. Additionally, the `profile` and `image` must be valid IBM Cloud resources.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance" "example" {
  name   = "example-instance"
  profile = "bx2-2x8"  # Must be a valid instance profile
  image   = ibm_is_image.example.id  # Must be a valid image ID
  vpc     = ibm_is_vpc.example.id
  zone    = "us-south-1"
  keys    = [ibm_is_ssh_key.example.id]
}
```