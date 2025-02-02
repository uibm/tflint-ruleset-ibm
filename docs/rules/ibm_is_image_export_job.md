# `ibm_is_image_export_job`

## Example
```hcl
resource "ibm_is_image_export_job" "example" {
  image = ibm_is_image.example.id
  name  = "export-job"
  format = "qcow2"

  storage_bucket {
    name = "my-bucket"
  }
}
```

```console
$ tflint
1 issue(s) found:
Error: name cannot be empty (ibm_is_image_export_job)
  on main.tf line 3:
   3:   name  = ""
```

## Why
The `image`, `name`, and `storage_bucket` block are required attributes for exporting an image. Additionally, the `format` must be one of the supported types (`qcow2`, `raw`, `vhd`, `vmdk`).

## How To Fix
Ensure all required attributes are specified with valid values:
```hcl
resource "ibm_is_image_export_job" "example" {
  image = ibm_is_image.example.id
  name  = "export-job"
  format = "qcow2"

  storage_bucket {
    name = "my-bucket"
  }
}
```