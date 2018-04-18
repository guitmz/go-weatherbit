[![Build Status](https://travis-ci.org/guitmz/go-weatherbit.svg?branch=master)](https://travis-ci.org/guitmz/go-weatherbit)
[![Go Report Card](https://goreportcard.com/badge/github.com/guitmz/go-weatherbit)](https://goreportcard.com/report/github.com/guitmz/go-weatherbit) ![Go Docs](https://godoc.org/github.com/guitmz/go-weatherbit?status.svg)

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
