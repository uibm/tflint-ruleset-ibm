# `ibm_is_security_group_rule`

## Example
```hcl
resource "ibm_is_security_group_rule" "example" {
  group      = ibm_is_security_group.example.id
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: only one protocol block (icmp, tcp, or udp) can be specified (ibm_is_security_group_rule)
  on main.tf line 7:
   7:   udp {
```

## Why
Only one protocol block (`icmp`, `tcp`, or `udp`) can be specified per security group rule. Additionally, port ranges must be valid.

## How To Fix
Ensure only one protocol block is specified with valid values:
```hcl
resource "ibm_is_security_group_rule" "example" {
  group      = ibm_is_security_group.example.id
  direction  = "inbound"
  remote     = "0.0.0.0/0"

  tcp {
    port_min = 80
    port_max = 80
  }
}
```