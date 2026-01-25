package main

import (
	"fmt"
)

type Solution struct{}

func (s Solution) best(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	mini := prices[0]
	maxProfit := 0

	for _, price := range prices {
		cost := price - mini

		if cost > maxProfit {
			maxProfit = cost
		}

		if price < mini {
			mini = price
		}
	}

	return maxProfit
}

func main() {
	sol := Solution{}

	prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{7,6,4,3,1}

	fmt.Print("Best sol : ")
	bestRes := sol.best(prices)
	fmt.Println(bestRes)
}

