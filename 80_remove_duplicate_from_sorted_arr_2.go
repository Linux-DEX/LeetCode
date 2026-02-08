/*
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/
package main

import (
	"fmt"
)

type Solution struct{}

// Brute-force: O(n) time, O(n) space
func (s *Solution) Brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := []int{}

	for _, num := range nums {
		if len(result) < 2 || result[len(result)-2] != num {
			result = append(result, num)
		}
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	return len(result)
}

// Optimal: O(n) time, O(1) space
func (s *Solution) Optimal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1 // write pointer
	count := 1

	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[j-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func main() {
	sol := Solution{}

	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Print("Brute sol : ")
	bruteRes := sol.Brute(nums)
	fmt.Printf("%d => ", bruteRes)
	for _, n := range nums {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

	// reset nums
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	fmt.Print("Optimal sol : ")
	optimalRes := sol.Optimal(nums)
	fmt.Printf("%d => ", optimalRes)
	for _, n := range nums {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()
}
