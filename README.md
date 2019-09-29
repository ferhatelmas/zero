## zero

[![Build Status](https://travis-ci.org/ferhatelmas/zero.svg?branch=master)](https://travis-ci.org/ferhatelmas/zero)
[![GoDoc](https://godoc.org/github.com/ferhatelmas/zero?status.svg)](https://godoc.org/github.com/ferhatelmas/zero)

> detect and replace zero width characters in golang

### install

> go get github.com/ferhatelmas/zero

### usage

```go

import "github.com/ferhatelmas/zero"

s := "Lorem​Ipsum​Dolor"
zero.Has(s)
// true

zero.Replace(s, "|")
// Lorem|Ipsum|Dolor

zero.Replace(s, "")
// LoremIpsumDolor
```
