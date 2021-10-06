package lorelai

import "strings"

func formatWords(textToFormat string) string {
	return strings.TrimSpace(strings.ToUpper(textToFormat[0:1])+textToFormat[1:]) + "."
}
