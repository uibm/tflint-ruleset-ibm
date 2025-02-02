# `ibm_is_instance_group_manager_action`

This rule checks the configuration of instance group manager actions.

## Example

```hcl
resource "ibm_is_instance_group_manager_action" "example" {
  name                     = "example-action"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_manager   = ibm_is_instance_group_manager.example.manager_id
  cron_spec                = "0 0 * * *"
  min_membership_count     = 2
  max_membership_count     = 10
}
```

```console
$ tflint
1 issue(s) found:
Error: `cron_spec` attribute must be specified (ibm_is_instance_group_manager_action)
  on main.tf line 1:
   1: resource "ibm_is_instance_group_manager_action" "example" {
```

## Why

The instance group manager action configuration requires specific attributes and valid values to function properly. For example, the `name`, `instance_group`, `instance_group_manager`, and `cron_spec` attributes are required. Additionally, the `cron_spec` must be a valid cron expression, and the `max_membership_count` must be greater than or equal to the `min_membership_count`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_manager_action" "example" {
  name                     = "example-action"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_manager   = ibm_is_instance_group_manager.example.manager_id
  cron_spec                = "0 0 * * *"  # Must be a valid cron expression
  min_membership_count     = 2
  max_membership_count     = 10  # Must be greater than or equal to min_membership_count
}
```