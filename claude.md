# Lorelai Project Overview

## What is Lorelai?

Lorelai is a Go package and CLI tool for generating Lorem Ipsum placeholder text. It provides both random word generation and classic Lorem Ipsum text generation with a clean, simple API.

## Project Structure

```
lorelai/
├── cmd/                    # CLI application
│   └── root.go            # Command-line interface implementation
├── pkg/                    # Core library package
│   ├── root.go            # Random Lorem generation functions
│   ├── classic.go         # Classic Lorem Ipsum functions
│   ├── convenience.go     # Utility functions (URL, Email, Domain)
│   ├── data.go            # Word corpus
│   ├── tlds.go            # Top-level domain data
│   ├── utils.go           # Internal helper functions
│   └── *_test.go          # Test files
├── example/               # Example usage code
└── readme.md              # Documentation
```

## Core Architecture

### Two Modes of Operation

**1. Random Mode** - Generates random Lorem Ipsum-like text:

- Uses a corpus of Latin words (`DATA` in `data.go`)
- Randomly selects words for generation
- Functions: `Word()`, `Sentence()`, `Paragraph()`, `Generate()`

**2. Classic Mode** - Uses traditional Lorem Ipsum text:

- Based on Cicero's "De finibus bonorum et malorum"
- Always starts with "Lorem ipsum dolor sit amet..."
- Functions: `ClassicWords()`, `ClassicSentence()`, `ClassicParagraph()`, `ClassicGenerate()`

### Key Design Patterns

#### 1. Consistent Naming Convention

All functions follow a clear pattern:

- Base functions for random mode: `Word()`, `Sentence()`, `Paragraph()`
- Classic variants prefixed: `ClassicWord()`, `ClassicSentence()`, `ClassicParagraph()`

#### 2. Composition Over Complexity

- Simple, focused functions that do one thing well
- No stateful services or complex object hierarchies
- Functions compose easily: `Sentence()` uses `LoremWords(8)`

#### 3. Performance Optimizations

- Uses `strings.Builder` with capacity hints (`b.Grow()`)
- Avoids unnecessary string concatenations
- Thread-safe random generation with `math/rand/v2`

## API Design Philosophy

### Simplicity First

```go
// Simple, direct API
text := lorelai.Paragraph()

// Structured generation with metadata
result := lorelai.Generate(3, 5)
fmt.Printf("Generated %d words\n", result.WordCount)
```

### Sensible Defaults

- Sentences are always 8 words
- Paragraphs are always 45 words
- All text is properly capitalized and punctuated

### No Configuration Required

- Zero-dependency operation (except stdlib)
- No initialization or setup needed
- Import and use immediately

## How Text Generation Works

### Random Mode Process

1. Select random words from `DATA` corpus
2. Join words with spaces
3. Capitalize first word
4. Add period at end
5. Return formatted string

### Classic Mode Process

1. Use sequential words from `classicText` array
2. Start from beginning: "lorem", "ipsum", "dolor"...
3. Wrap around if more words needed than available
4. Format with capitalization and punctuation

### Structured Generation

Both `Generate()` and `ClassicGenerate()`:

1. Calculate total capacity needed
2. Pre-allocate `strings.Builder` buffer
3. Loop through paragraphs and sentences
4. Insert separators (`\n\n` for paragraphs, space for sentences)
5. Return `Lorem` struct with text and metadata

## Testing Strategy

### Test Coverage: 97.1%

**Unit Tests** cover:

- Individual function outputs (word counts, formatting)
- Edge cases (0, negative numbers)
- Formatting correctness (capitalization, punctuation)
- Metadata accuracy

**Concurrent Tests** verify:

- Thread-safe random generation
- No race conditions in word selection

**Integration Tests** ensure:

- CLI flags work correctly
- File output functions properly
- Colors render as expected

## CLI vs Package Semantics

### Important Distinction

**CLI** (additive behavior):

```bash
lorelai -p 3 -s 5    # Generates 3 paragraphs + 5 sentences
```

**Package** (multiplicative behavior):

```go
lorelai.Generate(3, 5)    // Generates 3 paragraphs OF 5 sentences each
```

This difference is intentional:

- CLI optimized for quick, ad-hoc text generation
- Package optimized for structured, programmatic use

## Convenience Features

Beyond Lorem Ipsum text, lorelai provides:

### Domain Names

```go
lorelai.Domain()    // "neque.net"
```

### URLs

```go
lorelai.URL()       // "https://pellentesque.org"
```

### Email Addresses

```go
lorelai.Email()     // "bibendum@id.pe"
```

These use the same word corpus plus TLD (Top-Level Domain) data.

## Code Quality Practices

### 1. Clear Function Documentation

Every exported function has:

- Clear description of what it does
- Input/output specifications
- Usage examples where helpful

### 2. Efficient String Building

Always use `strings.Builder` with capacity hints:

```go
var b strings.Builder
b.Grow(expectedSize)  // Pre-allocate
```

### 3. DRY Principle

Common operations extracted to helpers:

- `formatWords()` - Capitalize and add punctuation
- `capitalizeFirstWord()` - Title case first character
- `trimSpaceAddDot()` - Clean up and punctuate

### 4. Defensive Programming

- Handle zero/negative inputs gracefully
- Validate array bounds before access
- Return empty strings for invalid input (not errors)

## Performance Characteristics

### Memory Efficiency

- Pre-allocated buffers reduce reallocations
- `strings.Builder` avoids intermediate string copies
- Fixed-size word corpus (`DATA` array)

### Time Complexity

- Random word selection: O(1)
- Text generation: O(n) where n is output length
- No complex algorithms or data structures

### Concurrency

- Thread-safe due to Go 1.22+ `math/rand/v2`
- No shared mutable state
- Functions are pure (aside from randomness)

## Common Use Cases

### 1. Web Development

Generate placeholder text for UI mockups:

```go
description := lorelai.Paragraph()
title := lorelai.Sentence()
```

### 2. Testing

Create realistic test data:

```go
for i := 0; i < 100; i++ {
    users[i].Bio = lorelai.Generate(2, 3).Text
}
```

### 3. Documentation

Fill example content in docs:

```go
example := lorelai.ClassicGenerate(1, 2)
```

### 4. CLI Quick Tasks

```bash
lorelai -p 5 | pbcopy  # Copy to clipboard
lorelai -w 100 > file.txt  # Save to file
```

## Extending Lorelai

### Adding New Functions

Follow the established patterns:

1. **Name consistently**: `Thing()` for random, `ClassicThing()` for classic
2. **Use helpers**: Reuse `formatWords()`, `capitalizeFirstWord()`
3. **Add tests**: Include edge cases and format verification
4. **Document**: Clear godoc comments with examples

### Adding New Features

Consider:

- Is it stateless? (Should be)
- Does it fit the "simple text generation" domain?
- Is the API obvious without reading docs?
- Does it maintain backward compatibility?

## Dependencies

### Production

- **Zero external dependencies**
- Uses only Go standard library
- Minimum Go version: 1.22 (for `math/rand/v2`)

### Development/CLI

- `github.com/spf13/cobra` - CLI framework
- `github.com/UltiRequiem/chigo/pkg` - Color output

## Project Philosophy

### Simplicity

> "Do one thing well" - Generate placeholder text, nothing more

### Pragmatism

> Sensible defaults over configuration. 8 words? Perfect. No need to change it.

### Accessibility

> Zero learning curve. Import, call a function, get text. Done.

### Performance

> Fast enough to not think about it. Pre-allocate, avoid copies, use `strings.Builder`.

## Maintenance Guidelines

### When Adding Code

- ✅ Maintain 95%+ test coverage
- ✅ Keep functions small and focused
- ✅ Follow existing naming patterns
- ✅ Update documentation and examples

### When Reviewing PRs

- ❓ Does it add real value?
- ❓ Is the API intuitive?
- ❓ Are there tests?
- ❓ Is it backward compatible?

### When Fixing Bugs

- 🐛 Add a test that reproduces the bug
- 🐛 Fix the issue
- 🐛 Verify the test passes
- 🐛 Consider similar issues elsewhere

## Historical Context

Lorelai was created to provide a simple, dependency-free Lorem Ipsum generator for Go projects. Unlike other libraries that require configuration or complex APIs, Lorelai focuses on:

1. **Immediate usability** - No setup, just import and use
2. **Predictable output** - Consistent word counts and formatting
3. **Both modes** - Random for variety, classic for tradition
4. **CLI included** - Quick shell usage without writing code

The project has grown to include convenience utilities (domains, URLs, emails) while maintaining its core simplicity.

---

_For implementation details and changes, see the Git history and PR discussions._
