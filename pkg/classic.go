package lorelai

import "strings"

// Classic Lorem Ipsum text from Cicero's "De finibus bonorum et malorum"
// This is the standard Lorem Ipsum that starts with "Lorem ipsum dolor sit amet"
var classicText = [...]string{
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

var wordsPerSentence = 8

// ClassicParagraph returns the classic Lorem Ipsum paragraph that starts with
// "Lorem ipsum dolor sit amet, consectetur adipiscing elit..."
func ClassicParagraph() string {
	return buildClassicText(len(classicText))
}

// ClassicSentence returns a sentence from the classic Lorem Ipsum text
func ClassicSentence() string {
	return buildClassicText(wordsPerSentence)
}

// ClassicWordsPerSentence returns the word count per sentence
func ClassicWordsPerSentence() int {
	return wordsPerSentence
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

// ClassicGenerate creates classic Lorem Ipsum text with specified paragraphs and sentences.
// Each paragraph contains exactly 'sentences' sentences.
// Each sentence contains 8 words from the classic Lorem Ipsum text.
//
// Example: ClassicGenerate(3, 5) creates 3 paragraphs, each with 5 sentences.
//
// This differs from the CLI behavior where -p and -s are additive.
func ClassicGenerate(paragraphs int, sentences int) Lorem {
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
			b.WriteString(ClassicSentence())
		}
	}

	text := b.String()

	return Lorem{
		Text:       text,
		Paragraphs: paragraphs,
		Sentences:  sentences,
		WordCount:  paragraphs * sentences * ClassicWordsPerSentence(),
	}
}

func buildClassicText(wordCount int) string {
	if wordCount <= 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(wordCount * wordsPerSentence)

	for i := 0; i < wordCount && i < len(classicText); i++ {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(classicText[i])
	}

	result := b.String()
	return formatWords(result)
}
