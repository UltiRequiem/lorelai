# Lorelai ![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/lorelai?color=blue&label=Total%20Lines) [![PkgGoDev](https://pkg.go.dev/badge/github.com/UltiRequiem/lorelai)](https://pkg.go.dev/github.com/UltiRequiem/lorelai) ![CodeQL](https://github.com/UltiRequiem/lorelai/workflows/CodeQL/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/lorelai)](https://goreportcard.com/report/github.com/UltiRequiem/lorelai)

Lorelai its a [package](#documentation) and [command line tool](#cli-tool) focused on generate [Lorem ipsum](https://en.wikipedia.org/wiki/Lorem_ipsum).

## Install

```bash
go get github.com/UltiRequiem/lorelai/pkg
```

## Examples

### Convenience Utilities

```go
package main

import (
	"fmt"
	"github.com/UltiRequiem/lorelai/pkg"
)

func main() {
	fmt.Println(fmt.Sprintf("My, my name is: %s.", lorelai.Word()))
	fmt.Println(fmt.Sprintf("My email address is %s", lorelai.Email()))
	fmt.Println(fmt.Sprintf("My website is: %s", lorelai.URL()))
}
```

For more examples check the [examples directory](./example/main.go).

### Documentation

This package exports 8 functions:

- [Word](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L29): Returns a Single word.

> Example: "sodales", "phasellus" , "diam", etc.

- [Sentence](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L34)

- [Paragraph](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L39)

- [FormattedLoremWords](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L24)

- [LoremWords](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L10)

- [Domain](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L10)

- [URL](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L17)

- [Email](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L22)

## CLI Tool

### Installation

```bash
go install github.com/UltiRequiem/lorelai@latest
```

Or use a binary from [releases](https://github.com/UltiRequiem/lorelai/releases/latest).

### License

This project is licensed under the [MIT License](./LICENSE.md).
