package cmd

import (
	"flag"
)

func resolveFlag[T comparable](longVal, shortVal *T, zero T, longName, shortName string) T {
	if *longVal != zero && *shortVal != zero {
		repeatedFlag(longName, shortName)
	}
	if *longVal != zero {
		return *longVal
	}
	return *shortVal
}

func flags() (bool, int, int, int, string, bool, bool, bool) {
	help := flag.Bool("help", false, "Display Help")
	helpShort := flag.Bool("h", false, "Display Help")

	words := flag.Int("words", 0, "Number of words")
	wordsShort := flag.Int("w", 0, "Number of words")

	paragraph := flag.Int("paragraph", 0, "Number of paragraphs.")
	paragraphShort := flag.Int("p", 0, "Number of paragraphs.")

	sentences := flag.Int("sentences", 0, "Number of sentences.")
	sentencesShort := flag.Int("s", 0, "Number of sentences.")

	output := flag.String("output", "", "File to write output.")
	outputShort := flag.String("o", "", "File to write output.")

	url := flag.Bool("url", false, "Print an URL")
	email := flag.Bool("email", false, "Print an Email Address")

	colorful := flag.Bool("color", false, "Print with colors?")

	flag.Usage = printHelp

	flag.Parse()

	finalHelp := resolveFlag(help, helpShort, false, "help", "h")
	finalWords := resolveFlag(words, wordsShort, 0, "words", "w")
	finalParagraph := resolveFlag(paragraph, paragraphShort, 0, "paragraph", "p")
	finalSentences := resolveFlag(sentences, sentencesShort, 0, "sentences", "s")
	finalOutput := resolveFlag(output, outputShort, "", "output", "o")

	return finalHelp, finalWords, finalParagraph, finalSentences, finalOutput, *url, *email, *colorful
}
