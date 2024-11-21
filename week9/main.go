package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a random odd or even number
func generateNumber(isEven bool, r *rand.Rand) int {
	num := r.Intn(62) // max number that can be generated is 0-63
	if isEven {
		if num%2 != 0 {
			num++
		}
	} else {
		if num%2 == 0 {
			num++
		}
	}
	return num
}

// Hash function h1(n) = n mod m
func hash1(n, m int) int {
	return n % m
}

// Hash function h2(n) = (h1(n1) + h1(n2) + h1(n3) + h1(n4) + h1(n5) + h1(n6)) mod m
func hash2(n []int, m int) int {
	sum := 0
	for _, ni := range n {
		sum += hash1(ni, m)
	}
	return sum % m
}

// Generate results and print as a table
func generateResults(caseName string, isNEven, isMOdd bool) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	fmt.Printf("\nResults for Case %s:\n", caseName)
	fmt.Printf("| n     | m     | n1    | n2    | n3    | n4    | n5    | n6    | h1(n, m) | h2(n, m) |\n")
	fmt.Printf("|-------|-------|-------|-------|-------|-------|-------|-------|----------|----------|\n")
	for i := 0; i < 5; i++ {
		n := generateNumber(isNEven, r)
		m := generateNumber(!isMOdd, r)
		n1 := generateNumber(isNEven, r)
		n2 := generateNumber(isNEven, r)
		n3 := generateNumber(isNEven, r)
		n4 := generateNumber(isNEven, r)
		n5 := generateNumber(isNEven, r)
		n6 := generateNumber(isNEven, r)

		h1 := hash1(n, m)
		h2 := hash2([]int{n1, n2, n3, n4, n5, n6}, m)

		fmt.Printf("| %5d | %5d | %5d | %5d | %5d | %5d | %5d | %5d | %8d | %8d |\n",
			n, m, n1, n2, n3, n4, n5, n6, h1, h2)
	}
}

func main() {
	fmt.Printf("h1(n) = n mod m\nh2(n) = (h1(n1) + h1(n2) + h1(n3) + h1(n4) + h1(n5) + h1(n6)) mod m\n")
	generateResults("Even n and Odd m", true, true)
	generateResults("Odd n and Even m", false, false)
}
