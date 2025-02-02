# `ibm_is_volume`

This rule checks block storage volume configuration.

## Example

```hcl
resource "ibm_is_volume" "example" {
  profile = "custom"
  zone    = "us-south1" # Invalid zone format
}
```

```console
$ tflint
1 issue(s) found:
Error: `capacity` attribute must be specified (ibm_is_volume)
  on main.tf line 1:
   1: resource "ibm_is_volume" "example" {
```

## Why

Volume configuration requires `name`, `profile`, `zone`, and `capacity` attributes. Custom profiles need IOPS specified. Capacity must be 10-16,000 GB, IOPS 100-48,000 (if specified), and zone format must be region-number (e.g., us-south-1).

## How To Fix

```hcl
resource "ibm_is_volume" "example" {
  name     = "example-vol"  # <= 63 chars
  profile  = "custom"       # general-purpose/5iops-tier/10iops-tier
  zone     = "us-south-1"   # Valid zone
  capacity = 100            # 10-16000
  iops     = 1000           # Required for custom profile
}
```