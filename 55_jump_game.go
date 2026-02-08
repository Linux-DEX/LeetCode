/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/
package main

import (
	"fmt"
)

type Solution struct{}

func (Solution) best(nums []int) bool {
	maxJump := 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if i > maxJump {
			return false
		}

		// update farthest reach
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
	}

	return true
}

func main() {
	sol := Solution{}

	nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}

	fmt.Println("Best sol :", sol.best(nums))
}
