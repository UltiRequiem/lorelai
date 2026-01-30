# Lorelai

[![PkgGoDev](https://pkg.go.dev/badge/github.com/UltiRequiem/lorelai)](https://pkg.go.dev/github.com/UltiRequiem/lorelai)
![CodeQL](https://github.com/UltiRequiem/lorelai/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/lorelai)](https://goreportcard.com/report/github.com/UltiRequiem/lorelai)

A [package](#documentation) and [command line tool](#cli-tool) to generate
[Lorem ipsum](https://en.wikipedia.org/wiki/Lorem_ipsum).

![Cover](./assets/cover.gif)

[Blog explaining this package](https://ultirequiem.hashnode.dev/lorelai)

Checkout the code coverage at: https://ulti.js.org/lorelai

## Install

```bash
go get github.com/UltiRequiem/lorelai/pkg
```

## Examples

### Generating Text

```go
func printTonsOfText() {
	for i := 0; i < 10; i++ {
          fmt.Println(lorelai.Paragraph())
	}
}
```

### Convenience Utilities

```go
func sayHi() {
	fmt.Printf("My, my name is: %s.\n", lorelai.Word())
	fmt.Printf("My email address is %s\n", lorelai.Email())
	fmt.Printf("My website is: %s\n", lorelai.URL())
}
```

For more examples check the [examples directory](./example/main.go).

## Documentation

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

E.g: "Libero malesuada duis massa luctus.", "Curabitur hendrerit sed.",
"Ligula.", etc.

- [LoremWords](https://github.com/UltiRequiem/lorelai/blob/main/pkg/root.go#L10):
  It receives a number and returns a string with the number of words you have
  indicated.

E.g: "arcu", "blandit porttitor a scelerisque", "donec justo lacinia", etc.

- [Domain](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L10):
  Returns a domain

E.g: "neque.net", "arcu.org" , "lorem.io", etc.

- [URL](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L17):
  Returns an URL

E.g: "https://pellentesque.org", "https://id.io" , "https://efficitur.com", etc.

- [Email](https://github.com/UltiRequiem/lorelai/blob/main/pkg/convenience.go#L22):
  Returns an email address

E.g: "bibendum@id.pe", "ornare@duis.pe" , "quisque@faucibus.org", etc.

## CLI Tool

### Installation

```bash
go install github.com/UltiRequiem/lorelai@latest
```

Or use a binary from
[releases](https://github.com/UltiRequiem/lorelai/releases/latest).

### Usage

If you don't pass any flag or you pass the help flag:

```
 lorelai 1.0.0
 Easily generate Lorem Ipsum on command line.

 -h or --help           Print this

 -w or --word           Number of words to print

 -p or --paragraph      Number of paragraphs to print

 -s or --sentences      Number of sentences to print

 -o or --output         If passed it will try to put the output in a file

 --url                  A random URL

 --email                A random Email Address

 --color                Print the output with colors?

  Examples:
      lorelai -w 55           # Will print 55 words
      lorelai -p 5            # Will print 5 paragraphs
      lorelai -s 5 --output b # Will write 5 sentences on file b if possible
      lorelai -w 55 -s 5      # Will print 55 words and 5 sentences

 If you need more help, found a bug or want to suggest a new feature:
 https://github.com/UltiRequiem/lorelai
```

## Contributing

Feel free to suggest new features or report bugs.

A big thanks to [Antoineio](https://github.com/Antoineio) for adding tests and
contributing to the CI.

## License

This project is licensed under the [MIT License](./license).
