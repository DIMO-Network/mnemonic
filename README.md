# Mnemonic

![GitHub license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![GoDoc](https://godoc.org/github.com/KevinJoiner/mnemonic?status.svg)](https://godoc.org/github.com/KevinJoiner/mnemonic)
[![Go Report Card](https://goreportcard.com/badge/github.com/KevinJoiner/mnemonic)](https://goreportcard.com/report/github.com/KevinJoiner/mnemonic)
## Overview

The Mnemonic package provides a flexible and customizable way to encode data into a mnemonic word list. It implements the [BIP-0039](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) specification using [`big.Int`](https://pkg.go.dev/math/big), allowing for the use of arbitrary entropy sizes.

## Features

- **BIP-0039 Specification:** The implementation adheres to the BIP-0039 specification, providing compatibility with various cryptographic systems.

- **Arbitrary Entropy Sizes:** This package allows users to work with arbitrary entropy sizes, providing flexibility in encoding different types of data.

- **Number Obfuscation:** For enhanced usability with IDs, the package offers optional number obfuscation. This feature uses Modular Multiplicative Inverse to convert the provided number into a seemingly random number before generating the mnemonic word list using https://github.com/KevinJoiner/mnemonic

## Getting Started

### Installation

To use this package in your Go project, run the following command:

```bash
go get github.com/KevinJoiner/mnemonic
```

### Example Usage
All Examples can be found in the [go docs](https://godoc.org/github.com/KevinJoiner/mnemonic) or [examples_test.go](./examples_test.go)

```go
package main

import (
	"fmt"
	"github.com/KevinJoiner/mnemonic"
)

func main() {
	// Example usage with a number
	number := "12345"
	mnemonic := mnemonic.FromInt(number)
	fmt.Println("Mnemonic Word List:", mnemonic)

	// Example usage with a word list
	words := []string{"apple", "banana", "orange"}
	decodedNumber := mnemonic.ToInt(words)
	fmt.Println("Int from Word List:", decodedNumber)
}
```

