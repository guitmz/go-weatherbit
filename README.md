# go-weatherbit
A Go package for interacting with the Weatherbit API.

# Installation
`$ go get -u "github.com/guitmz/go-weatherbit"`

# Usage
```go
package main

import (
	"fmt"

	"github.com/guitmz/go-weatherbit"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	weather, err := weatherbit.GetWeather("COUNTRY", "CITY", "API_KEY")
	check(err)
  fmt.Printf("%+v\n", weather)
}
```
