# `ibm_is_networkacls`

## Example
```hcl
resource "ibm_is_network_acls" "example" {
  name = "example-acl"
  vpc  = ibm_is_vpc.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: name cannot be longer than 63 characters (ibm_is_networkacls)
  on main.tf line 2:
   2:   name = "this-name-is-way-too-long-and-exceeds-the-character-limit-for-ibm-cloud-resources"
```

## Why
The `name` and `vpc` attributes are required for defining network ACLs. Additionally, the `name` must not exceed 63 characters to comply with IBM Cloud naming conventions.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_network_acls" "example" {
  name = "example-acl"
  vpc  = ibm_is_vpc.example.id
}
```