package main

import (
	"fmt"

	"github.com/UltiRequiem/lorelai/pkg"
)

func sayHi() {
	fmt.Println(fmt.Sprintf("My, my name is: %s.", lorelai.Word()))
	fmt.Println(fmt.Sprintf("My email address is %s", lorelai.Email()))
	fmt.Println(fmt.Sprintf("My website is: %s", lorelai.URL()))

	fmt.Println(fmt.Sprintf(`My favorite phrase is: "%s"`, lorelai.FormattedLoremWords(5)))

	fmt.Println("Let me tell you an history:")

	fmt.Println(lorelai.Paragraph())

	fmt.Println(fmt.Sprintf("Didn't you like the story? Let me tell you 4 words abou the author:"))

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
