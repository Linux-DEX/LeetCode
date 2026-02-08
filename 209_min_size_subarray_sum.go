/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/
package main

import (
	"fmt"
	"math"
)

// minSubArrayLen finds the minimal length of a contiguous subarray
// of which the sum is greater than or equal to target.
// Returns 0 if no such subarray exists.
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := math.MaxInt32
	sum := 0
	left := 0

	for right := 0; right < n; right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	result := minSubArrayLen(target, nums)
	fmt.Println("Minimum length of subarray:", result)
}
