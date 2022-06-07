# Data Sipper Data Source (Proof of Concept)
A simple proof of concept application connect to distributed data sources across multiple organizations, and send to a centralized database in a public cloud. This project provides an example application to extract data from a local datasource and send to the central server.

## Installation
To build and install:
```bash
go mod download
go install
```

# Usage
You can view usage details with `data-sipper --help`:
```
Usage:
  data-sipper [OPTIONS]

Application Options:
  -c, --config-file=                           The configuration file for the data-sipper client
  -t, --db-type=[mysql|postgres|mssql|sqlite3] The type of database in which to connect (default: mysql) [$DATASIPPER_DB_TYPE]
  -n, --db-host=                               The hostname or IP adress of the database server (default: localhost) [$DATASIPPER_DB_HOSTNAME]
  -r, --db-port=                               The TCP port of the database server (default: 3306) [$DATASIPPER_DB_PORT]
  -f, --db-file-path=                          The file path to the database [$DATASIPPER_DB_FILE_PATH]
  -d, --db-name=                               The name of the database or cluster [$DATASIPPER_DB_NAME]
  -u, --db-username=                           The username used for connecting to the database [$DATASIPPER_DB_USERNAME]
  -p, --db-password=                           The password used for connecting to the database [$DATASIPPER_DB_PASSWORD]
  -q, --db-query=                              The SQL query to be executed on the database with results uploaded to server
  -s, --server=                                The URL of the server's REST endpoint [$DATASIPPER_SERVER_URL]

Help Options:
  -h, --help                                   Show this help message
```

## Testing
To execute unit tests (requires DBs to be available):
```bash
docker-compose up
go test
```

## Cross-compiling
The following is how to build the project for multiple architecture. This is method is currently untested:
```bash
# Build all
make

# Build linux
make linux

# Build mac
make darwin

# Build windows
make windows
```