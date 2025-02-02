# `ibm_is_instance_network_interface`

This rule validates instance network interface configuration.

## Example

```hcl
resource "ibm_is_instance_network_interface" "example" {
  instance = ibm_is_instance.example.id
  subnet   = ibm_is_subnet.example.id
  name     = "nic-with-a-very-long-name-that-exceeds-maximum-allowed-length"
  primary_ipv4_address = "invalid-ip"
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid IPv4 address format (ibm_is_instance_network_interface)
  on main.tf line 5:
   5:   primary_ipv4_address = "invalid-ip"
```

## Why

Network interfaces require `instance`, `subnet`, and `name`. Names <=63 chars. Primary IPv4 must be valid format if specified.

## How To Fix

```hcl
resource "ibm_is_instance_network_interface" "example" {
  instance = ibm_is_instance.example.id
  subnet   = ibm_is_subnet.example.id
  name     = "nic-1"
  primary_ipv4_address = "192.168.1.10"  # Valid IPv4
}
```