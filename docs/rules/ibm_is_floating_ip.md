# `ibm_is_floating_ip`

This rule checks the configuration of floating IPs.

## Example

```hcl
resource "ibm_is_floating_ip" "example" {
  name = "example-floating-ip"
  zone = "us-south-1"
}
```

```console
$ tflint
1 issue(s) found:
Error: `zone` attribute must be specified (ibm_is_floating_ip)
  on main.tf line 1:
   1: resource "ibm_is_floating_ip" "example" {
```

## Why

The floating IP configuration requires specific attributes and valid values to function properly. For example, the `name` and `zone` attributes are required. Additionally, the `name` cannot be longer than 63 characters, and the `zone` must be in a valid format (e.g., `us-south-1`).

## How To Fix

Ensure all required attributes are specified with valid values:

```hcl
resource "ibm_is_floating_ip" "example" {
  name = "example-floating-ip"  # Cannot be empty or longer than 63 characters
  zone = "us-south-1"  # Must be a valid zone
}
```