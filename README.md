# Lorelai ![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/lorelai?color=blue&label=Total%20Lines) [![PkgGoDev](https://pkg.go.dev/badge/github.com/UltiRequiem/lorelai)](https://pkg.go.dev/github.com/UltiRequiem/lorelai) ![CodeQL](https://github.com/UltiRequiem/lorelai/workflows/CodeQL/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/lorelai)](https://goreportcard.com/report/github.com/UltiRequiem/lorelai)

A [package](#documentation) and [command line tool](#cli-tool) to generate
[Lorem ipsum](https://en.wikipedia.org/wiki/Lorem_ipsum).

## Install

```bash
go get github.com/UltiRequiem/lorelai/pkg
```

## Examples

### Convenience Utilities

```go
func sayHi() {
	fmt.Println(fmt.Sprintf("My, my name is: %s.", lorelai.Word()))
	fmt.Println(fmt.Sprintf("My email address is %s", lorelai.Email()))
	fmt.Println(fmt.Sprintf("My website is: %s", lorelai.URL()))
}
```

For more examples check the [examples directory](./example/main.go).

### Documentation

This package exports 8 functions:

- [Word](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L29):
  Returns 1 Word

E.g: "sodales", "phasellus" , "diam", etc.

- [Sentence](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L34):
  Returns 8 Words

E.g: "Varius sed imperdiet amet laoreet ex sapien placerat.", etc.

- [Paragraph](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L39):
  Returns 45 Words

E.g: "Nisi lacinia ante non nunc eros nibh mattis enim orci ante in ornare
accumsan iaculis vel..."

- [FormattedLoremWords](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L24):
  It receives a number and returns a string with the number of words you have
  indicated. The first letter will be capital and the sentence will end with a
  dot.

E.g: "Libero malesuada duis massa luctus.", "Curabitur hendrerit sed.", "Ligula.", etc.

- [LoremWords](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L10): It receives a number and returns a string with the number of words you have indicated.

E.g: "arcu", "blandit porttitor a scelerisque", "donec justo lacinia", etc.

- [Domain](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L10): Returns a domain

E.g: "neque.net", "arcu.org" , "lorem.io", etc.

- [URL](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L17): Returns an URL

E.g: "https://pellentesque.org", "https://id.io" , "https://efficitur.com", etc.

- [Email](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L22): Retuns an email address

E.g: "bibendum@id.pe", "ornare@duis.pe" , "quisque@faucibus.org", etc.

## CLI Tool

### Installation

```bash
go install github.com/UltiRequiem/lorelai@latest
```

Or use a binary from
[releases](https://github.com/UltiRequiem/lorelai/releases/latest).

### License

This project is licensed under the [MIT License](./LICENSE.md).
