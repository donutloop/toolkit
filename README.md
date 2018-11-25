# toolkit

[![Build Status](https://travis-ci.org/donutloop/toolkit.svg?branch=master)](https://travis-ci.org/donutloop/toolkit)
[![Coverage Status](https://coveralls.io/repos/github/donutloop/toolkit/badge.svg)](https://coveralls.io/github/donutloop/toolkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/donutloop/toolkit)](https://goreportcard.com/report/github.com/donutloop/toolkit)

## Introduction

These patterns can you use to solve common problems when designing an application or system.

## Requirements

* [golang](https://golang.org/) >=1.11.x - The Go Programming Language

## Installation

```sh
go get github.com/donutloop/toolkit
```

## Patterns 

* [Worker](https://github.com/donutloop/toolkit/blob/master/worker/README.md)
* [Job schedule](https://github.com/donutloop/toolkit/blob/master/schedule/README.md)
* [Singleton](https://github.com/donutloop/toolkit/blob/master/singleton/README.md)
* [Retry](https://github.com/donutloop/toolkit/blob/master/retry/README.md) 
* [Promise](https://github.com/donutloop/toolkit/blob/master/promise/README.md) 
* [Multierror](https://github.com/donutloop/toolkit/blob/master/multierror/README.md)
* [Loop](https://github.com/donutloop/toolkit/blob/master/loop/README.md) 
* [Lease](https://github.com/donutloop/toolkit/blob/master/lease/README.md)
* [Event-system](https://github.com/donutloop/toolkit/blob/master/event/README.md)
* [Debugutil](https://github.com/donutloop/toolkit/blob/master/debugutil/README.md)
* [Concurrent runner](https://github.com/donutloop/toolkit/blob/master/concurrent/README.md)
* [Bus-system](https://github.com/donutloop/toolkit/blob/master/bus/README.md)

## Examples 

In each sub directory is a set of examples 

## Code generation

The code generation tool generates for a pattern an none generic version for spefici type

### Supported pattern

* [Worker](https://github.com/donutloop/toolkit/blob/master/worker/README.md)

### Build

```bash
mkdir -p $GOPATH/src/github.com/donutloop/ && cd $GOPATH/src/github.com/donutloop/

git clone git@github.com:donutloop/toolkit.git

cd toolkit

go install ./cmd/xcode
```

### Usage

```bash
USAGE
  xcode [flags]

FLAGS
  -in     input file
  -out    output file
  -pkg    package name
  -type   type
```

#### Example generation 

```bash 
xcode -in $GOPATH/src/github.com/donutloop/toolkit/worker/worker.go -out $GOPATH/src/github.com/donutloop/toolkit/worker/new_worker.go -pkg test -type int32 
```

#### Example call for generated code 

```bash
workerHandler := func(v int32) {
    fmt.Println(v)
}

queue := worker.New(2, workerHandler, 10)

queue <- int32(3)
```

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from
anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to toolkit, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base to ensure those changes are in line with the general philosophy of the project and/or get some
early feedback which can make both your efforts much lighter as well as our review and merge
procedures quick and simple.

Please read and follow our [Contributing](https://github.com/donutloop/toolkit/blob/master/CONTRIBUTING.md).

## Code of Conduct

Please read and follow our [Code of Conduct](https://github.com/donutloop/toolkit/blob/master/CODE_OF_CONDUCT.md).
