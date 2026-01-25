package main

import (
	"fmt"
)

type Solution struct{}

// -------------------------
// Brute Force
// Time  : O(n)
// Space : O(n)
// -------------------------
func (s Solution) Brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	unique := []int{}

	for _, num := range nums {
		if len(unique) == 0 || num != unique[len(unique)-1] {
			unique = append(unique, num)
		}
	}

	// Copy back to nums
	for i := 0; i < len(unique); i++ {
		nums[i] = unique[i]
	}

	return len(unique)
}

// -------------------------
// Optimal: In-place
// Time  : O(n)
// Space : O(1)
// -------------------------
func (s Solution) Optimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	sol := Solution{}

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	bruteRes := sol.Brute(nums)
	fmt.Print("Brute sol   :", bruteRes, " => ")
	fmt.Println(nums)

	// Reset nums for optimal test
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	optimalRes := sol.Optimal(nums)
	fmt.Print("Optimal sol :", optimalRes, " => ")
	fmt.Println(nums)
}
