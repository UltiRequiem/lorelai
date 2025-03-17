package cmd

import (
	"fmt"
	"os"

	chigo "github.com/UltiRequiem/chigo/pkg"
	lorelai "github.com/UltiRequiem/lorelai/pkg"
)

const VERSION = "1.1.1"

func Main() {
	help, words, paragraphs, sentences, fileToWrite, url, email, colors := flags()

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

	if fileToWrite != "" {
		err := os.WriteFile(fileToWrite, []byte(text), 0664)

		if err != nil {
			error(fmt.Sprintf("Error while trying to write %s.", fileToWrite))
		}

		return
	}

	if text != "" {
		textNoNewline := text[0 : len(text)-2] // Remove the new line at the end

		if colors {
			chigo.PrintWithColors(textNoNewline)
			return
		}

		fmt.Println(textNoNewline)
		return
	}

	printHelp()
}
