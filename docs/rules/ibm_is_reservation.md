# `ibm_is_reservation`

## Example
```hcl
resource "ibm_is_reservation" "example" {
  zone = "us-south-1"

  capacity {
    total = 100
  }

  committed_use {
    term = "one_year"
  }

  profile {
    name          = "bx2-2x8"
    resource_type = "instance_profile"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: term must be either 'one_year' or 'three_year' (ibm_is_reservation)
  on main.tf line 9:
   9:     term = "invalid-term"
```

## Why
The `zone`, `capacity`, `committed_use`, and `profile` blocks are required for defining reservations. The `term` must be either `one_year` or `three_year`.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_reservation" "example" {
  zone = "us-south-1"

  capacity {
    total = 100
  }

  committed_use {
    term = "one_year"
  }

  profile {
    name          = "bx2-2x8"
    resource_type = "instance_profile"
  }
}
```