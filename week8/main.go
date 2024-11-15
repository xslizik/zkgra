package main

import (
	"flag"
	"fmt"
	"strings"
)

// Function to compute base^exp % mod using modular exponentiation
func modExponentiation(base, exp, mod int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

// Convert a number to a character (1 = 'a', 26 = 'z')
func numberToChar(num int) rune {
	if num >= 1 && num <= 26 {
		return rune('a' + num - 1)
	}
	return ' '
}

// Euclidean algorithm to find GCD of two numbers
func GCDEuclidean(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find values of e and d such that e*d ≡ 1 (mod fi)
func findFiCoprimes(fi, maxJ int) (int, int) {
	fmt.Printf("Possible e, d combinations\n")
	e_out := 0
	d_out := 0
	for e := 2; e < fi; e++ {
		if GCDEuclidean(e, fi) == 1 {
			for j := 0; j <= maxJ; j++ {
				if (1+fi*j)%e == 0 {
					d := (1 + fi*j) / e
					if d != 1 && e != d {
						fmt.Printf("e=%d d=%d j=%d\n", e, d, j)
						e_out = e
						d_out = d
						break
					}
				}
			}
		}
	}
	return e_out, d_out
}

func main() {
	// Define command-line flags
	word := flag.String("w", "bank", "Input message to encrypt or decrypt")
	encryptPrivate := flag.Bool("ep", false, "Use public key for encryption (default: false)")
	p := flag.Int("p", 9, "Prime number p")
	q := flag.Int("q", 11, "Prime number q")
	maxJ := flag.Int("j", 80, "Maximum j for finding d")

	flag.Parse()

	*word = strings.ToLower(*word)
	n := *p * *q
	fi := (*p - 1) * (*q - 1)

	var numbers []int
	var transformedNumbers []int
	var finalTransformedNumbers []int
	var finalChars []rune

	e, d := findFiCoprimes(fi, *maxJ)
	if e == 0 || d == 0 {
		fmt.Println("Failed to find suitable values for e and d")
		return
	}

	// Encryption process
	for _, char := range *word {
		var transformed int

		position := int(char - 'a' + 1)
		numbers = append(numbers, position)

		if *encryptPrivate {
			transformed = modExponentiation(position, d, n)
		} else {
			transformed = modExponentiation(position, e, n)
		}
		transformedNumbers = append(transformedNumbers, transformed)
	}

	// Decryption process
	for _, transformed := range transformedNumbers {
		var finalTransformed int

		if *encryptPrivate {
			finalTransformed = modExponentiation(transformed, e, n)
		} else {
			finalTransformed = modExponentiation(transformed, d, n)
		}
		finalTransformedNumbers = append(finalTransformedNumbers, finalTransformed)
		finalChars = append(finalChars, numberToChar(finalTransformed))
	}

	// Print results
	fmt.Printf("p: %d q: %d fi(n): %d\n", *p, *q, fi)
	fmt.Println("Input word:", *word)
	fmt.Printf("Public Key (e, n): (%d, %d)\n", e, n)
	fmt.Printf("Private Key (d, n): (%d, %d)\n", d, n)
	fmt.Println("Original message int:", numbers)
	if *encryptPrivate {
		fmt.Println("Encrypted message int (base^d mod n):", transformedNumbers)
		fmt.Println("Decrypted message int (base^e mod n):", finalTransformedNumbers)
	} else {
		fmt.Println("Encrypted message int (base^e mod n):", transformedNumbers)
		fmt.Println("Decrypted message int (base^d mod n):", finalTransformedNumbers)
	}
	fmt.Println("Decrypted message:", string(finalChars))
}
