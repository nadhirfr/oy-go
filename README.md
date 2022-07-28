# OY! Indonesia API Go Client

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/nadhirfr/oy-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/nadhirfr/oy-go)](https://goreportcard.com/report/github.com/nadhirfr/oy-go)

This library is the abstraction of OY! Indonesia API for access from applications written with Go.

- [Documentation](#documentation)
- [Installation](#installation)
  - [Go Module Support](#go-module-support)
- [Usage](#usage)
  - [Without Client](#without-client)
  - [With Client](#with-client)
  - [Sub-Packages Documentations](#sub-packages-documentations)
- [Contribute](#contribute)
- [License](#license)



## Documentation

For the API documentation, check [OY! API Reference](https://api-docs.oyindonesia.com).

For the details of this library, see the [GoDoc](https://pkg.go.dev/github.com/nadhirfr/oy-go).

## Installation

Install oy-go with:

```sh
go get -u github.com/nadhirfr/oy-go
```

Then, import it using:

```go
import (
    oygo "github.com/nadhirfr/oy-go"
    "github.com/nadhirfr/oy-go/$product$"
)
```

with `$product$` is the product of OY! such as `checkout` and `balance`.


### Go Module Support

This library supports Go modules by default. Simply require oy-go in `go.mod` with a version like so:

```go
module github.com/my/package

go 1.17

require (
  github.com/nadhirfr/oy-go v1.0.0
)
```

And use the same style of import paths as above:

```go
import (
    oygo "github.com/nadhirfr/oy-go"
    "github.com/nadhirfr/oy-go/$product$"
)
```

with `$product$` is the product of OY! such as `checkout` and `balance`.


## Usage

It is recommended you use **With Client** method. The following pattern is applied throughout the library for a given `$product$`:

### Without Client

If you're only dealing with a single `secret key`, you can simply import the packages required for the products you're interacting with without the need to create a client.

```go
import (
    oygo "github.com/nadhirfr/oy-go"
    "github.com/nadhirfr/oy-go/$product$"
)

// Setup
oygo.Opt.Username = "exampleusername"
oygo.Opt.SecretKey = "examplesecretkey"


// Create
resp, err := $product$.Create($product$.CreateParams)

// Get
resp, err := $product$.Get($product$.GetParams)

// GetAll
resp, err := $product$.GetAll($product$.GetAllParams)
```

### With Client

If you're dealing with multiple `secret key`s, it is recommended you use `client.API`. This allows you to create as many clients as needed, each with their own individual key.

```go
import (
    oygo "github.com/nadhirfr/oy-go"
    "github.com/nadhirfr/oy-go/client"
)

// Basic setup
oyClient := client.New("examplesecretkey", "exampleusername")

// Create
resp, err := oyClient.$product$.Create($product$.CreateParams)

// Get
resp, err := oyClient.$product$.Get($product$.GetParams)

// GetAll
resp, err := oyClient.$product$.GetAll($product$.GetAllParams)
```

### Sub-Packages Documentations

The following is a list of pointers to documentations for sub-packages of [oy-go](https://github.com/nadhirfr/oy-go).

- [Balance](https://pkg.go.dev/github.com/nadhirfr/oy-go/balance)
- [Disbursement](https://pkg.go.dev/github.com/nadhirfr/oy-go/disbursement)
- [Scheduled Disbursement](https://pkg.go.dev/github.com/nadhirfr/oy-go/scheduleddisbursement)
- [Checkout](https://pkg.go.dev/github.com/nadhirfr/oy-go/checkout)
- [Account Inquiry](https://pkg.go.dev/github.com/nadhirfr/oy-go/accountinquiry)

## Contribute

For any requests, bugs, or comments, please [open an issue](https://github.com/nadhirfr/oy-go/issues/new) or [submit a pull request](https://github.com/nadhirfr/oy-go/pulls).

## License
MIT © [LICENSE](https://github.com/nadhirfr/oy-go/blob/main/LICENSE)
