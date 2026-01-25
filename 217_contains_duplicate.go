package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

/*
Time: O(n^2)
Space: O(1)
*/
func (s Solution) brute(arr []int) bool {
	size := len(arr)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

/*
Time: O(n log n)
Space: O(1) (ignoring sort space)
*/
func (s Solution) best(arr []int) bool {
	// copy slice to avoid modifying original
	nums := make([]int, len(arr))
	copy(nums, arr)

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

/*
Time: O(n)
Space: O(n)
*/
func (s Solution) optimal(arr []int) bool {
	set := make(map[int]bool)

	for _, v := range arr {
		if set[v] {
			return true
		}
		set[v] = true
	}
	return false
}

func main() {
	sol := Solution{}
	arr := []int{1, 4, 3, 4, 5}

	fmt.Println("Brute force :", sol.brute(arr))
	fmt.Println("Best solution :", sol.best(arr))
	fmt.Println("Optimal solution :", sol.optimal(arr))
}

