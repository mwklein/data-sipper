# Data Sipper Data Source (Proof of Concept)
A simple proof of concept application connect to distributed data sources across multiple organizations, and send to a centralized database in a public cloud. This project provides an example application to extract data from a local datasource and send to the central server.

## Dev Environment Setup
For those new to GO, there are some prerequistes, required to setup your environment in order to build the Data Sipper application

First, GO requires a specific directory structure in order to build/compile your application

    |─── bin/
    |    |─── cmd1                                  # command executable
    |    |─── cmd2                                  # command executable
    |─── pkg/
    |    └─── linux_amd64/
    |         └─── github.com/golang/example/
    |              |─── stringutil.a                # package object
    |─── src/
    |         └─── github.com/golang/example/
    |              └─── .git/                       # Git repository metadata
    |         └─── cmd1/
    |              |─── cmd1.go                     # command source
    |         └─── cmd2/
    |              |─── cmd2.go
    |              |─── cmd2_test.go
    |         └─── github.com/mwklein/data-sipper/  # this GIT repository
    |              |─── main.go                     # main source code


You will need to set your environment variables to point to the root of the above structure.
```bash
export GOPATH=$HOME/path-to-workspace
export PATH=$PATH:$GOPATH/bin
export DATASIPPER_DB_HOSTNAME=<IP Address of your default docker machine>
```

To use the automated build and linting tools, you will need to install the following tools:
```bash
go get golang.org/x/tools/cmd/godoc
go get github.com/golang/lint/golint
go get golang.org/x/tools/cmd/vet
go get github.com/robertkrimen/godocdown/godocdown
go get -u gopkg.in/godo.v1/cmd/godo
```

## Installation
Within the $GOPATH/src/github.com/mwklein/data-sipper/ folder execute...
```bash
go get
go install
```

Or within the $GOPATH/
```bash
go get github.com/mwklein/data-sipper
go install github.com/mwklein/data-sipper
```

## Testing
Within $GOPATH/src/github.com/mwklein/data-sipper/ folder execute
```bash
go test
```

Or within the $GOPATH/
```bash
go test github.com/mwklein/data-sipper/
```

## Cross-compiling
There are many methods available to cross-compile the application for different OS and architectures. The method that has been tested on this project is:
```bash
go get github.com/karalabe/xgo
xgo .
```
This requires an active docker machine to be available in your environment, and downloads to a fairly large docker image to support the cross-compile. 