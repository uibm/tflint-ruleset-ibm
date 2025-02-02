# `ibm_is_share`

## Example
```hcl
resource "ibm_is_share" "example" {
  name     = "example-share"
  size     = 100
  profile  = "dp2"
  zone     = "us-south-1"
}
```

```console
$ tflint
1 issue(s) found:
Error: size must be between 10 and 16000 GB (ibm_is_share)
  on main.tf line 3:
   3:   size     = 5
```

## Why
The `name`, `size`, `profile`, and `zone` attributes are required. The `size` must be between 10 and 16000 GB, and the `profile` must be one of `dp2`, `dp4`, or `dp8`.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_share" "example" {
  name     = "example-share"
  size     = 100
  profile  = "dp2"
  zone     = "us-south-1"
}
```