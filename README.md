# Smallest Roman Numeral

API responsible for searching for the smallest prime roman number in a text.

## Summary

* [Dependencies](#dependencies)
* [Project Architecture](#project-architecture)
* [Environment](#environment)
* [Installing](#installing)
* [Running](#running)
* [Testing](#testing)
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

## Environment

## Installing

Cloning the Smallest Roman Numeral project:
``` bash
git clone https://github.com/GabrielFreitasP/smallest-roman-numeral.git
```

Access project folder:
``` bash
cd smallest-roman-numeral
```

Downloading dependencies, cleans up the unused dependencies and constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module:
```
make install
```

## Running

Building project:
``` bash
make build
```

Starting the application:
``` bash
make run
```

## Testing

Running the project's tests:
```bash
make test
```

Generating the test's coverage:
```bash
make coverage
```

## Documentation

Generating the swagger documentation:
```bash
make swagger
```

For more information, check the API documentation [here](http://localhost:8080/swagger/index.html#/RomanNumeral/post_search).
