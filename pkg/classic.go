package lorelai

import "strings"

// Classic Lorem Ipsum text from Cicero's "De finibus bonorum et malorum"
// This is the standard Lorem Ipsum that starts with "Lorem ipsum dolor sit amet"
var classicText = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
	"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
	"magna", "aliqua", "ut", "enim", "ad", "minim", "veniam", "quis", "nostrud",
	"exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea",
	"commodo", "consequat", "duis", "aute", "irure", "dolor", "in", "reprehenderit",
	"in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla",
	"pariatur", "excepteur", "sint", "occaecat", "cupidatat", "non", "proident",
	"sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id",
	"est", "laborum",
}

// ClassicParagraph returns the classic Lorem Ipsum paragraph that starts with
// "Lorem ipsum dolor sit amet, consectetur adipiscing elit..."
func ClassicParagraph() string {
	return buildClassicText(len(classicText))
}

// ClassicSentence returns a sentence from the classic Lorem Ipsum text
func ClassicSentence() string {
	return buildClassicText(8)
}

// ClassicWordsPerSentence returns the word count per sentence
func ClassicWordsPerSentence() int {
	return 8
}

// ClassicWordsPerSentence returns the word count per praragraph
func ClassicWordsPerParagraph() int {
	return len(classicText)
}

// ClassicWords returns the specified number of words from the classic Lorem Ipsum text
func ClassicWords(quantity int) string {
	if quantity <= 0 {
		return ""
	}
	if quantity > len(classicText) {
		quantity = len(classicText)
	}
	return buildClassicText(quantity)
}

func buildClassicText(wordCount int) string {
	if wordCount <= 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(wordCount * 8)

	for i := 0; i < wordCount && i < len(classicText); i++ {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(classicText[i])
	}

	result := b.String()
	return capitalizeFirstWord(result) + "."
}
