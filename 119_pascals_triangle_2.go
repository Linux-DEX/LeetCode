package main

import (
	"fmt"
)

type Solution struct{}

func ncr(n int, r int) int {
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

// Brute-sol: O(n x r) time , O(r) space
func (s Solution) brute(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for c := 0; c <= rowIndex; c++ {
		res[c] = ncr(rowIndex, c)
	}

	return res
}

// Optimal-sol: O(n) time , O(1) space
func (s Solution) optimal(rowIndex int) []int {
	res := []int{1}
	ans := 1
	for i := 1; i <= rowIndex; i++ {
		ans = ans * (rowIndex - i + 1)
		ans = ans / i
		res = append(res, ans)
	}

	return res
}

func main() {
	sol := Solution{}

	rowIndex1 := 3
	rowIndex2 := 0
	rowIndex3 := 1

	fmt.Println("Brute Sol :")
	fmt.Println("row", rowIndex1, " : ", sol.brute(rowIndex1))
	fmt.Println("row", rowIndex2, " : ", sol.brute(rowIndex2))
	fmt.Println("row", rowIndex3, " : ", sol.brute(rowIndex3))

	fmt.Println("Optimal Sol :")
	fmt.Println("row", rowIndex1, " : ", sol.optimal(rowIndex1))
	fmt.Println("row", rowIndex2, " : ", sol.optimal(rowIndex2))
	fmt.Println("row", rowIndex3, " : ", sol.optimal(rowIndex3))
}
