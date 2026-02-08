/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/
package main

import (
	"fmt"
)

type Solution struct{}

// Brute-force: O(n^3) time, O(1) space
func (s *Solution) brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_val := nums[0]
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			max_val = max(sum, max_val)
		}
	}

	return max_val
}

// Better-Sol: O(n^2) time, O(1) space
func (s *Solution) better(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_val := nums[0]
	n := len(nums)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			max_val = max(sum, max_val)
		}
	}

	return max_val
}

// Optimal-Sol: O(n) time, O(1) space
func (s *Solution) optimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_val := nums[0]
	current_sum := nums[0]

	for i := 1; i < len(nums); i++ {
		current_sum = max(nums[i], current_sum+nums[i])
		max_val = max(max_val, current_sum)
	}

	return max_val
}

func main() {
	sol := Solution{}

	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}

	fmt.Println("\nBrute sol : ")
	res1 := sol.brute(nums1)
	res2 := sol.brute(nums2)
	res3 := sol.brute(nums3)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	fmt.Println("\nBetter sol : ")
	res1 = sol.better(nums1)
	res2 = sol.better(nums2)
	res3 = sol.better(nums3)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)

	fmt.Println("\nOptimal sol : ")
	res1 = sol.optimal(nums1)
	res2 = sol.optimal(nums2)
	res3 = sol.optimal(nums3)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
