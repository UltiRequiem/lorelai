package cmd

import (
	"fmt"
	"os"
	"strings"

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

	var text strings.Builder

	for i := 0; i < words; i++ {
		text.WriteString(lorelai.Word())
		text.WriteByte('\n')
	}

	for i := 0; i < sentences; i++ {
		text.WriteString(lorelai.Sentence())
		text.WriteString("\n\n")
	}

	for i := 0; i < paragraphs; i++ {
		text.WriteString(lorelai.Paragraph())
		text.WriteString("\n\n")
	}

	if fileToWrite != "" {
		err := os.WriteFile(fileToWrite, []byte(text.String()), 0664)

		if err != nil {
			error(fmt.Sprintf("Error while trying to write %s.", fileToWrite))
		}

		return
	}

	if text.Len() > 0 {
		textStr := text.String()
		textNoNewline := textStr[0 : len(textStr)-2] // Remove the new line at the end

		if colors {
			chigo.PrintWithColors(textNoNewline)
			return
		}

		fmt.Println(textNoNewline)
		return
	}

	printHelp()
}
