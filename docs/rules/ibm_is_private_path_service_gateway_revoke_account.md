# `ibm_is_private_path_service_gateway_revoke_account`

This rule checks the configuration of IBM Cloud private path service gateway revoke account operations.

## Example

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_revoke_account" "example" {
  account                      = "1234567890abcdef1234567890abcdef"
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `account` attribute must be specified (ibm_is_private_path_service_gateway_revoke_account)

  on main.tf line 1:
   1: resource "ibm_is_private_path_service_gateway_revoke_account" "example" {
```

## Why

Private path service gateway revoke account configuration requires the `account` and `private_path_service_gateway` attributes. The `account` must be a valid IBM Cloud account ID (typically 32 characters).

## How To Fix

Ensure both required attributes are specified and that the `account` ID is valid:

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_revoke_account" "example" {
  account                      = "1234567890abcdef1234567890abcdef" # Valid account ID
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```