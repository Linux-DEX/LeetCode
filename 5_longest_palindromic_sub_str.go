/*
# Longest Palindromic substring

Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"
*/

package main

import "fmt"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}

// Brute Solution : O(n^3) time, O(1) space
func longestPalindromeBrute(s string) string {
	n := len(s)
	longest := ""

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sub := s[i : j+1]
			if isPalindrome(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}

	return longest
}

func expandFromCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

// Better Solution : O(n^2) time, O(1) space
func longestPalindromeBetter(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Odd length palindrome
		l1, r1 := expandFromCenter(s, i, i)
		// Even length palindrome
		l2, r2 := expandFromCenter(s, i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// Optimal Solution (Manacher's Algorithm) : O(n) time , O(n) space
func longestPalindromeManacher(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Transform string: "^#a#b#a#d#$"
	t := "^"
	for _, ch := range s {
		t += "#" + string(ch)
	}
	t += "#$"

	n := len(t)
	p := make([]int, n)
	center, right := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*center - i

		if i < right {
			if p[mirror] < right-i {
				p[i] = p[mirror]
			} else {
				p[i] = right - i
			}
		}

		// Expand around center
		for t[i+(1+p[i])] == t[i-(1+p[i])] {
			p[i]++
		}

		// Update center and right
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	// Find max palindrome
	maxLen, centerIndex := 0, 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}


func main() {
	fmt.Println("Brute Solution : ")
	fmt.Println(longestPalindromeBrute("babad"))
	fmt.Println(longestPalindromeBrute("cbbd"))

	fmt.Println("Better Solution : ")
	fmt.Println(longestPalindromeBetter("babad"))
	fmt.Println(longestPalindromeBetter("cbbd"))

	fmt.Println("Optimal Solution : ")
	fmt.Println(longestPalindromeManacher("babad"))
	fmt.Println(longestPalindromeManacher("cbbd"))
}

