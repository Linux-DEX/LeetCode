/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

# Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

# Example 2:
Input: height = [1,1]
Output: 1
*/
package main

import (
	"fmt"
)

type Solution struct{}

// time complexity: O(n)
// space complexity: O(1)
func (s Solution) maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxWater := 0

	for left < right {
		width := right - left

		// min(height[left], height[right])
		h := height[left]
		if height[right] < h {
			h = height[right]
		}

		area := width * h

		// max(maxWater, area)
		if area > maxWater {
			maxWater = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	sol := Solution{}

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println("Max water stored:", sol.maxArea(height))
}
