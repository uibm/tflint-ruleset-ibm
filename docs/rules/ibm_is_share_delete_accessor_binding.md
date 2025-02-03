# `ibm_is_share_delete_accessor_binding`

This rule checks the configuration of IBM Cloud share delete accessor binding operations.

## Example

```hcl
resource "ibm_is_share" "example" {
  name = "example-share"
}

resource "ibm_is_share_accessor_binding" "example" {
  share           = ibm_is_share.example.id
  access_type     = "ip"
  ip_address_range = "10.0.0.0/24"
}

resource "ibm_is_share_delete_accessor_binding" "example" {
  share            = ibm_is_share.example.id
  accessor_binding = ibm_is_share_accessor_binding.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `share` attribute must be specified (ibm_is_share_delete_accessor_binding)

  on main.tf line 1:
   1: resource "ibm_is_share_delete_accessor_binding" "example" {
```

## Why

Share delete accessor binding configuration requires the `share` and `accessor_binding` attributes. The `accessor_binding` cannot be empty.

## How To Fix

Ensure both required attributes are specified and that the `accessor_binding` is not empty:

```hcl
resource "ibm_is_share" "example" {
  name = "example-share"
}

resource "ibm_is_share_accessor_binding" "example" {
  share           = ibm_is_share.example.id
  access_type     = "ip"
  ip_address_range = "10.0.0.0/24"
}

resource "ibm_is_share_delete_accessor_binding" "example" {
  share            = ibm_is_share.example.id
  accessor_binding = ibm_is_share_accessor_binding.example.id # Cannot be empty
}
```