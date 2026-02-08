/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

# Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

# Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

# Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/
package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

// ThreeSum returns all unique triplets that sum to zero
func (s Solution) ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate values
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	sol := Solution{}
	nums := []int{-1, 0, 1, 2, -1, -4}

	result := sol.ThreeSum(nums)

	fmt.Println("Unique triplets:")
	for _, triplet := range result {
		fmt.Printf("[%d, %d, %d]\n", triplet[0], triplet[1], triplet[2])
	}
}
