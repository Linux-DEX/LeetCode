package main

import (
	"fmt"
)

type Solution struct{}

// Best approach: preserve order
func (s Solution) Best(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// Optimal approach: change order allowed
func (s Solution) Optimal(nums []int, val int) int {
	size := len(nums)
	i := 0
	for i < size {
		if nums[i] == val {
			nums[i] = nums[size-1]
			size--
		} else {
			i++
		}
	}
	return size
}

func main() {
	sol := Solution{}

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	// Best solution
	numsCopy := append([]int(nil), nums...) // make a copy
	k := sol.Best(numsCopy, val)
	fmt.Printf("Best sol : %d -> ", k)
	for _, n := range numsCopy {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

	// Optimal solution
	numsCopy = append([]int(nil), nums...) // make a copy
	k = sol.Optimal(numsCopy, val)
	fmt.Printf("Optimal sol : %d -> ", k)
	for _, n := range numsCopy {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}
