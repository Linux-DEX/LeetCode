/*
# Reverse Pairs

Given an integer array nums, return the number of reverse pairs in the array.

A reverse pair is a pair (i, j) where:

0 <= i < j < nums.length and
nums[i] > 2 * nums[j].

Example 1:

Input: nums = [1,3,2,3,1]
Output: 2
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1

Example 2:

Input: nums = [2,4,3,5,1]
Output: 3
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
(2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
*/

package main

import (
	"fmt"
)

// Brute Solution : O(n^2) time, O(1) space
func reversePairsBrute(nums []int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > 2*nums[j] {
				count++
			}
		}
	}

	return count
}

// Optimal Solution : O(n log n) time, O(n) space
func reversePairsOptimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSortAndCount(nums, 0, len(nums)-1)
}

func mergeSortAndCount(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	count := 0

	count += mergeSortAndCount(nums, left, mid)
	count += mergeSortAndCount(nums, mid+1, right)
	count += countReversePairs(nums, left, mid, right)
	merge(nums, left, mid, right)

	return count
}

func countReversePairs(nums []int, left, mid, right int) int {
	count := 0
	j := mid + 1

	for i := left; i <= mid; i++ {
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		count += j - (mid + 1)
	}

	return count
}

func merge(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}

	for i := left; i <= right; i++ {
		nums[i] = temp[i-left]
	}
}

func main() {
	fmt.Println("Brute Solution : ")
	fmt.Println(reversePairsBrute([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairsBrute([]int{2, 4, 3, 5, 1}))
	fmt.Println()

	fmt.Println("Optimal Solution : ")
	fmt.Println(reversePairsOptimal([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairsOptimal([]int{2, 4, 3, 5, 1}))
}
