package main

import (
	"fmt"

	"github.com/bobadilla-tech/lorelai/pkg"
)

func sayHi() {
	fmt.Printf("My, my name is: %s.\n", lorelai.Word())
	fmt.Printf("My email address is %s\n", lorelai.Email())
	fmt.Printf("My website is: %s\n", lorelai.URL())

	fmt.Printf(`My favorite phrase is: "%s"`+"\n", lorelai.FormattedLoremWords(5))

	fmt.Println("Let me tell you a history:")

	fmt.Println(lorelai.Paragraph())

	fmt.Println("Didn't you like the story? Let me tell you 4 words about the author:")

	fmt.Println(lorelai.LoremWords(4))
}

func printTonsOfText() {
	for i := 0; i < 10; i++ {
		fmt.Println(lorelai.Paragraph())
	}
}

func generateStructuredContent() {
	fmt.Println("\n=== Structured Random Lorem ===")
	// Generate 3 paragraphs with 5 sentences each
	result := lorelai.Generate(3, 5)
	fmt.Println(result.Text)
	fmt.Printf("\nGenerated %d words across %d paragraphs (%d sentences per paragraph)\n",
		result.WordCount, result.Paragraphs, result.Sentences)
}

func generateClassicContent() {
	fmt.Println("\n=== Classic Lorem Ipsum ===")
	// Generate classic Lorem Ipsum text: 2 paragraphs with 4 sentences each
	result := lorelai.ClassicGenerate(2, 4)
	fmt.Println(result.Text)
	fmt.Printf("\nGenerated %d words of classic Lorem Ipsum\n", result.WordCount)
}

func main() {
	generateStructuredContent()
	generateClassicContent()
}
