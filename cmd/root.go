package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	chigo "github.com/UltiRequiem/chigo/pkg"
	lorelai "github.com/UltiRequiem/lorelai/pkg"
)

const VERSION = "1.1.1"

var (
	words      int
	paragraphs int
	sentences  int
	output     string
	url        bool
	email      bool
	colors     bool
	classic    bool
)

var rootCmd = &cobra.Command{
	Use:     "lorelai",
	Short:   "Easily generate Lorem Ipsum on command line",
	Version: VERSION,
	Long: `lorelai ` + VERSION + `
Easily generate Lorem Ipsum on command line.

Examples:
    lorelai -w 55           # Will print 55 words
    lorelai -p 5            # Will print 5 paragraphs
    lorelai -s 5 --output b # Will write 5 sentences on file b if possible
    lorelai -w 55 -s 5      # Will print 55 words and 5 sentences

If you need more help, found a bug or want to suggest a new feature:
https://github.com/UltiRequiem/lorelai`,
	RunE: run,
}

func init() {
	rootCmd.Flags().IntVarP(&words, "words", "w", 0, "Number of words to print")
	rootCmd.Flags().IntVarP(&paragraphs, "paragraph", "p", 0, "Number of paragraphs to print")
	rootCmd.Flags().IntVarP(&sentences, "sentences", "s", 0, "Number of sentences to print")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "If passed it will try to put the output in a file")
	rootCmd.Flags().BoolVar(&url, "url", false, "A random URL")
	rootCmd.Flags().BoolVar(&email, "email", false, "A random Email Address")
	rootCmd.Flags().BoolVar(&colors, "color", false, "Print the output with colors?")
	rootCmd.Flags().BoolVar(&classic, "classic", false, "Use classic Lorem Ipsum text (starts with 'Lorem ipsum dolor sit amet...')")
}

func run(cmd *cobra.Command, args []string) error {
	if url {
		fmt.Println(lorelai.URL())
		return nil
	}

	if email {
		fmt.Println(lorelai.Email())
		return nil
	}

	var text strings.Builder

	if classic {
		// Classic Lorem Ipsum mode
		if words > 0 {
			text.WriteString(lorelai.ClassicWords(words))
			text.WriteString("\n")
		}

		for i := 0; i < sentences; i++ {
			text.WriteString(lorelai.ClassicSentence())
			text.WriteString("\n\n")
		}

		for i := 0; i < paragraphs; i++ {
			text.WriteString(lorelai.ClassicParagraph())
			text.WriteString("\n\n")
		}
	} else {
		// Random Lorem Ipsum mode
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
	}

	if output != "" {
		err := os.WriteFile(output, []byte(text.String()), 0664)
		if err != nil {
			return fmt.Errorf("error while trying to write %s: %w", output, err)
		}
		return nil
	}

	if text.Len() > 0 {
		textStr := text.String()
		textNoNewline := textStr[0 : len(textStr)-2] // Remove the new line at the end

		if colors {
			chigo.PrintWithColors(textNoNewline)
			return nil
		}

		fmt.Println(textNoNewline)
		return nil
	}

	// If no flags provided, show help
	return cmd.Help()
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
