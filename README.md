# go-gold
A simple implementation of the martian robots kata.
![Build](https://github.com/carlosroman/go-gold/workflows/Run%20tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/carlosroman/go-gold/badge.svg?branch=master)](https://coveralls.io/github/carlosroman/go-gold?branch=main)
---

## Table of Contents

- [Getting Started](#getting-started)
    - [Quick start](#quick-start)
    - [Running the application without Docker](#running-the-application-without-docker)
    - [Building standalone binary](#building-standalone-binary)
- [Development](#development)
- [TODO](#TODO)

## Getting Started

First clone this repo and cd into the directory:

```
git clone https://github.com/carlosroman/go-gold.git
cd go-gold
```

### Quick start

For the quickest start you just need:

* Make
* Docker

With those two items you can then just run the following to run the application:

```
make quick-start
```

This will read the file from [test/data/sample-transactions.csv](test/data/sample-transactions.csv) and output a file called [test/data/out.csv](test/data/out.csv)  

### Running the application without Docker

If you want to run the application without Docker, then you need the following:

* Make
* Golang 1.4+

The application can be run by using the following command

```
make start
```

If you want to override the input then run it as `make start INPUT=<path to file>`.
The output can be changed by set the env var `OUTPUT`, i.e. `make start OUTPUT=<path to file>`. 

### Building standalone binary

If you want to run the application without Docker, then you need the following:

* Make
* Golang 1.4+

```
make build
```

After the command has finished you'll have a binary called `gold` in directory called `bin`.
By default, this binary will be built to run on the local OS and architecture.

The binary takes two arguments:

* the input CSV as the first argument
* the output CSV as the second argument

When run it will calculate the data for six months from today.

## Development

The simplest way to run the unit tests for this project is:

```
make test
```

## TODO

* Set up the main command to take a date
* Use viper to have proper argument flags
