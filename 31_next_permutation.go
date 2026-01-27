package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

// Brute-force: O(n! x n) time, O(n! x n) space
func (s *Solution) brute(nums []int) {
	var perms [][]int
	n := len(nums)

	var backtrack func(int)
	backtrack = func(first int) {
		if first == n {
			tmp := make([]int, n)
			copy(tmp, nums)
			perms = append(perms, tmp)
			return
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)

	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < n; k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})

	for i := 0; i < len(perms); i++ {
		match := true
		for j := 0; j < n; j++ {
			if perms[i][j] != nums[j] {
				match = false
				break
			}
		}
		if match {
			next := perms[(i+1)%len(perms)]
			copy(nums, next)
			return
		}
	}
}

// Optimal-Sol: O(n) time, O(1) space
func (s *Solution) optimal(nums []int) {
	n := len(nums)

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums, i+1, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	sol := Solution{}

	nums1 := []int{1, 2, 3}
	nums2 := []int{3, 2, 1}
	nums3 := []int{1, 1, 5}

	fmt.Print("Brute sol : ")
	sol.brute(nums1)
	sol.brute(nums2)
	sol.brute(nums3)

	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)

	nums1 = []int{1, 2, 3}
	nums2 = []int{3, 2, 1}
	nums3 = []int{1, 1, 5}

	fmt.Print("Optimal sol : ")
	sol.optimal(nums1)
	sol.optimal(nums2)
	sol.optimal(nums3)

	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)
}
