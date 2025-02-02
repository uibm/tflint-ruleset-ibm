# `ibm_is_public_gateway`

## Example
```hcl
resource "ibm_is_public_gateway" "example" {
  name = "example-gateway"
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
}
```

```console
$ tflint
1 issue(s) found:
Error: invalid zone format. Must be in format: region-number (e.g., us-south-1) (ibm_is_public_gateway)
  on main.tf line 4:
   4:   zone = "invalid-zone"
```

## Why
The `name`, `vpc`, and `zone` attributes are required for defining public gateways. The `zone` must follow the format `region-number`.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_public_gateway" "example" {
  name = "example-gateway"
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
}
```