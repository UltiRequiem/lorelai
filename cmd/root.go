package cmd

import (
	"fmt"
	"github.com/UltiRequiem/lorelai/pkg"
)

func Main() {
	help, words, paragraphs, sentences, output := flags()

	fmt.Println(help, words, paragraphs, sentences, output)

	fmt.Println(lorelai.URL())
}
