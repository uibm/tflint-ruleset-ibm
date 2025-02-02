# Rules

This documentation describes a list of rules available by enabling this ruleset for IBM Cloud resources.

## Possible Errors

These rules warn of possible errors that can occur during `terraform apply`. Rules marked with `Deep` are only used when enabling [Deep Checking](../deep_checking.md).

| Rule                                | Description                                      | Deep | Enabled by Default |
|-------------------------------------|--------------------------------------------------|------|--------------------|
| ibm_is_instance_invalid_profile     | Disallow using invalid instance profiles         | ✔    | ✔                  |
| ibm_is_vpc_invalid_name             | Disallow using invalid VPC names                 |      | ✔                  |
| ibm_is_backup_policy_invalid_match  | Disallow using invalid match criteria            |      | ✔                  |
| ibm_is_floating_ip_invalid_target   | Disallow using invalid floating IP targets       | ✔    | ✔                  |
| ibm_is_security_group_invalid_rule  | Disallow using invalid security group rules      |      | ✔                  |

## Best Practices/Naming Conventions

These rules enforce best practices and naming conventions for IBM Cloud resources.

| Rule                                | Description                                      | Enabled by Default |
|-------------------------------------|--------------------------------------------------|--------------------|
| ibm_is_instance_previous_type       | Disallow using previous generation instance types| ✔                  |
| ibm_is_vpc_default_resource_group   | Disallow using default resource groups           | ✔                  |

## SDK-based Validations

Rules based on IBM Cloud SDK validations are also available:

| Rule                                | Enabled by Default |
|-------------------------------------|--------------------|
| ibm_is_instance_invalid_zone        | ✔                  |
| ibm_is_vpc_invalid_region           | ✔                  |
| ibm_is_backup_policy_invalid_name   | ✔                  |

For detailed information about each rule, refer to its dedicated documentation.