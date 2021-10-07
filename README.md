# This is still a wip!

# Lorelai

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/lorelai?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/lorelai/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/lorelai)](https://goreportcard.com/report/github.com/UltiRequiem/lorelai)

## Install

```bash
go get github.com/UltiRequiem/lorelai/pkg
```

## Usage

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
