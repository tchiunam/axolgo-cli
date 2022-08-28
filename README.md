<img src="images/axolgo-logo-transparent.png" width="50%" />

# axolgo-cli, the Axolotl CLI Library in Golang
#### Release
<div align="left">
  <a href="https://github.com/tchiunam/axolgo-cli/releases">
    <img alt="Version" src="https://img.shields.io/github/v/release/tchiunam/axolgo-cli?sort=semver" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cli/releases">
    <img alt="Release Date" src="https://img.shields.io/github/release-date/tchiunam/axolgo-cli" />
  </a>
  <a href="https://pkg.go.dev/github.com/tchiunam/axolgo-cli">
    <img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/tchiunam/axolgo-cli" />
  </a>
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/tchiunam/axolgo-cli" />
  <img alt="Language" src="https://img.shields.io/github/languages/count/tchiunam/axolgo-cli" />
  <img alt="File Count" src="https://img.shields.io/github/directory-file-count/tchiunam/axolgo-cli" />
  <img alt="Repository Size" src="https://img.shields.io/github/repo-size/tchiunam/axolgo-cli.svg?label=Repo%20size" />
</div>

#### Code Quality
<div align="left">
  <a href="https://github.com/tchiunam/axolgo-cli/actions/workflows/go.yml">
    <img alt="Go" src="https://github.com/tchiunam/axolgo-cli/actions/workflows/go.yml/badge.svg" />
  </a>
  <a href="https://codecov.io/gh/tchiunam/axolgo-cli">
    <img alt="codecov" src="https://codecov.io/gh/tchiunam/axolgo-cli/branch/main/graph/badge.svg?token=R38VYBN1AL" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cli/actions/workflows/codeql-analysis.yml">
    <img alt="CodeQL" src="https://github.com/tchiunam/axolgo-cli/actions/workflows/codeql-analysis.yml/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/tchiunam/axolgo-cli">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/tchiunam/axolgo-cli" />
  </a>
</div>

#### Activity
<div align="left">
  <a href="https://github.com/tchiunam/axolgo-cli/commits/main">
    <img alt="Last Commit" src="https://img.shields.io/github/last-commit/tchiunam/axolgo-cli" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cli/issues?q=is%3Aissue+is%3Aclosed">
    <img alt="Closed Issues" src="https://img.shields.io/github/issues-closed/tchiunam/axolgo-cli" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cli/pulls?q=is%3Apr+is%3Aclosed">
    <img alt="Closed Pull Requests" src="https://img.shields.io/github/issues-pr-closed/tchiunam/axolgo-cli" />
  </a>
</div>

#### License
<div align="left">
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/github/license/tchiunam/axolgo-cli" />
  </a>
  <a href="https://app.fossa.com/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli?ref=badge_shield">
    <img alt="FOSSA Status" src="https://app.fossa.com/api/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli.svg?type=shield" />
  </a>
</div>

<br />
This is the CLI library of the Axolotl series in Golang. Command is 
designed to fit daily operational usage and the sub-command is 
added for better experience. You may configure **axolgo** through 
configuration file or command line parameters.

Go package: https://pkg.go.dev/github.com/tchiunam/axolgo-cli

## Use it with your Go module
To add as dependency for your package or upgrade to the latest version:
```console
go get github.com/tchiunam/axolgo-cli
```

To upgrade or downgrade to a specific version:
```console
go get github.com/tchiunam/axolgo-cli@v1.2.3
```

To remove dependency on your module and downgrade modules:
```console
go get github.com/tchiunam/axolgo-cli@none
```

See 'go help get' or https://golang.org/ref/mod#go-get for details.

## Build
Download the source and run:
```console
go build -o axolgo
```

See 'go help build' or https://golang.org/ref/mod#go-build for details.

## Install
To install latest version:
```console
go install github.com/tchiunam/axolgo-cli@latest
```

To build and install version in module-aware mode:
```console
go install github.com/tchiunam/axolgo-cli@v1.2.3
```

See 'go help install' or https://golang.org/ref/mod#go-install for details.

## Run test
To run test:
```console
go test ./...
```

To run test with coverage result:
```console
go test -coverpkg=./... ./...
```

## Examples
### AWS
To update database cluster parameter group:
```console
axolgo aws rds modifyDBClusterParameterGroup --name <parameter_group_name> --parameter-file <yaml_file_containing_parameters>
```

To describe EC2 instances with given Instance IDs:
```console
axolgo aws ec2 describeInstances --instance-id <instance_id> --instance-id <instance_id>
```

To describe EC2 instances with given Private IP Addresses:
```console
axolgo aws ec2 describeInstances --private-ip-address 127.0.0.1 --private-ip-address 127.0.0.2
```
### GCP
To list all compute engine instances in a zone:
```console
axolgo gcp compute listInstances --project proj1 --zone asia-east1-a
```

To list compute engine instances with the given ID:
```console
axolgo gcp compute listInstances --project proj1 --zone asia-east1-a --id 7452065390813417482
```

### Cryptography
To encrypt a message:
```console
axolgo cryptography encrypt --key-file secret.key --message "Galaxy is a big space where there are many planets and stars."
```

To decrypt a message:
```console
axolgo cryptography decrypt --key-file secret.key --message 5602259f17f41c7a1289d26c4cc540bdfe939a93d4f01f4f6830739f2488c04e7875890280d966a3509217b73e290c2bcdbbf2b8676395bd87ae000dcac8343ac3cf3e389968e185f96810302947ea6809d62dfee404e318d2
```


## Test report
## Code Coverage graph
![Code Coverage graph](https://codecov.io/gh/tchiunam/axolgo-cli/branch/main/graphs/tree.svg?token=R38VYBN1AL)

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-cloud](https://github.com/tchiunam/axolgo-cloud) for using cloud library (AWS SDK and GCP API)

## License
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli.svg?type=large)](https://app.fossa.com/projects/custom%2B32310%2Fgithub.com%2Ftchiunam%2Faxolgo-cli?ref=badge_large)