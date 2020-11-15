# timekeeper

[![Go Report Card](https://goreportcard.com/badge/syeo66/timekeeper)](https://goreportcard.com/report/syeo66/timekeeper)

A small helper for tracking time spent in a function in Go

## Installation

```bash
go get github.com/syeo66/timekeeper
```

## Usage

```go
package main

import "github.com/syeo66/timekeeper"

func main() {
    timekeeper.StartTime()
    defer timekeeper.EndTime()

    // do something else
}
```

## Documentation

See [godoc](https://pkg.go.dev/github.com/syeo66/timekeeper#section-documentation) documentation 
