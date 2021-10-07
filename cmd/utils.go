package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func error(text string) {
	color.Red(text)
	os.Exit(1)
}

func repeatedFlag(longName string, shortName string) {
	error(fmt.Sprintf("You cannot pass --%s and -%s at the same time!", longName, shortName))
}

func printHelp() {
	fmt.Println("TODO")
}
