/*
# Majority Element 2

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:

Input: nums = [3,2,3]
Output: [3]

Example 2:

Input: nums = [1]
Output: [1]

Example 3:

Input: nums = [1,2]
Output: [1,2]
*/

package main

import "fmt"

// Brute Solution : O(n^2) time, O(1) space
func majorityElementBrute(nums []int) []int {
	var result []int
	n := len(nums)

	for i := 0; i < n; i++ {
		alreadyAdded := false
		for _, v := range result {
			if v == nums[i] {
				alreadyAdded = true
				break
			}
		}
		if alreadyAdded {
			continue
		}

		count := 0
		for j := 0; j < n; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}

		if count > n/3 {
			result = append(result, nums[i])
		}
	}

	return result
}

// Better Solution : O(n) time, O(n) space
func majorityElementBetter(nums []int) []int {
	freq := make(map[int]int)
	n := len(nums)
	var result []int

	for _, nums := range nums {
		freq[nums]++
	}

	for num, count := range freq {
		if count > n/3 {
			result = append(result, num)
		}
	}

	return result
}

// Optimal Solution : O(n) time, O(1) space
func majorityElementOptimal(nums []int) []int {
	var candidate1, candidate2 int
	count1, count2 := 0, 0

	// First pass: Find potential candidates
	for _, num := range nums {
		if count1 > 0 && num == candidate1 {
			count1++
		} else if count2 > 0 && num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	// Second pass: Verify candidates
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	var result []int
	n := len(nums)

	if count1 > n/3 {
		result = append(result, candidate1)
	}
	if count2 > n/3 {
		result = append(result, candidate2)
	}

	return result
}

func main() {
	nums1 := []int{3, 2, 3}
	nums2 := []int{1}
	nums3 := []int{1, 2}

	fmt.Println("Brute Solution : ")
	fmt.Println(majorityElementBrute(nums1))
	fmt.Println(majorityElementBrute(nums2))
	fmt.Println(majorityElementBrute(nums3))

	fmt.Println("Better Solution : ")
	fmt.Println(majorityElementBetter(nums1))
	fmt.Println(majorityElementBetter(nums2))
	fmt.Println(majorityElementBetter(nums3))

	fmt.Println("Optimal Solution : ")
	fmt.Println(majorityElementOptimal(nums1))
	fmt.Println(majorityElementOptimal(nums2))
	fmt.Println(majorityElementOptimal(nums3))
}

