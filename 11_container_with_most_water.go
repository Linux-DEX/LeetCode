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
