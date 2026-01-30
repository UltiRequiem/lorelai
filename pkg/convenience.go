package lorelai

import (
	"math/rand/v2"
	"strings"
)

// Get a random Domain
func Domain() string {
	return trimSpaceAddDot(LoremWords(1)) + TLDS[rand.IntN(len(TLDS))]
}

// Get a random URL
func URL() string {
	return "https://" + Domain()
}

// Get a random email
func Email() string {
	return strings.TrimSpace(LoremWords(1)) + "@" + Domain()
}
