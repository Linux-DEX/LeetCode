/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at index 0.

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at index i, you can jump to any index (i + j) where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach index n - 1. The test cases are generated such that you can reach index n - 1.

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [2,3,0,1,4]
Output: 2
*/
package main

import (
	"fmt"
	"math"
)

type Solution struct{}

/*
   -------------------------
   Brute Force (DFS + Memo)
   Time  : O(n^2)
   Space : O(n)
   -------------------------
*/

func (s Solution) minJumpsHelper(nums []int, currentPos int, memo []int) int {
	n := len(nums)

	// reached or crossed last index
	if currentPos >= n-1 {
		return 0
	}

	if memo[currentPos] != -1 {
		return memo[currentPos]
	}

	jumps := math.MaxInt32

	for i := 1; i <= nums[currentPos]; i++ {
		if currentPos+i < n {
			result := s.minJumpsHelper(nums, currentPos+i, memo)
			if result != math.MaxInt32 {
				jumps = min(jumps, 1+result)
			}
		}
	}

	memo[currentPos] = jumps
	return jumps
}

func (s Solution) Brute(nums []int) int {
	n := len(nums)
	memo := make([]int, n)

	for i := range memo {
		memo[i] = -1
	}

	return s.minJumpsHelper(nums, 0, memo)
}

/*
   -------------------------
   Optimal Greedy
   Time  : O(n)
   Space : O(1)
   -------------------------
*/

func (s Solution) Optimal(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == currentEnd {
			jumps++
			currentEnd = farthest

			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

/*
   Utility functions
*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
   Main
*/

func main() {
	sol := Solution{}

	nums := []int{2, 3, 1, 1, 4}

	fmt.Println("Brute sol :", sol.Brute(nums))
	fmt.Println("Optimal sol :", sol.Optimal(nums))
}
