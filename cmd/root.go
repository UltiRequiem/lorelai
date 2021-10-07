package cmd

import (
	"fmt"
	"os"

	"github.com/UltiRequiem/lorelai/pkg"
	"github.com/fatih/color"
)

func Main() {
	help, words, paragraphs, sentences, output, url, email := flags()

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

		color.Blue("Done!")

		return
	}

	fmt.Println(text[0 : len(text)-1])
}
