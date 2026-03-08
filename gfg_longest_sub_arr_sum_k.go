/*
# Longest Subarray with Sum K

Given an array arr[] containing integers and an integer k, your task is to find the length of the longest subarray where the sum of its elements is equal to the given value k. If there is no subarray with sum equal to k, return 0.

Examples:

Input: arr[] = [10, 5, 2, 7, 1, -10], k = 15
Output: 6
Explanation: Subarrays with sum = 15 are [5, 2, 7, 1], [10, 5] and [10, 5, 2, 7, 1, -10]. The length of the longest subarray with a sum of 15 is 6.

Input: arr[] = [-5, 8, -14, 2, 4, 12], k = -5
Output: 5
Explanation: Subarrays with sum = -5 are [-5] and [-5, 8, -14, 2, 4]. The length of the longest subarray with a sum of -5 is 5.

Input: arr[] = [10, -10, 20, 30], k = 5
Output: 0
Explanation: No subarray with sum = 5 is present in arr[].
*/

package main

import (
	"fmt"
)

// Brute Solution : O(n^2) time, O(1) space
func longestSubarrayBrute(arr []int, k int) int {
	n := len(arr)
	maxLen := 0

	for i := 0; i < n; i++ {
		sum := 0

		for j := i; j < n; j++ {
			sum += arr[j]

			if sum == k {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}

	return maxLen
}

// Better Solution : O(n) time, O(n) space
func longestSubarrayOptimal(arr []int, k int) int {
	preSumMap := make(map[int]int)
	sum := 0
	maxLen := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum == k {
			maxLen = i + 1
		}

		rem := sum - k
		if firstIndex, ok := preSumMap[rem]; ok {
			length := i - firstIndex
			if length > maxLen {
				maxLen = length
			}
		}

		if _, ok := preSumMap[sum]; !ok {
			preSumMap[sum] = i
		}
	}

	return maxLen
}

func main() {
	fmt.Println("Brute Solution : ")
	fmt.Println(longestSubarrayBrute([]int{10, 5, 2, 7, 1, -10}, 15))
	fmt.Println(longestSubarrayBrute([]int{-5, 8, -14, 2, 4, 12}, -5))
	fmt.Println(longestSubarrayBrute([]int{10, -10, 20, 30}, 5))
	fmt.Println()

	fmt.Println("Optimal Solution : ")
	fmt.Println(longestSubarrayOptimal([]int{10, 5, 2, 7, 1, -10}, 15))
	fmt.Println(longestSubarrayOptimal([]int{-5, 8, -14, 2, 4, 12}, -5))
	fmt.Println(longestSubarrayOptimal([]int{10, -10, 20, 30}, 5))
	fmt.Println()
}
