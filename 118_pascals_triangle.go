package main

import (
	"fmt"
)

type Solution struct{}

func nCr(n int, r int) int {
	if r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}

	res := 1
	for i := 0; i < r; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

// Brute-Sol: O(nxnxr) time
func (s *Solution) brute(numRows int) [][]int {
	res := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		temp := make([]int, 0, row+1)
		for col := 0; col <= row; col++ {
			temp = append(temp, nCr(row, col))
		}
		res[row] = temp
	}

	return res
}

func generateRow(row int) []int {
	ansRow := []int{1}
	ans := 1
	for i := 1; i <= row; i++ {
		ans = ans * (row - i + 1)
		ans = ans / i
		ansRow = append(ansRow, ans)
	}

	return ansRow
}

// Optimal-Sol:
func (s *Solution) optimal(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = generateRow(i)
	}

	return res
}

func main() {
	sol := Solution{}

	numRows1 := 5
	numRows2 := 1

	fmt.Println("Brute Sol :")
	fmt.Println(sol.brute(numRows1))
	fmt.Println(sol.brute(numRows2))

	fmt.Println("Optimal Sol :")
	fmt.Println(sol.optimal(numRows1))
	fmt.Println(sol.optimal(numRows2))
}
