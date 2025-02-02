# `ibm_is_share_replica_operations`

## Example
```hcl
resource "ibm_is_share_replica_operations" "example" {
  share_replica = ibm_is_share_replica.example.id
  split_share   = true
}
```

```console
$ tflint
1 issue(s) found:
Error: cannot specify both split_share and fallback_policy (ibm_is_share_replica_operations)
  on main.tf line 4:
   4:   fallback_policy = "split"
```

## Why
The `share_replica` attribute is required, and either `split_share` or `fallback_policy` must be specified—but not both. Additionally, `fallback_policy` must be one of `split` or `failover`.

## How To Fix
Ensure only one of `split_share` or `fallback_policy` is specified:
```hcl
resource "ibm_is_share_replica_operations" "example" {
  share_replica = ibm_is_share_replica.example.id
  split_share   = true
}
```