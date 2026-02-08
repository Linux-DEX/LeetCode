/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
*/
package main

import (
	"fmt"
)

type Solution struct{}

// -------------------------
// Best: Merge with extra space
// Time  : O(n + m)
// Space : O(n + m)
// -------------------------
func (s Solution) Best(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	left, right, index := 0, 0, 0

	for left < m && right < n {
		if nums1[left] <= nums2[right] {
			nums3[index] = nums1[left]
			left++
		} else {
			nums3[index] = nums2[right]
			right++
		}
		index++
	}

	for left < m {
		nums3[index] = nums1[left]
		left++
		index++
	}

	for right < n {
		nums3[index] = nums2[right]
		right++
		index++
	}

	// Copy back to nums1
	for i := 0; i < m+n; i++ {
		nums1[i] = nums3[i]
	}
}

// -------------------------
// Optimal: Merge in-place from end
// Time  : O(m + n)
// Space : O(1)
// -------------------------
func (s Solution) Optimal(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func main() {
	sol := Solution{}

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	fmt.Print("Best sol : ")
	sol.Best(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}

	fmt.Print("Optimal sol : ")
	sol.Optimal(nums1, m, nums2, n)
	fmt.Println(nums1)
}
