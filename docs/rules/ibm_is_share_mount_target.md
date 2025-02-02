# `ibm_is_share_mount_target`

## Example
```hcl
resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.example.id
  name  = "example-mount-target"

  virtual_network_interface {
    subnet = ibm_is_subnet.example.id

    primary_ip {
      name = "example-ip"
    }
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: primary_ip block must be specified in virtual_network_interface (ibm_is_share_mount_target)
  on main.tf line 5:
   5:   virtual_network_interface {
```

## Why
The `share`, `name`, and `virtual_network_interface` block are required. The `primary_ip` block must also be specified within the `virtual_network_interface`.

## How To Fix
Ensure all required attributes and blocks are specified:
```hcl
resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.example.id
  name  = "example-mount-target"

  virtual_network_interface {
    subnet = ibm_is_subnet.example.id

    primary_ip {
      name = "example-ip"
    }
  }
}
```