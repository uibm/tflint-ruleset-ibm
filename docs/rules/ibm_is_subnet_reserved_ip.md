# `ibm_is_subnet_reserved_ip`

This rule checks the configuration of subnet reserved IPs.

## Example

```hcl
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet  = ibm_is_subnet.example.id
  address = "invalid-ip-format"
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid IPv4 address format (ibm_is_subnet_reserved_ip)
  on main.tf line 3:
   3:   address = "invalid-ip-format"
```

## Why

The reserved IP configuration requires at least the `subnet` attribute. When specifying an IP address, it must be a valid IPv4 format. Names are optional but must be <= 63 characters if provided.

## How To Fix

Ensure valid configuration with required attributes:

```hcl
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet  = ibm_is_subnet.example.id
  name    = "reserved-ip-1"  # <= 63 characters if specified
  address = "192.168.1.100"  # Valid IPv4 format
  target  = ibm_is_instance.example.primary_network_interface.0.id
}
```