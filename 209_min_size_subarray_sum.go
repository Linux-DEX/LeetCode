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
