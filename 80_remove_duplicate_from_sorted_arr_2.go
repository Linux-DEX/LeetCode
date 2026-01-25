package main

import (
	"fmt"
)

type Solution struct{}

// Brute-force: O(n) time, O(n) space
func (s *Solution) Brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := []int{}

	for _, num := range nums {
		if len(result) < 2 || result[len(result)-2] != num {
			result = append(result, num)
		}
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	return len(result)
}

// Optimal: O(n) time, O(1) space
func (s *Solution) Optimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1 // write pointer
	count := 1

	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[j-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func main() {
	sol := Solution{}

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Print("Brute sol : ")
	bruteRes := sol.Brute(nums)
	fmt.Printf("%d => ", bruteRes)
	for _, n := range nums {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

	// reset nums
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Print("Optimal sol : ")
	optimalRes := sol.Optimal(nums)
	fmt.Printf("%d => ", optimalRes)
	for _, n := range nums {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}
