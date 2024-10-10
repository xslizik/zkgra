package main

import (
	"flag"
	"fmt"
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

func cipher(text string, limitedAlphabetMap map[rune]rune) string {
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

func caesarCipher(text string, shift int) string {
	shiftedText := make([]rune, len(text))
	for i, char := range text {
		switch {
		case char >= 'A' && char <= 'Z':
			shiftedText[i] = 'A' + (char-'A'+rune(shift))%26
		case char >= 'a' && char <= 'z':
			shiftedText[i] = 'a' + (char-'a'+rune(shift))%26
		default:
			shiftedText[i] = char
		}
	}
	return string(shiftedText)
}

func decipher(text string, shift int) string {
	return caesarCipher(text, -shift+26)
}

func bruteDecipher(input string) {
	limitResults := make([]string, 0)
	for limit := 0; limit < 26; limit++ {
		limitResults = append(limitResults, fmt.Sprintf("%s", cipher(input, createLimitedAlphabetMap(limit))))
	}

	shiftResults := make([]string, 0)
	for shift := 1; shift < 26; shift++ {
		shiftResults = append(shiftResults, fmt.Sprintf("%s", decipher(input, shift)))
	}

	fmt.Println(limitResults)
	fmt.Println(shiftResults)
}

func main() {
	input := flag.String("t", "Lebsbd", "text to cipher use quotation marks for sentences")
	flag.Parse()

	bruteDecipher(*input)
}
