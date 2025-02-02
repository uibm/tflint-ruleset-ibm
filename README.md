# TFLint Ruleset for Terraform Provider IBM Cloud

[![Build Status](https://github.com/uibm/tflint-ruleset-ibm/workflows/build/badge.svg?branch=master)](https://github.com/uibm/tflint-ruleset-ibm/actions)
[![GitHub release](https://img.shields.io/github/release/uibm/tflint-ruleset-ibm.svg)](https://github.com/uibm/tflint-ruleset-ibm/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

This is a TFLint ruleset plugin for the [Terraform IBM Cloud Provider](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest). The ruleset focuses on identifying possible errors, best practices, and misconfigurations specific to IBM Cloud resources.

Many rules are enabled by default and warn against code that might fail during `terraform apply`, or configurations that are not recommended.

---

## Requirements

- **TFLint**: v0.42+
- **Go**: v1.23+

---

## Installation

You can install the plugin by adding the following configuration to your `.tflint.hcl` file and running `tflint --init`:

```hcl
plugin "ibm" {
    enabled = true
    version = "0.1.0"
    source  = "github.com/uibm/tflint-ruleset-ibm"
}
```

For more details about configuring the plugin, see [Plugin Configuration](docs/configuration.md).

---

## Getting Started

Terraform is an excellent tool for Infrastructure as Code (IaC), but it doesn't validate provider-specific issues. For example, consider the following Terraform configuration:

```hcl
resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "invalid-image-id" # Invalid image ID!
  profile = "invalid-profile" # Invalid instance profile!
  primary_network_interface {
    subnet = ibm_is_subnet.subnet.id
  }
  vpc    = ibm_is_vpc.vpc.id
  zone   = "us-south-1"
}
```

In this example:
- `invalid-image-id` is not a valid image ID.
- `invalid-profile` is not a valid instance profile.

While Terraform (`terraform validate` and `terraform plan`) will not detect these issues, they will cause errors during `terraform apply`. This is where **TFLint with the IBM Cloud ruleset** comes in.

By running TFLint with this ruleset, you can catch such issues early in your development workflow, preventing them from causing failures in production CI/CD pipelines.

![demo](docs/assets/demo.gif)

---

## Rules

The following rules are currently implemented:

### Instance Rules
- **`ibm_is_instance`**: Validates that the `profile`, `image` attribute of `ibm_is_instance`. 

### VPC Rules
- **`ibm_is_vpc_name`**: Ensures that the `name` attribute of `ibm_is_vpc` is specified.

More rules will be added in future releases. For a complete list of rules, see [Rules](docs/rules/README.md).

---

## Building the Plugin

To build the plugin locally, clone the repository and run the following commands:

```bash
$ git clone https://github.com/uibm/tflint-ruleset-ibm.git
$ cd tflint-ruleset-ibm
$ make
```

You can easily install the built plugin using:

```bash
$ make install
```

> **Note**: If you install the plugin with `make install`, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "ibm" {
    enabled = true
}
```

---

## Contributing

We welcome contributions to improve this ruleset! Here's how you can help:
- **Report Issues**: If you find a bug or have a feature request, please open an issue.
- **Submit Pull Requests**: Add new rules, fix bugs, or improve documentation.
- **Improve Documentation**: Help us enhance the documentation for better clarity.

For more details, see [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---