package main

import (
	"fmt"
)

func maxChickensProtected(n int, k int, positions []int) int {
	maxChickens := 0
	start := 0

	// slicing window technique
	for end := 0; end < n; end++ {
		for positions[end]-positions[start] >= k {
			start++
		}
		currentChickens := end - start + 1
		if currentChickens > maxChickens {
			maxChickens = currentChickens
		}
	}

	return maxChickens
}

func main() {
	// Test cases
	n1 := 5
	k1 := 5
	positions1 := []int{2, 5, 10, 12, 15}
	fmt.Println(maxChickensProtected(n1, k1, positions1))

	n2 := 6
	k2 := 10
	positions2 := []int{1, 11, 30, 34, 35, 37}
	fmt.Println(maxChickensProtected(n2, k2, positions2))
}
