# `ibm_is_instance_group`

This rule ensures that the required attributes are specified for the `ibm_is_instance_group` resource.

## Example

```hcl
resource "ibm_is_instance_group" "example" {
  name              = "example-group"
  instance_template = ibm_is_instance_template.example.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.example.id]
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_instance_group)
  on main.tf line 1:
   1: resource "ibm_is_instance_group" "example" {
```

## Why

The `name`, `instance_template`, and `subnets` attributes are required to define a valid instance group. Missing these attributes will result in errors during Terraform apply. Additionally, the `instance_count` attribute, if specified, must be a non-negative integer.

## How To Fix

Ensure all required attributes are specified and valid:

```hcl
resource "ibm_is_instance_group" "example" {
  name              = "example-group"
  instance_template = ibm_is_instance_template.example.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.example.id]
}
```

---

# `ibm_is_instance_group_manager`

This rule checks the configuration of instance group managers.

## Example

```hcl
resource "ibm_is_instance_group_manager" "example" {
  name                 = "example-manager"
  instance_group       = ibm_is_instance_group.example.id
  manager_type         = "autoscale"
  enable_manager       = true
  max_membership_count = 10
  min_membership_count = 2
}
```

```console
$ tflint
1 issue(s) found:
Error: `manager_type` must be either 'autoscale' or 'scheduled' (ibm_is_instance_group_manager)
  on main.tf line 4:
   4:   manager_type         = "invalid"
```

## Why

The manager configuration requires specific attributes and valid values to function properly. For example, the `manager_type` must be either `autoscale` or `scheduled`. Additionally, the `max_membership_count` and `min_membership_count` must be logically consistent.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_manager" "example" {
  name                 = "example-manager"
  instance_group       = ibm_is_instance_group.example.id
  manager_type         = "autoscale"  # Must be 'autoscale' or 'scheduled'
  enable_manager       = true
  max_membership_count = 10
  min_membership_count = 2
}
```

---

# `ibm_is_instance_group_manager_policy`

This rule validates instance group manager policies.

## Example

```hcl
resource "ibm_is_instance_group_manager_policy" "example" {
  instance_group         = ibm_is_instance_group.example.id
  instance_group_manager = ibm_is_instance_group_manager.example.manager_id
  metric_type            = "cpu"
  metric_value           = 70
  policy_type            = "target"
  name                   = "example-policy"
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid metric_type. Must be one of: cpu, memory, network_in, network_out, custom (ibm_is_instance_group_manager_policy)
  on main.tf line 4:
   4:   metric_type            = "invalid"
```

## Why

Policies require specific attributes and valid values to function properly. For example, the `metric_type` must be one of the supported types (`cpu`, `memory`, `network_in`, `network_out`, `custom`), and the `policy_type` must be one of the supported types (`target`, `range`, `percentage`).

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_instance_group_manager_policy" "example" {
  instance_group         = ibm_is_instance_group.example.id
  instance_group_manager = ibm_is_instance_group_manager.example.manager_id
  metric_type            = "cpu"  # Must be one of: cpu, memory, network_in, network_out, custom
  metric_value           = 70
  policy_type            = "target"  # Must be one of: target, range, percentage
  name                   = "example-policy"
}
```

---

# Additional Notes

1. **Validation Scope**:
   - The `ibm_is_instance_group` rule ensures that required attributes (`name`, `instance_template`, `subnets`) are present.
   - The `ibm_is_instance_group_manager` rule validates the `manager_type` and ensures logical consistency between `max_membership_count` and `min_membership_count`.
   - The `ibm_is_instance_group_manager_policy` rule validates the `metric_type` and `policy_type` against their respective allowed values.

2. **Best Practices**:
   - Always specify all required attributes to avoid runtime errors.
   - Use valid values for enumerated fields like `manager_type`, `metric_type`, and `policy_type`.
   - Ensure that `max_membership_count` is greater than or equal to `min_membership_count`.

3. **References**:
   - [IBM Cloud Documentation for Instance Groups](https://cloud.ibm.com/docs/vpc?topic=vpc-instance-groups)
   - [TFLint Ruleset Documentation](https://github.com/uibm/tflint-ruleset-ibm)
