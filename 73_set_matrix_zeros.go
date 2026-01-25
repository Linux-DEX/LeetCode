package main

import (
	"fmt"
)

type Solution struct{}

func markRow(matrix [][]int, row int) {
	for k := 0; k < len(matrix[row]); k++ {
		if matrix[row][k] != 0 {
			matrix[row][k] = -1
		}
	}
}

func markCol(matrix [][]int, col int) {
	for k := 0; k < len(matrix); k++ {
		if matrix[k][col] != 0 {
			matrix[k][col] = -1
		}
	}
}

// brute-force: O(n^3) time, O(1) space
func (s Solution) brute(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				markRow(matrix, r)
				markCol(matrix, c)
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == -1 {
				matrix[r][c] = 0
			}
		}
	}
}

// better-sol: O(2 x n x m) time, O(n) + O(m) space
func (s Solution) better(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	rowZero := make([]bool, rows)
	colZero := make([]bool, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// Optimal-sol: O(n*m) time, O(1) extra space
func (s Solution) optimal(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	firstRowZero := false
	firstColZero := false

	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

	if firstColZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	sol := Solution{}

	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	fmt.Println("\nbrute Sol : ")
	sol.brute(matrix1)
	fmt.Println(matrix1)

	sol.brute(matrix2)
	fmt.Println(matrix2)

	matrix1 = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix2 = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	fmt.Println("\nbetter Sol : ")
	sol.better(matrix1)
	fmt.Println(matrix1)

	sol.better(matrix2)
	fmt.Println(matrix2)

	matrix1 = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix2 = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	fmt.Println("\nOptimal Sol : ")
	sol.optimal(matrix1)
	fmt.Println(matrix1)

	sol.optimal(matrix2)
	fmt.Println(matrix2)
}
