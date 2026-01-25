package main

import (
	"fmt"
)

type Solution struct{}

// -------------------------
// Brute Force
// Time  : O(n^2)
// Space : O(1)
// -------------------------
func (s Solution) Brute(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > n/2 {
			return nums[i]
		}
	}

	return -1
}

// -------------------------
// Hashmap Approach
// Time  : O(n)
// Space : O(n)
// -------------------------
func (s Solution) Best(nums []int) int {
	n := len(nums)
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
		if countMap[num] > n/2 {
			return num
		}
	}

	return -1
}

// -------------------------
// Boyer-Moore Voting Algorithm
// Time  : O(n)
// Space : O(1)
// -------------------------
func (s Solution) Optimal(nums []int) int {
	count := 0
	candidate := -1

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	sol := Solution{}

	nums := []int{3, 2, 3}
	// nums := []int{4, 4, 1, 1, 1, 4, 4}

	bruteRes := sol.Brute(nums)
	fmt.Println("Brute sol   :", bruteRes)

	bestRes := sol.Best(nums)
	fmt.Println("Best sol    :", bestRes)

	optimalRes := sol.Optimal(nums)
	fmt.Println("Optimal sol :", optimalRes)
}
