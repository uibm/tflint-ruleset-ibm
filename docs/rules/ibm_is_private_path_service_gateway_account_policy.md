# `ibm_is_private_path_service_gateway_account_policy`

This rule checks the configuration of IBM Cloud private path service gateway account policies.

## Example

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_account_policy" "example" {
  access_policy                = "permit"
  account                      = "1234567890abcdef1234567890abcdef"
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `access_policy` attribute must be specified (ibm_is_private_path_service_gateway_account_policy)

  on main.tf line 1:
   1: resource "ibm_is_private_path_service_gateway_account_policy" "example" {
```

## Why

Private path service gateway account policy configuration requires the `access_policy`, `account`, and `private_path_service_gateway` attributes. The `access_policy` must be either `permit` or `deny`. The `account` must be a valid IBM Cloud account ID (typically 32 characters).

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  name = "example-gateway"
}

resource "ibm_is_private_path_service_gateway_account_policy" "example" {
  access_policy                = "permit" # Must be "permit" or "deny"
  account                      = "1234567890abcdef1234567890abcdef" # Valid account ID
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```