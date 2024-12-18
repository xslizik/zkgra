package main

import (
	"fmt"
	"sort"
)

func createTable(probabilities []float64) [][]float64 {
	var table [][]float64

	sort.Sort(sort.Reverse(sort.Float64Slice(probabilities)))

	table = append(table, append([]float64(nil), probabilities...))

	// Iterate until the sum of probabilities is 1
	for len(probabilities) > 1 {
		// Add the last two elements
		last := probabilities[len(probabilities)-1]
		secondLast := probabilities[len(probabilities)-2]

		// Remove the last two elements
		probabilities = probabilities[:len(probabilities)-2]

		// Add their sum back into the probabilities
		probabilities = append(probabilities, last+secondLast)

		// Sort the probabilities again (descending order)
		sort.Sort(sort.Reverse(sort.Float64Slice(probabilities)))

		// Add the current column to the table
		table = append(table, append([]float64(nil), probabilities...))
	}

	return table
}

func printTable(table [][]float64) {
	for i, row := range table {
		fmt.Printf("Column %d: ", i+1)
		for _, prob := range row {
			fmt.Printf("%.2f ", prob)
		}
		fmt.Println()
	}
}

func main() {
	probabilities := []float64{0.25, 0.20, 0.18, 0.15, 0.10, 0.08, 0.03, 0.01}

	table := createTable(probabilities)
	printTable(table)
}
