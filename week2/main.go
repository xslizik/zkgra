package main

import (
	"flag"
	"fmt"
	"unicode"
)

func caesarCipher(text string, K int) string {
	Xni := make([]rune, len(text))
	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			Xni[i] = 'A' + (char-'A'+rune(K))%26
		} else if char >= 'a' && char <= 'z' {
			Xni[i] = 'a' + (char-'a'+rune(K))%26
		} else {
			Xni[i] = char
		}
	}
	return string(Xni)
}

func main() {
	input := flag.String("t", "Slizik", "Text to encrypt")
	flag.Parse()
	fmt.Printf("Ciphertext: %s\n", caesarCipher(*input, (int(unicode.ToLower(rune((*input)[0]))-'a'))%26+1))
}
