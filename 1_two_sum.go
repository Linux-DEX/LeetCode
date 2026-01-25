package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

// Brute force: O(n^2)
func (s Solution) Brute(arr []int, target int) (int, int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] > target {
			continue
		}
		val := target - arr[i]
		for j := i + 1; j < size; j++ {
			if arr[j] == val {
				return i, j
			}
		}
	}
	return -1, -1
}

// Binary search helper
func BinarySearch(arr []int, start int, target int) int {
	left, right := start, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Best approach: sort + binary search: O(n log n)
func (s Solution) Best(arr []int, target int) (int, int) {
	size := len(arr)
	// Make a copy to avoid modifying original
	nums := make([]int, size)
	copy(nums, arr)

	sort.Ints(nums)

	for i := 0; i < size; i++ {
		if nums[i] > target {
			continue
		}
		val := target - nums[i]
		index := BinarySearch(nums, i+1, val)
		if index != -1 {
			return i, index
		}
	}
	return -1, -1
}

// Optimal approach: hash map O(n)
func (s Solution) Optimal(arr []int, target int) (int, int) {
	hashmap := make(map[int]int)
	for i, num := range arr {
		complement := target - num
		if idx, found := hashmap[complement]; found {
			return idx, i
		}
		hashmap[num] = i
	}
	return -1, -1
}

func main() {
	sol := Solution{}

	arr := []int{15, 7, 2, 11}
	target := 9

	fmt.Print("Brute force : ")
	i, j := sol.Brute(arr, target)
	fmt.Println(i, j)

	fmt.Print("Best sol     : ")
	i, j = sol.Best(arr, target)
	fmt.Println(i, j)

	fmt.Print("Optimal sol  : ")
	i, j = sol.Optimal(arr, target)
	fmt.Println(i, j)
}
