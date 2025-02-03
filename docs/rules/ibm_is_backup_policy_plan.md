# `ibm_is_backup_policy_plan`

This rule checks the configuration of IBM Cloud backup policy plans.

## Example

```hcl
resource "ibm_is_backup_policy" "example" {
  name = "example-backup-policy"
}

resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = ibm_is_backup_policy.example.id
  cron_spec        = "0 0 * * *"
  name             = "example-backup-policy-plan"
  clone_policy {
 zones = ["us-south-1", "us-east-1"]
 max_snapshots = 10
  }
}
```

```console
$ tflint
1 issue(s) found:

Error: `backup_policy_id` attribute must be specified (ibm_is_backup_policy_plan)

  on main.tf line 1:
   1: resource "ibm_is_backup_policy_plan" "example" {
```

## Why

Backup policy plan configuration requires the `backup_policy_id`, `cron_spec`, and `name` attributes. The `cron_spec` must be a valid cron expression. The `name` must be between 1 and 63 characters. If a `clone_policy` block is present, it must include the `zones` and `max_snapshots` attributes. The `zones` attribute cannot be empty and must contain valid zone formats (e.g., `region-number`, like `us-south-1`). The `max_snapshots` attribute must be greater than 0.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_backup_policy" "example" {
  name = "example-backup-policy"
}

resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = ibm_is_backup_policy.example.id
  cron_spec        = "0 0 * * *" # Valid cron expression
  name             = "example-backup-policy-plan" # Between 1 and 63 characters
  clone_policy {
 zones = ["us-south-1", "us-east-1"] # Valid zone formats
 max_snapshots = 10 # Greater than 0
  }
}
```