# Finfini SDK

[![GoDoc](https://pkg.go.dev/badge/github.com/finfini/go-sdk?status.svg)](https://pkg.go.dev/github.com/finfini/go-sdk?tab=doc)
[![Release](https://img.shields.io/github/release/finfini/go-sdk.svg?style=flat-square)](https://github.com/finfini/go-sdk/releases)

an SDK for integrating to Finfini' services.

# Examples:

## Instalation

```
go get github.com/finfini/go-sdk
```

## Check String Similarity

```go
package main

import (
    "github.com/finfini/go-sdk/ekyc"
    "fmt"
)

func main() {
    first := "abcde"
    second := "defgh"
    threshold := 80.0

    result := ekyc.StringSimilarity(first, second, threshold)

    fmt.Printf("Comparing \"%s\" with \"%s\"; Ratio to reach: %.02f\n", first, second, threshold)
    fmt.Printf("\tRatio: %f\n", result.Ratio)
    fmt.Printf("\tPassed: %v\n", result.Passed)
    fmt.Println("====================================\n")
}
```

or please see [this](/_example/ekyc/similarity.go) file.

# Copyright

2020 - Finfini.com

# Credits

- Patrick Mézard - [source](https://github.com/pmezard/go-difflib)
