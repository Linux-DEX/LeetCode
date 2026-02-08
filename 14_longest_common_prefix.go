/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

# Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

# Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
package main

import (
	"fmt"
)

type Solution struct{}

// -------------------------
// Brute Force
// Time Complexity : O(N*M)
// Space Complexity: O(1)
// -------------------------
func (s Solution) Brute(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// -------------------------
// Optimal Approach
// Time Complexity : O(n*L)  (n=strs length, L = length of first string)
// Space Complexity: O(1)
// -------------------------
func (s Solution) Optimal(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	sol := Solution{}

	// input := []string{"dog", "racecar", "car"}
	input := []string{"flower", "flow", "flight"}

	fmt.Println("Brute force  :", sol.Brute(input))
	fmt.Println("Optimal force:", sol.Optimal(input))
}
