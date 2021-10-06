package lorelai

import (
	"math/rand"
	"strings"
	"time"
)

// Get a random Domain
func Domain() string {
	rand.Seed(time.Now().Unix())

	return strings.ToLower(Word()) + [5]string{"com", "net", "org", "io", "pe"}[rand.Intn(5)]
}

// Get a random URL
func URL() string {
	return "https://" + Domain()
}

// Get a random email
func Email() string {
	return Word() + "@" + Domain()
}
