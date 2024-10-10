package main

import (
	"fmt"
	"strconv"
)

func xorBinaryValues(values []string) string {
	// create hashmap of occurences
	occurrences := make(map[string]int)

	// find number of occurences
	for _, value := range values {
		occurrences[value]++
	}

	result := 0
	// perform xor only on those that have odd number of occurences
	for value, count := range occurrences {
		if count%2 != 0 {
			num, _ := strconv.ParseInt(value, 2, 64)
			result ^= int(num)
		}
	}

	return strconv.FormatInt(int64(result), 2)
}

func main() {
	binarySets := [][]string{
		{"1011", "0110", "0100"},
		{"0101", "1110", "1101"},
		{"0001", "0101", "1010"},
		{"1010", "1010"},
		{"1000", "1100"},
	}

	for i, set := range binarySets {
		if i < 3 {
			result := xorBinaryValues([]string{set[0], set[1], set[2], set[0], set[1]})
			fmt.Printf("For a = %s, b = %s, c = %s: a^b^c^a^b = %s\n", set[0], set[1], set[2], result)
		} else if i == 3 {
			result := xorBinaryValues([]string{set[0], set[0], set[0], set[0], set[0]})
			fmt.Printf("For a = %s: a^a^a^a^a = %s\n", set[0], result)
		} else {
			result := xorBinaryValues([]string{set[0], set[1], set[1], set[1], set[1]})
			fmt.Printf("For a = %s, b = %s: a^b^b^b^b = %s\n", set[0], set[1], result)
		}
	}
}
