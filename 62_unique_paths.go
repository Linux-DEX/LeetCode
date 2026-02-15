/*
# Unique Paths

There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:

Input: m = 3, n = 7
Output: 28

Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/

package main

import (
	"fmt"
)

func dfs(i, j, m, n int) int {
	// out of bounds
	if i >= m || j >= n {
		return 0
	}

	// reached destination
	if i == m-1 && j == n-1 {
		return 1
	}

	// move down + move right
	return dfs(i+1, j, m, n) + dfs(i, j+1, m, n)
}

// Brute Solution : O(2^(m+n)) time , O(m+n) space
func uniquePathsBrute(m int, n int) int {
	return dfs(0, 0, m, n)
}

// Better Solution : O(mxn) time, O(mxn) space
func uniquePathsBetter(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// Optimal Solution : O(min(m, n)) time, O(1) space
func uniquePathsOptimal(m int, n int) int {
	totalMoves := m + n - 2

	r := m - 1
	if n-1 < r {
		r = n - 1
	}

	result := 1
	for i := 1; i <= r; i++ {
		result = result * (totalMoves - r + i) / i
	}

	return result
}

func main() {
	fmt.Println("Brute Solution :")
	fmt.Println(uniquePathsBrute(3, 7))
	fmt.Println(uniquePathsBrute(3, 2))
	fmt.Println()

	fmt.Println("Better Solution :")
	fmt.Println(uniquePathsBetter(3, 7))
	fmt.Println(uniquePathsBetter(3, 2))
	fmt.Println()

	fmt.Println("Optimal Solution :")
	fmt.Println(uniquePathsOptimal(3, 7))
	fmt.Println(uniquePathsOptimal(3, 2))
	fmt.Println()
}
