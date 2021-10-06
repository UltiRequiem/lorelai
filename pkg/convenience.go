package lorelai

import (
	"math/rand"
	"time"
)

func Domain() string {
	rand.Seed(time.Now().Unix())

	return Word() + [5]string{".com", ".net", ".org", ".io", ".pe"}[rand.Intn(5)]
}

func URL() string {
	return "https://" + Domain()
}

func Email() string {
	return Word() + "@" + Domain()
}
