# `ibm_is_instance_disk_management`

This rule checks the configuration of IBM Cloud instance disk management.

## Example

```hcl
resource "ibm_is_instance" "example" {
  name = "example-instance"
}

resource "ibm_is_instance_disk_management" "example" {
  instance = ibm_is_instance.example.id
  disks {
 id = ibm_is_instance_disk.example.id
 name = "example-disk"
  }
}
```

```console
$ tflint
1 issue(s) found:

Error: instance attribute must be specified (ibm_is_instance_disk_management)

  on main.tf line 1:
   1: resource "ibm_is_instance_disk_management" "example" {
```

## Why

Instance disk management configuration requires the `instance` attribute and at least one `disks` block. Each `disks` block must have an `id` attribute. The `name` attribute within a `disks` block is optional but cannot be empty and must be no longer than 63 characters.

## How To Fix

Ensure the `instance` attribute and at least one `disks` block are specified, and that the `name` attribute within each `disks` block (if provided) adheres to the format:

```hcl
resource "ibm_is_instance" "example" {
  name = "example-instance"
}

resource "ibm_is_instance_disk_management" "example" {
  instance = ibm_is_instance.example.id
  disks {
 id = ibm_is_instance_disk.example.id
 name = "example-disk" # Optional, but if provided, must be <= 63 characters
  }
}
```