# `ibm_is_snapshot_consistency_group`

This rule checks the configuration of snapshot consistency groups.

## Example

```hcl
resource "ibm_is_snapshot_consistency_group" "example" {
  delete_snapshots_on_delete = true

  snapshots {
    source_volume = ibm_is_volume.example.id
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_snapshot_consistency_group)
  on main.tf line 1:
   1: resource "ibm_is_snapshot_consistency_group" "example" {
```

## Why

The snapshot consistency group configuration requires specific attributes and valid values to function properly. The `name` attribute is required, and at least one `snapshots` block must be specified with both `name` and `source_volume` attributes. Names cannot be empty or exceed 63 characters.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_snapshot_consistency_group" "example" {
  name                       = "example-group"  # Cannot be empty or >63 chars
  delete_snapshots_on_delete = true

  snapshots {
    name          = "snapshot-1"  # Required in each snapshots block
    source_volume = ibm_is_volume.example.id  # Required in each snapshots block
  }
}
```