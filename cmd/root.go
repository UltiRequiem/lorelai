package cmd

import (
	"fmt"
	"os"

	"github.com/UltiRequiem/lorelai/pkg"
	"github.com/fatih/color"
)

func Main() {
	help, words, paragraphs, sentences, output, url, email := flags()

	if url {
		fmt.Println(lorelai.URL())
                return
	}

	if email {
		fmt.Println(lorelai.Email())
                return
	}

	if help {
		printHelp()
		return
	}

	text := ""

	for i := 0; i < words; i++ {
		text += lorelai.Word()
	}

	for i := 0; i < paragraphs; i++ {
		text += lorelai.Paragraph()
	}

	for i := 0; i < sentences; i++ {
		text += lorelai.Sentence()
	}

	if output != "" {
		err := os.WriteFile(output, []byte(text), 0664)
		if err != nil {
			error(fmt.Sprintf("Error while trying to write %s.", output))
		}

		color.Blue("Done!")

		return
	}

	fmt.Println(text)
}
