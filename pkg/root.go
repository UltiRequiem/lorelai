// Provides Random Lorem Ipsum, and other utilities.
package lorelai

import (
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Get [quantity] words
func LoremWords(quantity int) string {
	if quantity <= 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(quantity * 8) // heuristic: avg word + space

	for i := range quantity {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(DATA[rng.Intn(len(DATA))])
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

func Tittle() string {
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
