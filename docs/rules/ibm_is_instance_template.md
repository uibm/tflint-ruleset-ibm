# `ibm_is_instance_template`

This rule checks instance template configuration.

## Example

```hcl
resource "ibm_is_instance_template" "example" {
  name   = "template-1"
  profile = "invalid-profile"
  zone   = "us-south1" # Invalid format
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid instance profile specified (ibm_is_instance_template)
  on main.tf line 3:
   3:   profile = "invalid-profile"
```

## Why

Templates require `name`, valid `profile`, `image`, `vpc`, `zone`, and `keys`. Zone format must be region-number. Primary network attachment needs virtual network interface ID.

## How To Fix

```hcl
resource "ibm_is_instance_template" "example" {
  name    = "template-1"
  profile = "bx2-2x8"  # Valid profile
  image   = ibm_is_image.example.id
  vpc     = ibm_is_vpc.example.id
  zone    = "us-south-1"
  keys    = [ibm_is_ssh_key.example.id]

  primary_network_attachment {
    name = "primary-nic"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.example.id
    }
  }
}
```