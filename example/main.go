package main

import (
	"fmt"

	"github.com/UltiRequiem/lorelai/pkg"
)

func main() {
	fmt.Println(fmt.Sprintf("My, my name is: %s.", lorelai.Word()))
	fmt.Println(fmt.Sprintf("My email address is %s", lorelai.Email()))
	fmt.Println(fmt.Sprintf("My website is: %s", lorelai.URL()))

	fmt.Println(fmt.Sprintf(`My favorite phrase is: "%s"`, lorelai.FormattedLoremWords(5)))

	fmt.Println("Let me tell you an history:")

	fmt.Println(lorelai.Paragraph())
}
