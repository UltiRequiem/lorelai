package lorelai

import "strings"

func capitalizeFirstWord(text string) string {
    if len(text) == 0 {
        return text
    }
    return strings.ToUpper(text[:1]) + text[1:]
}

// Trim spaces and adds a dot at the end.
func formatWords(textToFormat string) string {
	return trimSpaceAddDot(capitalizeFirstWord(textToFormat))
}

func trimSpaceAddDot(text string) string {
	return strings.TrimSpace(text) + "."
}
