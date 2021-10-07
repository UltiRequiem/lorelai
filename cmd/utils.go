package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func repeatedFlag(longName string, shortName string) {
	color.Red(fmt.Sprintf("You cannot pass --%s and -%s at the same time!", longName, shortName))
	os.Exit(1)
}
