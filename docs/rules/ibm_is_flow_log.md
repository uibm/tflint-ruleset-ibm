# `ibm_is_flow_log`

This rule checks the configuration of IBM Cloud flow logs.

## Example

```hcl
resource "ibm_is_flow_log" "example" {
  name           = "example-flow-log"
  target         = ibm_is_vpc.example.id
  active         = true
  storage_bucket = ibm_cos_bucket.example.bucket_name
  resource_group = ibm_resource_group.example.id
}
```

```console
$ tflint
1 issue(s) found:

Error: `name` attribute must be specified (ibm_is_flow_log)

  on main.tf line 1:
   1: resource "ibm_is_flow_log" "example" {
```

## Why

Flow log configuration requires the `name`, `target`, and `storage_bucket` attributes. The `name` must be between 1 and 63 characters. The `target` must be a valid target ID. The `storage_bucket` must be a valid COS bucket name.

## How To Fix

Ensure all required attributes are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_flow_log" "example" {
  name           = "example-flow-log" # Between 1 and 63 characters
  target         = ibm_is_vpc.example.id # Valid target ID
  active         = true
  storage_bucket = ibm_cos_bucket.example.bucket_name # Valid COS bucket name
  resource_group = ibm_resource_group.example.id
}
```