package main

import (
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	value := rand.Intn(max-min) + min
	return value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stringInSlice(query string, list []string) bool {
	for _, b := range list {
		if b == query {
			return true
		}
	}
	return false
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
