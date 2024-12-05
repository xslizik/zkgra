package main

import (
	"fmt"
	"strings"
)

func main() {
	ciphertext := []int{20, 24, 12}

	for x := 1; x <= 26; x++ {
		decodedText := decode(ciphertext, x)
		if decodedText != "" {
			fmt.Printf("Key [x] = %d, Decoded Message [m] = %s\n", x, decodedText)
		}
	}
}

func decode(ciphertext []int, key int) string {
	var decoded strings.Builder
	for _, num := range ciphertext {
		decodedNum := (num - key)
		if decodedNum < 1 {
			decodedNum += 26
		}
		letter := rune(decodedNum + 'A' - 1)
		decoded.WriteRune(letter)
	}
	return decoded.String()
}
