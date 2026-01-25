package main

import "fmt"

type Solution struct{}

// Helper to print array
func (s Solution) PrintArr(nums []int) {
	for _, n := range nums {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}

// Brute approach: rotate one by one (O(n*k), O(1))
func (s Solution) Brute(nums []int, k int) {
	size := len(nums)
	if size == 0 {
		return
	}

	k = k % size

	for i := 0; i < k; i++ {
		last := nums[size-1]
		for j := size - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// Best approach: use extra array (O(n), O(n))
func (s Solution) Best(nums []int, k int) {
	size := len(nums)
	if size == 0 {
		return
	}

	k = k % size
	rotate := make([]int, size)

	for i := 0; i < size; i++ {
		rotate[(i+k)%size] = nums[i]
	}

	for i := 0; i < size; i++ {
		nums[i] = rotate[i]
	}
}

// Helper to reverse a slice in-place
func (s Solution) Reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Optimal approach: reverse parts of array (O(n), O(1))
func (s Solution) Optimal(nums []int, k int) {
	size := len(nums)
	if size == 0 {
		return
	}

	k = k % size
	if k == 0 {
		return
	}

	s.Reverse(nums, 0, size-1)
	s.Reverse(nums, 0, k-1)
	s.Reverse(nums, k, size-1)
}

func main() {
	sol := Solution{}

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Print("Brute sol   : ")
	sol.Brute(nums, k)
	sol.PrintArr(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print("Best sol    : ")
	sol.Best(nums, k)
	sol.PrintArr(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print("Optimal sol : ")
	sol.Optimal(nums, k)
	sol.PrintArr(nums)
}
