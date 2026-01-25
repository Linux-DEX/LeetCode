package main

import (
	"fmt"
)

type Solution struct{}

// Brute force: O(n^2)
func (s Solution) Brute(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1} // 1-indexed
			}
		}
	}
	return []int{}
}

// Binary search: O(n log n)
func (s Solution) Best(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		low := i + 1
		high := len(numbers) - 1
		complement := target - numbers[i]

		for low <= high {
			mid := low + (high-low)/2
			if numbers[mid] == complement {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] < complement {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return []int{}
}

// Optimal approach: two pointers, O(n)
func (s Solution) Optimal(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // 1-indexed
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

// Helper function to print result
func PrintResult(label string, result []int) {
	fmt.Printf("%s : [", label)
	for i, val := range result {
		fmt.Print(val)
		if i != len(result)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func main() {
	sol := Solution{}

	numbers := []int{2, 7, 11, 15}
	target := 9

	PrintResult("Brute sol", sol.Brute(numbers, target))
	PrintResult("Best sol", sol.Best(numbers, target))
	PrintResult("Optimal sol", sol.Optimal(numbers, target))
}
