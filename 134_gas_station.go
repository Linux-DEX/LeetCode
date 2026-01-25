package main

import (
	"fmt"
)

type Solution struct{}

// Brute Force - O(n^2)
func (s Solution) best(gas []int, cost []int) int {
	n := len(gas)

	for start := 0; start < n; start++ {
		currentGas := 0
		canComplete := true

		for i := 0; i < n; i++ {
			idx := (start + i) % n

			currentGas += gas[idx]
			currentGas -= cost[idx]

			if currentGas < 0 {
				canComplete = false
				break
			}
		}

		if canComplete {
			return start
		}
	}

	return -1
}

// Optimal - O(n)
func (s Solution) optimal(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currentGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		currentGas += gas[i] - cost[i]

		// reset starting point
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	if totalGas >= totalCost {
		return start
	}
	return -1
}

func main() {
	sol := Solution{}

	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Print("Best sol : ")
	bestRes := sol.best(gas, cost)
	fmt.Println(bestRes)

	fmt.Print("Optimal sol : ")
	optimalRes := sol.optimal(gas, cost)
	fmt.Println(optimalRes)
}

