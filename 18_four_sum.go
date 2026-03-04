/*
# 4Sum

Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
*/

package main

import (
	"fmt"
	"sort"
)

// Brute Solution : O(n^4) time, O(1) space
func fourSumBrute(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for k := j + 1; k < n-1; k++ {
				for l := k + 1; l < n; l++ {
					sum := nums[i] + nums[j] + nums[k] + nums[l]
					if sum == target {
						quad := []int{nums[i], nums[j], nums[k], nums[l]}
						result = append(result, quad)
					}
				}
			}
		}
	}

	return result
}

// Better Solution : O(n^3) time, O(1) space
func fourSumBetter(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {

			seen := make(map[int]bool)

			for k := j + 1; k < n; k++ {
				complement := target - nums[i] - nums[j] - nums[k]

				if seen[complement] {
					quad := []int{nums[i], nums[j], complement, nums[k]}
					sort.Ints(quad)
					result = append(result, quad)
				}
				seen[nums[k]] = true
			}
		}
	}

	return result
}

// Optimal Solution : O(n^3) time, O(1) space
func fourSumOptimal(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{
						nums[i], nums[j], nums[left], nums[right],
					})

					left++
					right--

					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Brute Solution : ")
	fmt.Println(fourSumBrute([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSumBrute([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println()

	fmt.Println("Better Solution : ")
	fmt.Println(fourSumBetter([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSumBetter([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println()

	fmt.Println("Optimal Solution : ")
	fmt.Println(fourSumOptimal([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSumOptimal([]int{2, 2, 2, 2, 2}, 8))
}
