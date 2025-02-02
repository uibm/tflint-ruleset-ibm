# `ibm_is_ike_policy`

This rule checks the configuration of IBM Cloud IKE policies.

## Example

```hcl
resource "ibm_is_ike_policy" "example" {
  name                    = "example-ike-policy"
  authentication_algorithm = "sha256"
  encryption_algorithm   = "aes256"
  dh_group                = 14
  ike_version             = 2
  key_lifetime            = 28800
  resource_group          = ibm_resource_group.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_ike_policy)
  on main.tf line 1:
   1: resource "ibm_is_ike_policy" "example" {
```

## Why

IKE policy configuration requires specific attributes with valid values to ensure secure communication.  The `name`, `authentication_algorithm`, `encryption_algorithm`, `dh_group`, and `ike_version` attributes are required.  The `authentication_algorithm` must be one of `md5`, `sha1`, or `sha256`. The `encryption_algorithm` must be one of `triple_des`, `aes128`, `aes192`, or `aes256`. The `dh_group` must be one of `2`, `5`, `14`, or `19`. The `ike_version` must be either `1` or `2`. The optional `key_lifetime` attribute, if specified, must be between 300 and 86400 seconds.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_ike_policy" "example" {
  name                    = "example-ike-policy"
  authentication_algorithm = "sha256" # Must be one of: md5, sha1, sha256
  encryption_algorithm   = "aes256"   # Must be one of: triple_des, aes128, aes192, aes256
  dh_group                = 14         # Must be one of: 2, 5, 14, 19
  ike_version             = 2          # Must be either 1 or 2
  key_lifetime            = 28800      # Optional, must be between 300 and 86400
  resource_group          = ibm_resource_group.example.id
}
```