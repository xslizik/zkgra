package main

import (
	"flag"
	"fmt"
)

// polynomialHashWeighted calculates a weighted polynomial hash for a string
func polynomialHashWeighted(s string, p, m int, weights map[rune]int) int {
	hash := 0
	power := 1

	for _, char := range s {
		// Get the weight of the character; default to 1 if not in map
		weight, exists := weights[char]
		if !exists {
			weight = 1
		}
		// Update the hash with the weighted value
		hash = (hash + weight*int(char)*power) % m
		// Update the power of p modulo m
		power = (power * p) % m
	}

	return hash
}

func generateAlphabetWeights() map[rune]int {
	weights := make(map[rune]int)
	// Assign weights to a-Z
	for i, char := range "vRTisVGaCkeFyHNOjqfuPLcSbrYQEMxpoIhBgJlXtKwAUzmWdnDZ" {
		weights[char] = i + 1
	}
	return weights
}

func main() {
	// Generate weights for the entire alphabet
	weights := generateAlphabetWeights()

	p := flag.Int("p", 31, "Small prime number")
	m := flag.Int("m", 7417, "Large prime number to prevent overflow")
	input := flag.String("i", "TRALALA", "Message to be processed")
	flag.Parse()

	hashValue := polynomialHashWeighted(*input, *p, *m, weights)

	fmt.Printf("The weighted polynomial hash of '%s' is: %d\n", *input, hashValue)

}
