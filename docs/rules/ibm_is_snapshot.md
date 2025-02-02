# `ibm_is_snapshot`

This rule checks the configuration of volume snapshots.

## Example

```hcl
resource "ibm_is_snapshot" "example" {
  name = "example-snapshot"
  clones = ["us-south-1", "invalid-zone"]
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid zone format in clones: invalid-zone. Must be in format: region-number (e.g., us-south-1) (ibm_is_snapshot)
  on main.tf line 3:
   3:   clones = ["us-south-1", "invalid-zone"]
```

## Why

The snapshot configuration requires specific attributes and valid values to function properly. The `name` and `source_volume` attributes are required. The `name` cannot be empty or exceed 63 characters. Clone zones must be in valid format, and timeouts (if specified) must use valid duration strings.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_snapshot" "example" {
  name         = "example-snapshot"  # Cannot be empty or >63 chars
  source_volume = ibm_is_volume.example.id  # Required
  clones       = ["us-south-1", "us-south-2"]  # Valid zones

  timeouts {
    create = "30m"  # Valid duration format
    delete = "15m"
  }
}
```