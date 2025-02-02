# `ibm_is_ssh_key`

This rule checks the configuration of SSH keys.

## Example

```hcl
resource "ibm_is_ssh_key" "example" {
  name = "example-key"
  public_key = "invalid-key-format"
  type = "dsa"
}
```

```console
$ tflint
1 issue(s) found:
Error: type must be either 'rsa' or 'ed25519' (ibm_is_ssh_key)
  on main.tf line 4:
   4:   type = "dsa"
```

## Why

The SSH key configuration requires specific attributes and valid values to function properly. The `name` and `public_key` attributes are required. Names cannot be empty or exceed 63 characters. Public keys must be in valid SSH format, and key type must be either `rsa` or `ed25519`.

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_ssh_key" "example" {
  name       = "example-key"  # Cannot be empty or >63 chars
  public_key = "ssh-rsa AAAAB3NzaC1yc2E..."  # Valid SSH public key
  type       = "rsa"  # Must be 'rsa' or 'ed25519'
}
```