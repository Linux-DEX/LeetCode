/*
Given an unsorted array arr[] of size n, containing elements from the range 1 to n, it is known that one number in this range is missing, and another number occurs twice in the array, find both the duplicate number and the missing number.

Examples:

Input: arr[] = [3, 1, 3]
Output: [3, 2]
Explanation: 3 is occurs twice and 2 is missing.

Input: arr[] = [4, 3, 6, 2, 1, 1]
Output: [1, 5]
Explanation: 1 is occurs twice and 5 is missing.
*/
package main

import (
	"fmt"
	"math"
)

// 1. Map Approach: O(n) Time, O(n) Space
func findTwoElementMap(nums []int) (int, int) {
	repeating, missing := -1, -1
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	for i := 1; i <= len(nums); i++ {
		if freq[i] == 0 {
			missing = i
		} else if freq[i] == 2 {
			repeating = i
		}
	}
	return repeating, missing
}

// 2. Mathematical Approach: O(n) Time, O(1) Space
// Uses Sum of N and Sum of Squares formulas
func findTwoElementMath(nums []int) (int, int) {
	n := int64(len(nums))
	// Sum of first n numbers: (n*(n+1))/2
	sn := (n * (n + 1)) / 2
	// Sum of squares of first n numbers: (n*(n+1)*(2n+1))/6
	s2n := (n * (n + 1) * (2*n + 1)) / 6

	var s, s2 int64
	for _, v := range nums {
		s += int64(v)
		s2 += int64(v) * int64(v)
	}

	// eq1: s - sn = x - y
	// eq2: s2 - s2n = x^2 - y^2
	val1 := s - sn            // x - y
	val2 := (s2 - s2n) / val1 // (x^2 - y^2) / (x - y) = x + y

	repeating := (val1 + val2) / 2
	missing := val2 - repeating

	return int(repeating), int(missing)
}

// 3. XOR Approach: O(n) Time, O(1) Space
// Most robust for interviews (no overflow issues)
func findTwoElementXOR(nums []int) (int, int) {
	resXor := 0
	for i, v := range nums {
		resXor ^= v
		resXor ^= (i + 1)
	}

	// Get the rightmost set bit to create two groups
	setBit := resXor & -resXor
	group0, group1 := 0, 0

	for i := 0; i < len(nums); i++ {
		// Divide array elements
		if nums[i]&setBit != 0 {
			group1 ^= nums[i]
		} else {
			group0 ^= nums[i]
		}
		// Divide numbers from 1 to N
		val := i + 1
		if val&setBit != 0 {
			group1 ^= val
		} else {
			group0 ^= val
		}
	}

	// Verify which one is in the array to distinguish repeating from missing
	for _, v := range nums {
		if v == group0 {
			return group0, group1
		}
	}
	return group1, group0
}

// 4. Index Marking Approach: O(n) Time, O(1) Space
// Modifies the input array to track visits
func findTwoElementIndexing(nums []int) (int, int) {
	// Create a copy so we don't ruin the original array for other tests
	temp := make([]int, len(nums))
	copy(temp, nums)

	repeating, missing := -1, -1

	for i := 0; i < len(temp); i++ {
		val := int(math.Abs(float64(temp[i])))
		index := val - 1
		if temp[index] < 0 {
			repeating = val
		} else {
			temp[index] = -temp[index]
		}
	}

	for i := 0; i < len(temp); i++ {
		if temp[i] > 0 {
			missing = i + 1
			break
		}
	}

	return repeating, missing
}

func main() {
	testCase := []int{4, 3, 6, 2, 1, 1}
	fmt.Printf("Input Array: %v\n", testCase)
	fmt.Println("------------------------------------")

	r1, m1 := findTwoElementMap(testCase)
	fmt.Printf("1. Map Method:       Repeating: %d, Missing: %d\n", r1, m1)

	r2, m2 := findTwoElementMath(testCase)
	fmt.Printf("2. Math Method:      Repeating: %d, Missing: %d\n", r2, m2)

	r3, m3 := findTwoElementXOR(testCase)
	fmt.Printf("3. XOR Method:       Repeating: %d, Missing: %d\n", r3, m3)

	r4, m4 := findTwoElementIndexing(testCase)
	fmt.Printf("4. Indexing Method:  Repeating: %d, Missing: %d\n", r4, m4)

	testCase = []int{3, 1, 3}
	fmt.Printf("Input Array: %v\n", testCase)
	fmt.Println("------------------------------------")

	r1, m1 = findTwoElementMap(testCase)
	fmt.Printf("1. Map Method:       Repeating: %d, Missing: %d\n", r1, m1)

	r2, m2 = findTwoElementMath(testCase)
	fmt.Printf("2. Math Method:      Repeating: %d, Missing: %d\n", r2, m2)

	r3, m3 = findTwoElementXOR(testCase)
	fmt.Printf("3. XOR Method:       Repeating: %d, Missing: %d\n", r3, m3)

	r4, m4 = findTwoElementIndexing(testCase)
	fmt.Printf("4. Indexing Method:  Repeating: %d, Missing: %d\n", r4, m4)
}
