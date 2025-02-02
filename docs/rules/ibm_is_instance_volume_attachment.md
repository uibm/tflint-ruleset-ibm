# `ibm_is_instance_volume_attachment`

This rule validates instance volume attachments.

## Example

```hcl
resource "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id
  name     = "attch-1"
  volume   = ibm_is_volume.example.id
  volume_name = "vol-1" # Both specified
}
```

```console
$ tflint
1 issue(s) found:
Error: cannot specify both volume and volume_name (ibm_is_instance_volume_attachment)
  on main.tf line 1:
   1: resource "ibm_is_instance_volume_attachment" "example" {
```

## Why

Attachments require `instance` and `name`. Must specify either `volume` or `volume_name` (not both). Capacity 10-2000GB if creating new volume.

## How To Fix

```hcl
resource "ibm_is_instance_volume_attachment" "example" {
  instance = ibm_is_instance.example.id
  name     = "attch-1"
  volume   = ibm_is_volume.example.id  # OR volume_name = "new-vol"
  profile  = "general-purpose"         # Required if volume_name used
  capacity = 100                       # Required if volume_name used
}
```