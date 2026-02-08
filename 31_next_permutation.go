/*
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.

Example 1:

Input: nums = [1,2,3]
Output: [1,3,2]

Example 2:

Input: nums = [3,2,1]
Output: [1,2,3]

Example 3:

Input: nums = [1,1,5]
Output: [1,5,1]
*/
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
