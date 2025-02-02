# `ibm_is_instance_network_attachment`

This rule checks instance network attachment configuration.

## Example

```hcl
resource "ibm_is_instance_network_attachment" "example" {
  instance = ibm_is_instance.example.id
  # Missing name and virtual_network_interface
}
```

```console
$ tflint
1 issue(s) found:
Error: `name` attribute must be specified (ibm_is_instance_network_attachment)
  on main.tf line 1:
   1: resource "ibm_is_instance_network_attachment" "example" {
```

## Why

Network attachments require `instance`, `name`, and `virtual_network_interface` block with `id`. Names must be 1-63 characters.

## How To Fix

```hcl
resource "ibm_is_instance_network_attachment" "example" {
  instance = ibm_is_instance.example.id
  name     = "attachment-1"
  
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.example.id
  }
}
```