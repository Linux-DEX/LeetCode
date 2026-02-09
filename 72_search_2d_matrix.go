/*
74. Search a 2D Matrix

You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

package main

import (
	"fmt"
	"linux-dex/leetcode/algorithms"
)

// Brute Solution : O(n x m) time , O(1) space
func searchMatrixBrute(matrix [][]int, target int) bool {
	// linear traversal
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if matrix[row][col] == target {
				return true
			}
		}
	}
	return false
}

// Better Solution : O(n log n) time , O(1) space
func searchMatrixBetter(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		if target < row[0] || target > row[len(row)-1] {
			continue
		}

		if algorithms.BinarySearch(row, target) {
			return true
		}
	}
	return false
}

// Better Solution : O(n + m) time, O(1) space
func searchMatrixBetter2(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}

	return false
}

// Optimal Solution : O(n + m) time, O(1) space
func searchMatrixOptimal(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2

		row := mid / n
		col := mid % n
		value := matrix[row][col]

		if value == target {
			return true
		} else if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	matrix1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target1 := 3
	matrix2 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target2 := 13

	fmt.Println("Brute Solution :")
	fmt.Println(searchMatrixBrute(matrix1, target1))
	fmt.Println(searchMatrixBrute(matrix2, target2))

	fmt.Println("Better Solution :")
	fmt.Println(searchMatrixBetter(matrix1, target1))
	fmt.Println(searchMatrixBetter(matrix2, target2))

	fmt.Println("Better2 Solution :")
	fmt.Println(searchMatrixBetter2(matrix1, target1))
	fmt.Println(searchMatrixBetter2(matrix2, target2))

	fmt.Println("Optimal Solution :")
	fmt.Println(searchMatrixOptimal(matrix1, target1))
	fmt.Println(searchMatrixOptimal(matrix2, target2))
}
