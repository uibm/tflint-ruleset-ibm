# `ibm_is_image_deprecate`

## Example
```hcl
resource "ibm_is_image_deprecate" "example" {
  image = ibm_is_image.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: image attribute must be specified (ibm_is_image_deprecate)
  on main.tf line 1:
   1: resource "ibm_is_image_deprecate" "example" {
```

## Why
The `image` attribute is required to specify which image should be deprecated. Missing this attribute will result in errors during Terraform apply.

## How To Fix
Ensure the `image` attribute is specified with a valid IBM Cloud image ID:
```hcl
resource "ibm_is_image_deprecate" "example" {
  image = ibm_is_image.example.id
}
```