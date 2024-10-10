package main

import (
	"flag"
	"fmt"
	"strings"
)

// Polybius Square Structure
type PolybiusSquare struct {
	encode map[rune]string
	decode map[string]rune
}

// Initialize the Polybius Square
func newPolybiusSquare() PolybiusSquare {
	encode := map[rune]string{
		'A': "11", 'B': "12", 'C': "13", 'D': "14", 'E': "15", 'F': "16",
		'G': "21", 'H': "22", 'I': "23", 'J': "24", 'K': "25", 'L': "26",
		'M': "31", 'N': "32", 'O': "33", 'P': "34", 'Q': "35", 'R': "36",
		'S': "41", 'T': "42", 'U': "43", 'V': "44", 'W': "45", 'X': "46",
		'Y': "51", 'Z': "52", '0': "53", '1': "54", '2': "55", '3': "56",
		'4': "61", '5': "62", '6': "63", '7': "64", '8': "65", '9': "66",
	}

	// Reverse mapping for decryption
	decode := make(map[string]rune)
	for char, code := range encode {
		decode[code] = char
	}

	return PolybiusSquare{encode: encode, decode: decode}
}

// Encrypts a message using the Polybius square
func (ps *PolybiusSquare) encrypt(message string) string {
	var encrypted strings.Builder
	for _, char := range message {
		if code, exists := ps.encode[char]; exists {
			encrypted.WriteString(code)
		} else if char == ' ' { // Keep spaces as is
			encrypted.WriteString(" ")
		}
	}
	return encrypted.String()
}

// Decrypts a message using the Polybius square
func (ps *PolybiusSquare) decrypt(encoded string) string {
	var decrypted strings.Builder
	// Split the encoded string by spaces
	words := strings.Split(encoded, " ")
	for _, word := range words {
		for i := 0; i < len(word); i += 2 {
			if i+1 < len(word) {
				code := word[i : i+2]
				if char, exists := ps.decode[code]; exists {
					decrypted.WriteRune(char)
				}
			}
		}
		decrypted.WriteRune(' ') // Add space after each word
	}
	return strings.TrimSpace(decrypted.String()) // Trim trailing space
}

func main() {
	input := flag.String("t", "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "text to cipher")
	flag.Parse()
	ps := newPolybiusSquare()

	encoded := ps.encrypt(*input)
	fmt.Printf("Encoded: %s\n", encoded)

	decoded := ps.decrypt(encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
