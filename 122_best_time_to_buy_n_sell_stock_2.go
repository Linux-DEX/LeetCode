package main

import (
	"fmt"
)

type Solution struct{}

// time complexity: O(n)
// space complexity: O(1)
func (s Solution) best(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func main() {
	sol := Solution{}

	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2, 3, 4, 5}

	fmt.Print("Best sol : ")
	bestRes := sol.best(prices)
	fmt.Println(bestRes)
}

