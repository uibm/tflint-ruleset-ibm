# `ibm_is_instance_group_manager`

This rule checks the configuration of instance group managers.

## Example

```hcl
resource "ibm_is_instance_group_manager" "example" {
  name                     = "example-manager"
  instance_group           = ibm_is_instance_group.example.id
  manager_type             = "autoscale"
  enable_manager           = true
  max_membership_count     = 10
  min_membership_count     = 2
}
```

```console
$ tflint
1 issue(s) found:
Error: `manager_type` must be either 'autoscale' or 'scheduled' (ibm_is_instance_group_manager)
  on main.tf line 4:
   4:   manager_type             = "invalid"
```

## Why

The instance group manager configuration requires specific attributes and valid values to function properly. For example, the `name`, `instance_group`, `manager_type`, and `enable_manager` attributes are required. Additionally, the `manager_type` must be either `autoscale` or `scheduled`, and the `max_membership_count` must be greater than or equal to the `min_membership_count`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_manager" "example" {
  name                     = "example-manager"
  instance_group           = ibm_is_instance_group.example.id
  manager_type             = "autoscale"  # Must be either 'autoscale' or 'scheduled'
  enable_manager           = true
  max_membership_count     = 10
  min_membership_count     = 2  # Must be less than or equal to max_membership_count
}
```