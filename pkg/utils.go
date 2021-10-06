package lorelai

import "strings"

// Trim spaces and adds a dot at the end.
func formatWords(textToFormat string) string {
	return strings.TrimSpace(strings.ToUpper(textToFormat[0:1])+textToFormat[1:]) + "."
}
