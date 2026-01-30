package main

import (
	"fmt"

	"github.com/UltiRequiem/lorelai/pkg"
)

func sayHi() {
	fmt.Printf("My, my name is: %s.\n", lorelai.Word())
	fmt.Printf("My email address is %s\n", lorelai.Email())
	fmt.Printf("My website is: %s\n", lorelai.URL())

	fmt.Printf(`My favorite phrase is: "%s"`+"\n", lorelai.FormattedLoremWords(5))

	fmt.Println("Let me tell you an history:")

	fmt.Println(lorelai.Paragraph())

	fmt.Println("Didn't you like the story? Let me tell you 4 words abou the author:")

	fmt.Println(lorelai.LoremWords(4))
}

func printTonsOfText() {
	for i := 0; i < 10; i++ {
		fmt.Println(lorelai.Paragraph())
	}
}

func main() {
	printTonsOfText()
}
