/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and using only constant extra space.

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2

Example 2:

Input: nums = [3,1,3,4,2]
Output: 3

Example 3:

Input: nums = [3,3,3,3,3]
Output: 3
*/
package main

import (
	"fmt"
	"linux-dex/leetcode/algorithms"
)

// Brute Solution: O(n log n) time, O(1) space
func findDuplicateBrute(nums []int) int {
	// sorting array
	sorted := algorithms.MergeSort(nums)

	// traversal
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] == sorted[i+1] {
			return sorted[i]
		}
	}
	return -1
}

// Best Solution: O(n) time, O(n) space
func findDuplicateBest(nums []int) int {
	// hash map to track element
	seen := make(map[int]bool)

	// traverse
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}

	return -1
}

// Optimal Solution: O(n) time, O(1) space
func findDuplicateOptimal(nums []int) int {
	// finding intersection point in the cycle
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// finding entrance to the cycle
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {
	nums1 := []int{1, 3, 4, 2, 2}
	nums2 := []int{3, 1, 3, 4, 2}
	nums3 := []int{3, 3, 3, 3, 3}

	fmt.Println("Brute solution :")
	fmt.Println(findDuplicateBrute(nums1))
	fmt.Println(findDuplicateBrute(nums2))
	fmt.Println(findDuplicateBrute(nums3))

	fmt.Println("Best solution :")
	fmt.Println(findDuplicateBest(nums1))
	fmt.Println(findDuplicateBest(nums2))
	fmt.Println(findDuplicateBest(nums3))

	fmt.Println("Optimal solution :")
	fmt.Println(findDuplicateOptimal(nums1))
	fmt.Println(findDuplicateOptimal(nums2))
	fmt.Println(findDuplicateOptimal(nums3))
}
