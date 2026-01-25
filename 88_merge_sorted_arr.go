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
