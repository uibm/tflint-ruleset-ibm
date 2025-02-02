# `ibm_is_security_group_target`

## Example
```hcl
resource "ibm_is_security_group_target" "example" {
  security_group = ibm_is_security_group.example.id
  target         = ibm_is_instance.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: target attribute must be specified (ibm_is_security_group_target)
  on main.tf line 3:
   3:   target         = ibm_is_instance.example.id
```

## Why
The `security_group` and `target` attributes are required for defining security group targets. Missing these attributes will result in errors during Terraform apply.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_security_group_target" "example" {
  security_group = ibm_is_security_group.example.id
  target         = ibm_is_instance.example.id
}
```