package main

import (
	"fmt"
)

type Solution struct{}

func (s *Solution) Pprint(matrix [][]int) {
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("%d ", value)
			_ = j
		}
		fmt.Println()
		_ = i
	}
}

// Brute - O(n^2) time, O(n^2) space
func (s *Solution) brute(matrix [][]int) {
	n := len(matrix)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			arr[j][n-1-i] = matrix[i][j]
		}
	}

	// copy arr to matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = arr[i][j]
		}
	}
}

func (s *Solution) transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func (s *Solution) reverse(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

// optimal - O(n^2) time , O(1) space
func (s *Solution) optimal(matrix [][]int) {
	// transpose matrix
	s.transpose(matrix)

	// reverse matrix
	s.reverse(matrix)
}

func main() {
	sol := Solution{}

	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	fmt.Println("Brute solution :")
	sol.brute(matrix1)
	sol.Pprint(matrix1)
	fmt.Println()

	sol.brute(matrix2)
	sol.Pprint(matrix2)
	fmt.Println()

	matrix1 = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	fmt.Println("Optimal Solution :")
	sol.optimal(matrix1)
	sol.Pprint(matrix1)
	fmt.Println()

	sol.optimal(matrix2)
	sol.Pprint(matrix2)
	fmt.Println()
}

