# Data Sipper Data Source (Proof of Concept)
A simple proof of concept application connect to distributed data sources across multiple organizations, and send to a centralized database in a public cloud. This project provides an example application to extract data from a local datasource and send to the central server.

## Installation
To build and install:
```bash
go mod download
go install
```

## Testing
To execute unit tests (requires DBs to be available):
```bash
go test
```

## Cross-compiling
There are many methods available to cross-compile the application for different OS and architectures. The method that has been tested on this project is:
```bash
go get github.com/karalabe/xgo
xgo .
```
This requires an active docker machine to be available in your environment, and downloads to a fairly large docker image to support the cross-compile. 