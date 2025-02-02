# `ibm_is_backup_policy`

This rule ensures that the required attributes are specified for the `ibm_is_backup_policy` resource.

## Example

```hcl
resource "ibm_is_backup_policy" "example" {
  name                = "example-backup-policy"
  match_resource_type = "volume"
  match_user_tags     = ["env:prod"]
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_backup_policy)
  on main.tf line 1:
   1: resource "ibm_is_backup_policy" "example" {
```

## Why

The `name` and `match_resource_type` attributes are required to define a valid backup policy. Missing these attributes will result in errors during Terraform apply. Additionally, the `match_resource_type` attribute cannot be empty.

## How To Fix

Ensure all required attributes are specified and valid:

```hcl
resource "ibm_is_backup_policy" "example" {
  name                = "example-backup-policy"
  match_resource_type = "volume"
  match_user_tags     = ["env:prod"]
}
```