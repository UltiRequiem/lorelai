// Provides Random Lorem Ipsum
package lorelai

import (
	"math/rand"
	"time"
)

var LOREM_WORDS [1100]string = GetLoremArray()

func LoremWords(quantity int) string {
	rand.Seed(time.Now().Unix())

	lorem := ""

	for i := 0; i < quantity; i++ {
		lorem += LOREM_WORDS[rand.Intn(1100)] + " "
	}

	return lorem

}

func FormattedLoremWords(quantity int) string {
	return formatWords(LoremWords(quantity))
}

func Word() string {
	return formatWords(LoremWords(1))
}

func Phrase() string {
	return formatWords(LoremWords(8))
}

func Paragraph() string {
	return formatWords(LoremWords(45))
}
