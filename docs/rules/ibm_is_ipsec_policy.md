# `ibm_is_ipsec_policy`

This rule checks IPSec policy configuration.

## Example

```hcl
resource "ibm_is_ipsec_policy" "example" {
  name                    = "policy-1"
  authentication_algorithm = "sha3"  # Invalid
  pfs                     = "group_1" # Invalid
}
```

```console
$ tflint
1 issue(s) found:
Error: authentication_algorithm must be one of: md5, sha1, sha256, sha384, sha512 (ibm_is_ipsec_policy)
  on main.tf line 3:
   3:   authentication_algorithm = "sha3"
```

## Why

IPSec policies require `name`, `authentication_algorithm`, `encryption_algorithm`, and `pfs`. Valid authentication algorithms are md5/sha1/sha256/sha384/sha512. Encryption algorithms include triple_des/aes128/aes192/aes256. PFS must be disabled/group_2/group_5/group_14. Key lifetime must be 300-86400 seconds if specified.

## How To Fix

```hcl
resource "ibm_is_ipsec_policy" "example" {
  name                    = "policy-1"
  authentication_algorithm = "sha256"
  encryption_algorithm     = "aes256"
  pfs                     = "group_14"
  key_lifetime            = 3600  # 300-86400
}
```