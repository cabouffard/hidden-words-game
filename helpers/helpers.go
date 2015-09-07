package helpers

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	value := rand.Intn(max-min) + min
	return value
}

func StringInSlice(query string, list []string) bool {
	for _, b := range list {
		if b == query {
			return true
		}
	}
	return false
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
