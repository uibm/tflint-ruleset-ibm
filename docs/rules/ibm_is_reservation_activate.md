# `ibm_is_reservation_activate`

## Example
```hcl
resource "ibm_is_reservation_activate" "example" {
  reservation = ibm_is_reservation.example.id
}
```

```console
$ tflint
1 issue(s) found:
Error: reservation attribute must be specified (ibm_is_reservation_activate)
  on main.tf line 1:
   1: resource "ibm_is_reservation_activate" "example" {
```

## Why
The `reservation` attribute is required to specify which reservation should be activated. Missing this attribute will result in errors during Terraform apply.

## How To Fix
Ensure the `reservation` attribute is specified with a valid IBM Cloud reservation ID:
```hcl
resource "ibm_is_reservation_activate" "example" {
  reservation = ibm_is_reservation.example.id
}
```