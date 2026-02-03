# Lorelai

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bobadilla-tech/lorelai)](https://pkg.go.dev/github.com/bobadilla-tech/lorelai)
![CodeQL](https://github.com/bobadilla-tech/lorelai/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobadilla-tech/lorelai)](https://goreportcard.com/report/github.com/bobadilla-tech/lorelai)

A [package](#documentation) and [command line tool](#cli-tool) to generate
[Lorem ipsum](https://en.wikipedia.org/wiki/Lorem_ipsum).

![Cover](./assets/cover.gif)

[Blog explaining this package](https://ultirequiem.hashnode.dev/lorelai)

Checkout the code coverage at: https://bobadilla-tech.github.io/lorelai/#file0

## Install

```bash
go get github.com/bobadilla-tech/lorelai/pkg
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

### Structured Text Generation

```go
func generateStructuredContent() {
	// Generate 3 paragraphs with 5 sentences each
	result := lorelai.Generate(3, 5)
	fmt.Println(result.Text)
	fmt.Printf("Generated %d words across %d paragraphs\n",
		result.WordCount, result.Paragraphs)
}

func generateClassicLorem() {
	// Generate classic Lorem Ipsum text
	result := lorelai.ClassicGenerate(2, 4)
	fmt.Println(result.Text)
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

### Core Text Generation Functions

**Random Lorem Ipsum:**

- **[Word](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Returns 1 random word
  E.g: "sodales", "phasellus", "diam"

- **[Sentence](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Returns 8 random words (1 sentence)
  E.g: "Varius sed imperdiet amet laoreet ex sapien placerat."

- **[Paragraph](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Returns 45 random words (1 paragraph)
  E.g: "Nisi lacinia ante non nunc eros nibh mattis enim orci ante in ornare accumsan..."

- **[Generate](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Generate structured text with X paragraphs of Y sentences each
  Returns: `Lorem{Text, Paragraphs, Sentences, WordCount}`
  E.g: `Generate(3, 5)` creates 3 paragraphs with 5 sentences each (120 words)

**Classic Lorem Ipsum:**

- **[ClassicSentence](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/classic.go)**: Returns 8 words from classic text
  Always starts with "Lorem ipsum dolor sit amet..."

- **[ClassicParagraph](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/classic.go)**: Returns full classic Lorem paragraph
  The traditional Lorem Ipsum text

- **[ClassicGenerate](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/classic.go)**: Generate classic text with X paragraphs of Y sentences each
  Returns: `Lorem{Text, Paragraphs, Sentences, WordCount}`
  E.g: `ClassicGenerate(2, 4)` creates 2 paragraphs with 4 classic sentences each

- **[ClassicWords](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/classic.go)**: Returns N words from classic text

### Word Generation Functions

- **[LoremWords](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Returns N random words
  E.g: "arcu", "blandit porttitor a scelerisque", "donec justo lacinia"

- **[FormattedLoremWords](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/root.go)**: Returns N words with capitalization and period
  E.g: "Libero malesuada duis massa luctus.", "Curabitur hendrerit sed."

### Convenience Utilities

- **[Domain](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/convenience.go)**: Returns a random domain
  E.g: "neque.net", "arcu.org", "lorem.io"

- **[URL](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/convenience.go)**: Returns a random URL
  E.g: "https://pellentesque.org", "https://id.io", "https://efficitur.com"

- **[Email](https://github.com/bobadilla-tech/lorelai/blob/main/pkg/convenience.go)**: Returns a random email address
  E.g: "bibendum@id.pe", "ornare@duis.pe", "quisque@faucibus.org"

## CLI Tool

### Installation

```bash
go install github.com/bobadilla-tech/lorelai@latest
```

Or use a binary from
[releases](https://github.com/bobadilla-tech/lorelai/releases/latest).

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
 https://github.com/bobadilla-tech/lorelai
```

## Contributing

Feel free to suggest new features or report bugs.

A big thanks to [Antoineio](https://github.com/Antoineio) for adding tests and
contributing to the CI.

## License

This project is licensed under the [MIT License](./license).
