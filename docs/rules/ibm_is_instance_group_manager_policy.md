# `ibm_is_instance_group_manager_policy`

This rule validates instance group manager policies.

## Example

```hcl
resource "ibm_is_instance_group_manager_policy" "example" {
  name                     = "example-policy"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_manager   = ibm_is_instance_group_manager.example.manager_id
  metric_type              = "cpu"
  metric_value             = 70
  policy_type              = "target"
}
```

```console
$ tflint
1 issue(s) found:
Error: `metric_type` attribute must be specified (ibm_is_instance_group_manager_policy)
  on main.tf line 1:
   1: resource "ibm_is_instance_group_manager_policy" "example" {
```

## Why

The policy configuration requires specific attributes and valid values to function properly. For example, the `name`, `instance_group`, `instance_group_manager`, `metric_type`, `metric_value`, and `policy_type` attributes are required. Additionally, the `metric_type` must be one of the supported types (`cpu`, `memory`, `network_in`, `network_out`, `custom`), and the `policy_type` must be one of the supported types (`target`, `range`, `percentage`).

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_manager_policy" "example" {
  name                     = "example-policy"
  instance_group           = ibm_is_instance_group.example.id
  instance_group_manager   = ibm_is_instance_group_manager.example.manager_id
  metric_type              = "cpu"  # Must be one of: cpu, memory, network_in, network_out, custom
  metric_value             = 70
  policy_type              = "target"  # Must be one of: target, range, percentage
}
```
