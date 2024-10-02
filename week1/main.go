package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func mapLetters(mapping map[rune]rune, start rune, limitIndex int) {
	// Map first half of the alphabet
	for i := start; i < start+rune(limitIndex); i++ {
		mapping[i] = start + rune(limitIndex-1) - (i - start)
	}

	// Map second half of the alphabet
	for i := start + rune(limitIndex); i < start+26; i++ {
		mapping[i] = start + 25 - (i - (start + rune(limitIndex)))
	}
}

func createLimitedAlphabetMap(limitIndex int) map[rune]rune {
	mapping := make(map[rune]rune)

	mapLetters(mapping, 'A', limitIndex)
	mapLetters(mapping, 'a', limitIndex)

	return mapping
}

func cypher(text string, limitedAlphabetMap map[rune]rune) string {
	var result strings.Builder

	for _, char := range text {
		if mappedChar, exists := limitedAlphabetMap[char]; exists {
			result.WriteRune(mappedChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var text string
	var limit int

	flag.StringVar(&text, "t", "Slizik", "text to cypher use quotation marks for sentences")
	flag.IntVar(&limit, "l", 26, "limit to split the alphabet")
	flag.Parse()

	if limit > 26 || limit < 0 {
		fmt.Println("Error: limitIndex cannot be greater than 26")
		os.Exit(1)
	}

	limitedAlphabetMap := createLimitedAlphabetMap(limit)

	cipheredText := cypher(text, limitedAlphabetMap)

	fmt.Printf("%s\n", cipheredText)
}
