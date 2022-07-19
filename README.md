<img src="images/axolgo-logo-transparent.png" width="50%" />

# axolgo-cli, the Axolotl CLI Library in Golang
[![Go](https://github.com/tchiunam/axolgo-cli/actions/workflows/go.yml/badge.svg)](https://github.com/tchiunam/axolgo-cli/actions/workflows/go.yml)
[![Version](https://img.shields.io/github/v/release/tchiunam/axolgo-cli?sort=semver)](https://github.com/tchiunam/axolgo-cli/releases)
[![codecov](https://codecov.io/gh/tchiunam/axolgo-cli/branch/main/graph/badge.svg?token=R38VYBN1AL)](https://codecov.io/gh/tchiunam/axolgo-cli)
[![CodeQL](https://github.com/tchiunam/axolgo-cli/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/tchiunam/axolgo-cli/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tchiunam/axolgo-cli)](https://goreportcard.com/report/github.com/tchiunam/axolgo-cli)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tchiunam/axolgo-cli)](https://pkg.go.dev/github.com/tchiunam/axolgo-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli.svg?type=shield)](https://app.fossa.com/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli?ref=badge_shield)
[![Go Version](https://img.shields.io/github/go-mod/go-version/tchiunam/axolgo-cli)](https://github.com/tchiunam/axolgo-cli)
[![Repository Size](https://img.shields.io/github/repo-size/tchiunam/axolgo-cli.svg?label=Repo%20size)](https://github.com/tchiunam/axolgo-cli)

This is the CLI library of the Axolotl series in Golang. Command is 
designed to fit daily operational usage and the sub-command is 
added for better experience. You may configure **axolgo** through 
configuration file or command line parameters.

Go package: https://pkg.go.dev/github.com/tchiunam/axolgo-cli

## Use it with your Go module
To add as dependency for your package or upgrade to the latest version:
```
go get github.com/tchiunam/axolgo-cli
```

To upgrade or downgrade to a specific version:
```
go get github.com/tchiunam/axolgo-cli@v1.2.3
```

To remove dependency on your module and downgrade modules:
```
go get github.com/tchiunam/axolgo-cli@none
```

See 'go help get' or https://golang.org/ref/mod#go-get for details.

## Build
Download the source and run:
```
go build -o axolgo
```

See 'go help build' or https://golang.org/ref/mod#go-build for details.

## Install
To install latest version:
```
go install github.com/tchiunam/axolgo-cli@latest
```

To build and install version in module-aware mode:
```
go install github.com/tchiunam/axolgo-cli@v1.2.3
```

See 'go help install' or https://golang.org/ref/mod#go-install for details.

## Run test
To run test:
```
go test ./...
```

To run test with coverage result:
```
go test -coverpkg=./... ./...
```

## Examples
### AWS
To update database cluster parameter group:
```
axolgo aws rds modifyDBClusterParameterGroup --name <parameter_group_name> --parameter-file <yaml_file_containing_parameters>
```

To describe EC2 instances with given Instance IDs:
```
axolgo aws ec2 describeInstances --instance-id <instance_id> --instance-id <instance_id>
```

To describe EC2 instances with given Private IP Addresses:
```
axolgo aws ec2 describeInstances --private-ip-address 127.0.0.1 --private-ip-address 127.0.0.2
```
### GCP
To list all compute engine instances in a zone:
```
axolgo gcp compute listInstances --project proj1 --zone asia-east1-a
```

To list compute engine instances with the given ID:
```
axolgo gcp compute listInstances --project proj1 --zone asia-east1-a --id 7452065390813417482
```

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-cloud](https://github.com/tchiunam/axolgo-cloud) for using cloud library (AWS SDK and GCP API)

## License
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli.svg?type=large)](https://app.fossa.com/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli?ref=badge_large)