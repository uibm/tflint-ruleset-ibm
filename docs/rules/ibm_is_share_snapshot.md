# `ibm_is_share_snapshot`

## Example
```hcl
resource "ibm_is_share_snapshot" "example" {
  name  = "example-snapshot"
  share = ibm_is_share.example.id
  tags  = ["tag1", "tag2"]
}
```

```console
$ tflint
1 issue(s) found:
Error: tag length cannot exceed 128 characters (ibm_is_share_snapshot)
  on main.tf line 4:
   4:   tags  = ["tag1", "a-very-long-tag-that-exceeds-the-allowed-length-for-tags-in-ibm-cloud"]
```

## Why
The `name` and `share` attributes are required. Tags, if specified, must not exceed 128 characters each.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_share_snapshot" "example" {
  name  = "example-snapshot"
  share = ibm_is_share.example.id
  tags  = ["tag1", "tag2"]
}
```