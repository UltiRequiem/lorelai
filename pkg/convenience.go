package lorelai

import (
	"math/rand"
	"strings"
	"time"
)

func Domain() string {
	rand.Seed(time.Now().Unix())

	return strings.ToLower(Word() + [5]string{"com", "net", "org", "io", "pe"}[rand.Intn(5)])
}

func URL() string {
	return "https://" + Domain()
}

func Email() string {
	return Word() + "@" + Domain()
}
