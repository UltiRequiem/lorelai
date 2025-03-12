package cmd

import (
	"fmt"
	"os"

	chigo "github.com/UltiRequiem/chigo/pkg"
	"github.com/UltiRequiem/lorelai/pkg"
)

const VERSION = "1.0.0"

func Main() {
	help, words, paragraphs, sentences, output, url, email, colors := flags()

	if help {
		printHelp()
		return
	}

	if url {
		fmt.Println(lorelai.URL())
		return
	}

	if email {
		fmt.Println(lorelai.Email())
		return
	}

	text := ""

	for i := 0; i < words; i++ {
		text += lorelai.Word()
		text += "\n"
	}

	for i := 0; i < sentences; i++ {
		text += lorelai.Sentence()
		text += "\n\n"
	}

	for i := 0; i < paragraphs; i++ {
		text += lorelai.Paragraph()
		text += "\n\n"
	}

	if output != "" {
		err := os.WriteFile(output, []byte(text), 0664)
		if err != nil {
			error(fmt.Sprintf("Error while trying to write %s.", output))
		}

		return
	}

	if text != "" {
		textNoNewline := text[0 : len(text)-1]

		if colors {
			chigo.PrintWithColors(textNoNewline)
			return
		}

		fmt.Println(textNoNewline)
		return
	}

	printHelp()
}
