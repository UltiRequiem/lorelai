// Provides Random Lorem Ipsum, and other utilities.
package lorelai

import (
	"math/rand"
	"strings"
	"time"
)

// Get [quantity] words
func LoremWords(quantity int) string {
	rand.Seed(time.Now().UnixNano())

	lorem := ""

	for i := 0; i < quantity; i++ {
		lorem += DATA[rand.Intn(1100)] + " "
	}

	return lorem

}

// Get [quantity] words, the first word is capitalized, ends with a dot.
func FormattedLoremWords(quantity int) string {
	return formatWords(LoremWords(quantity))
}

// Get a single word
func Word() string {
	word := LoremWords(1)
	return strings.TrimSpace(strings.ToUpper(word[0:1]) + word[1:])
}

// Get a sentence, with the first word capitalized, of eight words that ends with a dot.
func Sentence() string {
	return formatWords(LoremWords(8))
}

// Get a paragraph, with the first word capitalized, of forty five words that ends with a dot.
func Paragraph() string {
	return formatWords(LoremWords(45))
}
