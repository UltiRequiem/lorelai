// Provides Random Lorem Ipsum, and other utilities.
package lorelai

import (
	"math/rand/v2"
	"strings"
)

// Lorem represents generated Lorem Ipsum text with metadata
type Lorem struct {
	Text       string
	Paragraphs int
	Sentences  int
	WordCount  int
}

// Get [quantity] words
func LoremWords(quantity int) string {
	if quantity <= 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(quantity * 8) // heuristic: avg word + space

	for i := 0; i < quantity; i++ {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(DATA[rand.IntN(len(DATA))])
	}

	return b.String()
}

// Get [quantity] words, the first word is capitalized, ends with a dot.
func FormattedLoremWords(quantity int) string {
	return formatWords(LoremWords(quantity))
}

// Get a single word
func Word() string {
	word := LoremWords(1)

	return strings.TrimSpace(strings.ToTitle(word))
}

func Title() string {
	return strings.ToTitle(LoremWords(1))
}

// Get a sentence, with the first word capitalized, of eight words that ends with a dot.
func Sentence() string {
	return formatWords(LoremWords(8))
}

// Get a paragraph, with the first word capitalized, of forty five words that ends with a dot.
func Paragraph() string {
	return formatWords(LoremWords(45))
}

// Generate creates Lorem Ipsum text with specified paragraphs and sentences.
// Each paragraph contains exactly 'sentences' sentences.
// Each sentence contains 8 random words.
//
// Example: Generate(3, 5) creates 3 paragraphs, each with 5 sentences.
//
// This differs from the CLI behavior where -p and -s are additive.
func Generate(paragraphs int, sentences int) Lorem {
	if paragraphs <= 0 || sentences <= 0 {
		return Lorem{
			Text:       "",
			Paragraphs: paragraphs,
			Sentences:  sentences,
			WordCount:  0,
		}
	}

	var b strings.Builder

	// Heuristic to reduce reallocations; exact sizing not required
	b.Grow(paragraphs * sentences * 64)

	for p := range paragraphs {
		if p > 0 {
			b.WriteString("\n\n")
		}
		for s := range sentences {
			if s > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(Sentence())
		}
	}

	text := b.String()

	return Lorem{
		Text:       text,
		Paragraphs: paragraphs,
		Sentences:  sentences,
		WordCount:  paragraphs * sentences * 8,
	}
}
