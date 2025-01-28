package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		word = strings.ToLower(word)
		words[i] = word
	}
	return words
}

func main() {
	fmt.Println("Hello, World!")
}
