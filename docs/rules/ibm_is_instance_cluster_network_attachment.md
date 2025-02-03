# `ibm_is_instance_cluster_network_attachment`

This rule checks the configuration of IBM Cloud instance cluster network attachments.

## Example

```hcl
resource "ibm_is_instance" "example" {
  name = "example-instance"
}

resource "ibm_is_instance_cluster_network_attachment" "example" {
  instance_id = ibm_is_instance.example.id
  name        = "example-attachment"
  cluster_network_interface {
 name = "eth1"
 subnet {
   id = ibm_is_cluster_network_subnet.example.id
 }
  }
}
```

```console
$ tflint
1 issue(s) found:

Error: `instance_id` attribute must be specified (ibm_is_instance_cluster_network_attachment)

  on main.tf line 1:
   1: resource "ibm_is_instance_cluster_network_attachment" "example" {
```

## Why

Instance cluster network attachment configuration requires the `instance_id` and `name` attributes. The `name` must be between 1 and 63 characters. A `cluster_network_interface` block must be specified, and it must contain a `subnet` block with an `id` attribute.

## How To Fix

Ensure all required attributes and blocks are specified and that they adhere to the correct format:

```hcl
resource "ibm_is_instance" "example" {
  name = "example-instance"
}

resource "ibm_is_instance_cluster_network_attachment" "example" {
  instance_id = ibm_is_instance.example.id
  name        = "example-attachment" # Between 1 and 63 characters
  cluster_network_interface {
 name = "eth1"
 subnet {
   id = ibm_is_cluster_network_subnet.example.id
 }
  }
}
```