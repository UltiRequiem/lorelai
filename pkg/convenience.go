package lorelai

import (
	"math/rand"
	"strings"
)

// Get a random Domain
func Domain() string {
	return trimSpaceAddDot(LoremWords(1)) + TLDS[rand.Intn(len(TLDS))]
}

// Get a random URL
func URL() string {
	return "https://" + Domain()
}

// Get a random email
func Email() string {
	return strings.TrimSpace(LoremWords(1)) + "@" + Domain()
}
