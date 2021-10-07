package cmd

import "flag"

func flags() (bool, int, int, int, string, bool, bool) {
	help := flag.Bool("help", false, "Display Help")
	helpShort := flag.Bool("h", false, "Display Help")

	words := flag.Int("words", 0, "Number of words")
	wordsShort := flag.Int("w", 0, "Number of words")

	paragraph := flag.Int("paragraph", 0, "Number of paragraphs.")
	paragraphShort := flag.Int("p", 0, "Number of paragraphs.")

	sentences := flag.Int("sentences", 0, "Number of sentences.")
	sentencesShort := flag.Int("s", 0, "Number of sentences.")

	output := flag.String("output", "", "Number of sentences.")
	outputShort := flag.String("o", "", "Number of sentences.")

	url := flag.Bool("url", false, "Print")
	email := flag.Bool("emal", false, "Print")

	flag.Parse()

	if *help && *helpShort {
		repeatedFlag("help", "h")
	}

	if *words != 0 && *wordsShort != 0 {
		repeatedFlag("words", "w")
	}

	if *paragraph != 0 && *paragraphShort != 0 {
		repeatedFlag("paragraph", "p")
	}

	if *sentences != 0 && *sentencesShort != 0 {
		repeatedFlag("sentences", "s")
	}

	if *output != "" && *outputShort != "" {
		repeatedFlag("output", "o")
	}

	if *words == 0 {
		words = wordsShort
	}

	if *paragraph == 0 {
		paragraph = paragraphShort
	}

	if *sentences == 0 {
		sentences = sentencesShort
	}

	if *output == "" {
		output = outputShort
	}

	return *help || *helpShort, *words, *paragraph, *sentences, *output, *url, *email
}
