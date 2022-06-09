# axolgo-cli, the Axolotl CLI Library in Golang
This is the CLI library of the Axolotl series in Golang. Command is 
designed to fit daily operational usage and the sub-command is 
added for better experience. You may configure **axolgo** through 
configuration file or command line parameters.

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

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-aws](https://github.com/tchiunam/axolgo-aws) for using AWS SDK
