/*
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import "fmt"

type LengthOfLongestSubstring struct{}

// Brute Force (O(n³))
func (l LengthOfLongestSubstring) Brute(s string) int {
	n := len(s)
	maxLen := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if l.allUnique(s, i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}

	return maxLen
}

func (l LengthOfLongestSubstring) allUnique(s string, start, end int) bool {
	seen := make(map[byte]bool)
	for i := start; i <= end; i++ {
		if seen[s[i]] {
			return false
		}
		seen[s[i]] = true
	}
	return true
}

// Better (O(n²))
func (l LengthOfLongestSubstring) Better(s string) int {
	n := len(s)
	maxLen := 0

	for i := 0; i < n; i++ {
		seen := make(map[byte]bool)
		for j := i; j < n; j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}

	return maxLen
}

// (Sliding Window) Optimal (O(n))
func (l LengthOfLongestSubstring) Optimal(s string) int {
	lastIndex := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		if idx, found := lastIndex[s[end]]; found && idx >= start {
			start = idx + 1
		}

		lastIndex[s[end]] = end
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}

	return maxLen
}

func main() {
	l := LengthOfLongestSubstring{}
	str1 := "abcabcbb"
	str2 := "bbbbb"
	str3 := "pwwkew"

	fmt.Println("string: ", str1)
	fmt.Println("Brute:", l.Brute(str1))
	fmt.Println("Better:", l.Better(str1))
	fmt.Println("Optimal:", l.Optimal(str1))
	fmt.Println("\n---\n")

	fmt.Println("string: ", str2)
	fmt.Println("Brute:", l.Brute(str2))
	fmt.Println("Better:", l.Better(str2))
	fmt.Println("Optimal:", l.Optimal(str2))
	fmt.Println("\n---\n")

	fmt.Println("string: ", str3)
	fmt.Println("Brute:", l.Brute(str3))
	fmt.Println("Better:", l.Better(str3))
	fmt.Println("Optimal:", l.Optimal(str3))
}
