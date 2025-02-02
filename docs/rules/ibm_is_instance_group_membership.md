# `ibm_is_instance_group_membership`

This rule checks the configuration of instance group memberships.

## Example

```hcl
resource "ibm_is_instance_group_membership" "example" {
  name                     = "example-membership"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_membership = ibm_is_instance_group_membership.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_instance_group_membership)
  on main.tf line 1:
   1: resource "ibm_is_instance_group_membership" "example" {
```

## Why

The instance group membership configuration requires specific attributes and valid values to function properly. For example, the `name`, `instance_group`, and `instance_group_membership` attributes are required. Additionally, the `name` cannot be empty.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_membership" "example" {
  name                     = "example-membership"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_membership = ibm_is_instance_group_membership.example.id
}
```