package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/UltiRequiem/chigo/pkg"
)

func error(text string) {
	printWithScanner(text)
	os.Exit(1)
}

func repeatedFlag(longName string, shortName string) {
	error(fmt.Sprintf("You cannot pass --%s and -%s at the same time!", longName, shortName))
}

func printWithScanner(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	var j int = 1

	for scanner.Scan() {
		r, g, b := chigo.GetRGB(j)

		fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, scanner.Text())

		j++
	}
}

func printHelp() {
	printWithScanner(fmt.Sprintf(` lorelai %s
 Easily generate Lorem Ipsum on command line.

 -h or --help           Print this

 -w or --word           Number of words to print

 -p or --paragraph      Number of paragraphs to print

 -s or --sentences      Number of sentences to print

 -o or --output         If passed it will try to put the output in a file

 --url                  A random URL

 --email                A random Email Address

 --color                Print the output with colors?

  Examples:
      lorelai -w 55           # Will print 55 words
      lorelai -p 5            # Will print 5 paragraphs
      lorelai -s 5 --output b # Will write 5 sentences on file b if possible
      lorelai -w 55 -s 5      # Will printe 55 words and 5 sentences

 If you need more help, found a bug or want to suggest a new feature:
 https://github.com/UltiRequiem/lorelai`, VERSION))
}
