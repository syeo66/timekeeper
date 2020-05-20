# go-timekeeper
A small helper for tracking time spent in a function

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
