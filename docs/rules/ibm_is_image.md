# `ibm_is_image`

## Example
```hcl
resource "ibm_is_image" "example" {
  name              = "example-image"
  href              = "cos://my-bucket/path/to/image.qcow2"
  operating_system  = "ubuntu-20-04-amd64"
}
```

```console
$ tflint
1 issue(s) found:
Error: href must be a valid Cloud Object Storage URL (ibm_is_image)
  on main.tf line 3:
   3:   href              = "invalid-url"
```

## Why
The `name`, `href`, and `operating_system` attributes are required to define a valid image. Additionally, the `href` must be a valid Cloud Object Storage URL.

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_image" "example" {
  name              = "example-image"
  href              = "cos://my-bucket/path/to/image.qcow2"
  operating_system  = "ubuntu-20-04-amd64"
}
```