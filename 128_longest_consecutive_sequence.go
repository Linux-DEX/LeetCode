/*
# Longest Consecutive Sequence

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Example 3:

Input: nums = [1,0,1,2]
Output: 3
*/

package main

import (
	"fmt"
	algo "linux-dex/leetcode/algorithms"
	"math"
)

func ls(arr []int, num int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
}

// Brute Solution : O(n^2) time, O(1) space
func longestConsecutiveBrute(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	longest := 1

	for i := 0; i < len(nums); i++ {

		x := nums[i]
		count := 1

		// check for next consecutive numbers
		for ls(nums, x+1) {
			x = x + 1
			count++
		}

		if count > longest {
			longest = count
		}
	}

	return longest
}

// Better Solution : O(n log n) time, O(1) space
func longestConsecutiveBetter(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nums = algo.MergeSort(nums)
	n := len(nums)
	lastSmaller := math.MinInt64
	cnt := 0
	longest := 1

	for i := 0; i < n; i++ {
		if nums[i]-1 == lastSmaller {
			cnt += 1
			lastSmaller = nums[i]
		} else if lastSmaller != nums[i] {
			cnt = 1
			lastSmaller = nums[i]
		}
		longest = max(longest, cnt)
	}

	return longest
}

// Optimal Solution : O(n) time, O(n) space
func longestConsecutiveOptimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	longest := 0

	for num := range set {
		// start of sequence
		if !set[num-1] {
			current := num
			count := 1

			for set[current+1] {
				current++
				count++
			}

			if count > longest {
				longest = count
			}
		}
	}

	return longest
}

func main() {
	fmt.Println("Brute Solution : ")
	fmt.Println(longestConsecutiveBrute([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveBrute([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutiveBrute([]int{1, 0, 1, 2}))
	fmt.Println()

	fmt.Println("Better Solution : ")
	fmt.Println(longestConsecutiveBetter([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveBetter([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutiveBetter([]int{1, 0, 1, 2}))
	fmt.Println()

	fmt.Println("Optimal Solution : ")
	fmt.Println(longestConsecutiveOptimal([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveOptimal([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutiveOptimal([]int{1, 0, 1, 2}))
}
