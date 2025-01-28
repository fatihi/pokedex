package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		output := fmt.Sprintf("Your command was: %s", words[0])
		fmt.Println(output)
	}
}
