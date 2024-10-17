package main

import (
	"flag"
	"fmt"
	"math"
)

// Function to calculate (a^x) % n
func modExp(a, x, n int) int {
	return int(math.Pow(float64(a), float64(x))) % n
}

func generateTable(n int) [][]int {
	// The table will have n-1 rows and n-1 columns
	table := make([][]int, n-1)
	for a := 1; a < n; a++ {
		table[a-1] = make([]int, n-1)
		for x := 1; x < n; x++ {
			table[a-1][x-1] = modExp(a, x, n)
		}
	}

	return table
}

func printTable(n int, table [][]int) {
	fmt.Printf("a\\x  ")
	for x := 1; x < n; x++ {
		if x >= 10 {
			fmt.Printf("%d ", x) // header x
		} else {
			fmt.Printf(" %d ", x) // header x
		}
	}
	fmt.Println()

	for a := 1; a < n; a++ {
		if a >= 10 {
			fmt.Printf("%d   ", a) // header a
		} else {
			fmt.Printf(" %d   ", a) // header a
		}

		for x := 1; x < n; x++ {
			if table[a-1][x-1] >= 10 {
				fmt.Printf("%d ", table[a-1][x-1])
			} else {
				fmt.Printf(" %d ", table[a-1][x-1])
			}
		}
		fmt.Println()
	}
}

func findXForAandY(a, y int, table [][]int) []int {
	var xValues []int

	// Loop through the a value row to find all 'x' where the value matches 'y'
	for x, value := range table[a-1] {
		if value == y {
			xValues = append(xValues, x+1)
		}
	}

	return xValues
}

func main() {
	n := flag.Int("n", 8, "n")
	s := flag.Bool("s", false, "Suppress table")
	a := flag.Int("a", 0, "a")
	y := flag.Int("y", 0, "y")
	flag.Parse()

	table := generateTable(*n)
	if !*s {
		printTable(*n, table)
	}

	if *a > 0 && *y > 0 {
		xValues := findXForAandY(*a, *y, table)
		fmt.Printf("For a = %d and y = %d, x values are: %v\n", *a, *y, xValues)
	}
}
