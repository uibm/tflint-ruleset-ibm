# `ibm_is_network_acl_rule`

## Example
```hcl
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "allow-http"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"

  tcp {
    port_min = 80
    port_max = 80
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: action must be either 'allow' or 'deny' (ibm_is_network_acl_rule)
  on main.tf line 5:
   5:   action      = "invalid-action"
```

## Why
The `action` must be either `allow` or `deny`. Additionally, protocol-specific blocks (`tcp`, `udp`, `icmp`) require valid port ranges or codes.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "allow-http"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"

  tcp {
    port_min = 80
    port_max = 80
  }
}
```