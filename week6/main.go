package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

// Initial Permutation table remains the same
var IP = [][]int{
	{16, 19, 5, 1, 13},
	{14, 4, 21, 10, 8},
	{24, 11, 3, 12, 22},
	{17, 9, 20, 7, 18},
	{23, 2, 6, 15, 25},
}

// Check if string contains only numbers
func isNumeric(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Convert input to binary - handles both text and numbers
func toBinary(s string) string {
	var binary strings.Builder

	if isNumeric(s) {
		// Handle numbers using 5 bits per digit
		for _, c := range s {
			val := int(c - '0') // Convert digit character to integer
			binary.WriteString(fmt.Sprintf("%05b", val))
		}
	} else {
		// Handle text using 5 bits per character
		for _, c := range s {
			val := int(c) % 32
			binary.WriteString(fmt.Sprintf("%05b", val))
		}
	}

	result := binary.String()
	if len(result) < 25 {
		result += strings.Repeat("0", 25-len(result))
	}
	return result[:25]
}

// Convert binary back to original format
func fromBinary(binary string, wasNumeric bool) string {
	var result strings.Builder

	for i := 0; i < len(binary); i += 5 {
		if i+5 > len(binary) {
			break
		}

		// Convert the 5-bit binary substring to an integer
		val := 0
		for j := 0; j < 5; j++ {
			val = val*2 + int(binary[i+j]-'0')
		}

		if wasNumeric {
			// For numbers, directly convert value to string
			if val < 10 { // Only add valid digits
				result.WriteString(fmt.Sprintf("%d", val))
			}
		} else {
			// For text, map to uppercase letters
			if val > 0 {
				result.WriteRune(rune(val + 64))
			}
		}
	}
	return result.String()
}

func padToLength(input string, length int) string {
	if len(input) < length {
		return input + strings.Repeat("0", length-len(input))
	}
	return input[:length]
}

func initialPermutation(input string) string {
	result := make([]byte, 25)
	for i, row := range IP {
		for j, val := range row {
			if val <= len(input) {
				result[i*5+j] = input[val-1]
			} else {
				result[i*5+j] = '0'
			}
		}
	}
	return string(result)
}

func finalPermutation(input string) string {
	input = padToLength(input, 25)
	result := make([]byte, 25)
	for i := 0; i < 25; i++ {
		found := false
		for rowIdx, row := range IP {
			for colIdx, val := range row {
				if val == i+1 {
					result[i] = input[rowIdx*5+colIdx]
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			result[i] = '0'
		}
	}
	return string(result)
}

func xorStrings(a, b string) string {
	length := len(a)
	if len(b) > length {
		length = len(b)
	}

	a = padToLength(a, length)
	b = padToLength(b, length)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}

func feistelFunction(R string, roundKey string) string {
	paddedR := padToLength(R, len(roundKey))
	return xorStrings(paddedR, roundKey)
}

func encrypt(message string, key string) (string, bool) {
	fmt.Printf("\n--- Starting Encryption ---\n")
	isNumber := isNumeric(message)
	binary := toBinary(strings.ToUpper(message))
	fmt.Printf("Binary input: %s (%d bits)\n", binary, len(binary))

	permuted := initialPermutation(binary)
	fmt.Printf("After initial permutation: %s (%d bits)\n", permuted, len(permuted))

	halfLength := len(permuted) / 2
	L := permuted[:halfLength]
	R := permuted[halfLength:]

	fmt.Printf("Initial split: L=%s (%d bits), R=%s (%d bits)\n", L, len(L), R, len(R))

	// Round 1
	R1 := feistelFunction(R, key)
	R1 = padToLength(R1, len(L))
	newR := xorStrings(L, R1)
	newL := R
	fmt.Printf("After round 1: L=%s (%d bits), R=%s (%d bits)\n", newL, len(newL), newR, len(newR))

	// Round 2
	R2 := feistelFunction(newR, key)
	R2 = padToLength(R2, len(newL))
	finalR := xorStrings(newL, R2)
	finalL := newR
	fmt.Printf("After round 2: L=%s (%d bits), R=%s (%d bits)\n", finalL, len(finalL), finalR, len(finalR))

	combined := finalL + finalR
	result := finalPermutation(combined)
	fmt.Printf("Final encrypted result: %s (%d bits)\n", result, len(result))
	return result, isNumber
}

func decrypt(ciphertext string, key string, wasNumeric bool) string {
	fmt.Printf("\n--- Starting Decryption ---\n")
	fmt.Printf("Binary input: %s (%d bits)\n", ciphertext, len(ciphertext))

	permuted := initialPermutation(ciphertext)
	fmt.Printf("After initial permutation: %s (%d bits)\n", permuted, len(permuted))

	halfLength := len(permuted) / 2
	L := permuted[:halfLength]
	R := permuted[halfLength:]

	fmt.Printf("Initial split: L=%s (%d bits), R=%s (%d bits)\n", L, len(L), R, len(R))

	// Round 2
	R2 := feistelFunction(L, key)
	R2 = padToLength(R2, len(R))
	newR := L
	newL := xorStrings(R, R2)
	fmt.Printf("After round 2: L=%s (%d bits), R=%s (%d bits)\n", newL, len(newL), newR, len(newR))

	// Round 1
	R1 := feistelFunction(newL, key)
	R1 = padToLength(R1, len(newR))
	finalR := newL
	finalL := xorStrings(newR, R1)
	fmt.Printf("After round 1: L=%s (%d bits), R=%s (%d bits)\n", finalL, len(finalL), finalR, len(finalR))

	combined := finalL + finalR
	decryptedBinary := finalPermutation(combined)
	decryptedMessage := fromBinary(decryptedBinary, wasNumeric)

	return decryptedMessage
}

func main() {
	message := flag.String("m", "110892", "Input message max 5 char case insensitive (default: 110892)")
	key := flag.String("k", "1010111011100110100001010", "Encryption key (default: 1010111010000100110000110)")

	flag.Parse()

	fmt.Printf("Original Message: %s\n", *message)

	encrypted, isNumber := encrypt(*message, *key)
	fmt.Printf("\nEncrypted Message: %s\n", encrypted)

	decrypted := decrypt(encrypted, *key, isNumber)
	fmt.Printf("\nDecrypted Message: %s\n", decrypted)
}
