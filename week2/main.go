package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

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
	for shift := 0; shift < 26; shift++ {
		fmt.Printf("Shift %d: %s\n", shift, decipher(input, shift))
	}
}

func main() {
	input := flag.String("t", "Slizik", "Text to encrypt or decrypt")
	shouldDecrypt := flag.Bool("d", false, "Set to true for decryption mode")
	shift := flag.Int("k", 0, "Shift value (for known shifts)")
	flag.Parse()

	if *shift > 26 || *shift < 0 {
		fmt.Println("Error: limitIndex cannot be greater than 26")
		os.Exit(1)
	}

	if *shouldDecrypt {
		if *shift > 0 && *shift < 26 {
			fmt.Println(decipher(*input, *shift))
		} else {
			bruteDecipher(*input)
		}
	} else {
		fmt.Println(caesarCipher(*input, (int(unicode.ToLower(rune((*input)[0]))-'a')%26)+1))
	}
}
