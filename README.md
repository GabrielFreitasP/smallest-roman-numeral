# Smallest Roman Numeral

API responsible for searching for the smallest prime roman numeral in a text.

## Summary

* [Dependencies](#dependencies)
* [Project Architecture](#project-architecture)
* [Environment](#environment)
* [Installing](#installing)
* [Running](#running)
* [Testing](#testing)
* [Linting](#linting)
* [Documentation](#documentation)

## Dependencies

- [Golang](https://golang.org/) (tested with 1.19.3 on linux)
- [GNU Make](https://www.gnu.org/software/make/) (tested with 4.3 on linux)

## Project Architecture

```
smallest-roman-numeral
  |-- cmd
    |-- api
  |-- config
  |-- docs
  |-- internal
    |-- {module_name}
      |-- delivery
      |-- usecase
    |-- middleware
    |-- models
    |-- server
  |-- pkg
```

## Configuration

Environment variables are defined in yaml files in the project settings directory. The configuration file to run the API
in the local environment is already included in the project and can be found by the path "./config/config-local". If it
is necessary to create another file, follow all the properties with their respective descriptions:

```yaml
server:
  AppVersion: string # Application version (project version), for example: 1.0.0
  Port: string # Application server port number (it must follow the following pattern: ":0000"
  Mode: string # Application mode (Development, Staging or Production)
  ReadTimeout: number # Timeout to read API request
  WriteTimeout: number # Timeout to write API request

logger:
  Encoding: string # Type of log display (console or json)
  Level: string # Log information level (debug, info, warn, error, panic or fatal)

metrics:
  Url: string # URL to run metrics service
  ServiceName: string # Metrics service name
```

## Installing

Clones the Smallest Roman Numeral project:

```bash
git clone https://github.com/GabrielFreitasP/smallest-roman-numeral.git
```

Access project folder:

```bash
cd smallest-roman-numeral
```

Downloads dependencies, cleans up the unused dependencies and constructs a directory named vendor in the main module's
root directory that contains copies of all packages needed to support builds and tests of packages in the main module:

```bash
make install
```

## Running

Builds project:

```bash
make build
```

Starts the application:

```bash
make run
```

## Testing

Generates the mocks (requires [mockgen](https://github.com/golang/mock)):

```bash
make mock
```

Runs the project's tests

```bash
make test
```

Generates the test coverage:

```bash
make coverage
```

## Analyzing

Analyzes the project code (requires [golangci-lint](https://github.com/golangci/golangci-lint)):

```bash
make lint
```

## Documentation

Generates the swagger documentation (requires [swag](https://github.com/swaggo/swag))

```bash
make swagger
```

For more information, check the API
documentation [here](http://localhost:8080/swagger/index.html#) (the server must be running).
