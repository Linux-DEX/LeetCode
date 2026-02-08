/*
# Count Inversions of an Array

Given an array of integers arr[]. You have to find the Inversion Count of the array.
Note : Inversion count is the number of pairs of elements (i, j) such that i < j and arr[i] > arr[j].

Examples:

Input: arr[] = [2, 4, 1, 3, 5]
Output: 3
Explanation: The sequence 2, 4, 1, 3, 5 has three inversions (2, 1), (4, 1), (4, 3).

Input: arr[] = [2, 3, 4, 5, 6]
Output: 0
Explanation: As the sequence is already sorted so there is no inversion count.

Input: arr[] = [10, 10, 10]
Output: 0
Explanation: As all the elements of array are same, so there is no inversion count.
*/

package main

import (
	"fmt"
)

// Brute Solution : O(n^2) time , O(1) space
func inversionCountBrute(arr []int) int {
	count := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				count += 1
			}
		}
	}

	return count
}

func countAndMerge(arr []int, l, m, r int) int {
	// Create temp slices
	left := make([]int, m-l+1)
	right := make([]int, r-m)

	copy(left, arr[l:m+1])
	copy(right, arr[m+1:r+1])

	i, j, k := 0, 0, l
	invCount := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			// Count inversions
			invCount += len(left) - i
		}
		k++
	}

	// Copy remaining elements
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}

	return invCount
}

func countInv(arr []int, l, r int) int {
	if l >= r {
		return 0
	}

	m := (l + r) / 2
	invCount := 0

	// Left inversions
	invCount += countInv(arr, l, m)
	// Right inversions
	invCount += countInv(arr, m+1, r)
	// Cross inversions
	invCount += countAndMerge(arr, l, m, r)

	return invCount
}

// Optimal Solution : O(n log n)
func inversionCountOptimal(arr []int) int {
	return countInv(arr, 0, len(arr)-1)
}

func main() {
	arr1 := []int{2, 4, 1, 3, 5}
	arr2 := []int{2, 3, 4, 5, 6}
	arr3 := []int{10, 10, 10}

	fmt.Println("Brute Solution :")
	fmt.Println(inversionCountBrute(arr1))
	fmt.Println(inversionCountBrute(arr2))
	fmt.Println(inversionCountBrute(arr3))

	fmt.Println("Optimal Solution :")
	fmt.Println(inversionCountOptimal(arr1))
	fmt.Println(inversionCountOptimal(arr2))
	fmt.Println(inversionCountOptimal(arr3))
}
