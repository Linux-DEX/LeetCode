/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

# Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

# Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

# Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/
package main

import (
	"fmt"
	"unicode"
)

type Solution struct{}

// Brute force: clean string and compare with its reverse (O(n) time, O(n) space)
func (s Solution) Brute(str string) bool {
	clean := ""
	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			clean += string(unicode.ToLower(c))
		}
	}
	rev := ""
	for i := len(clean) - 1; i >= 0; i-- {
		rev += string(clean[i])
	}
	return clean == rev
}

// Optimal: two-pointer with skipping non-alphanumeric characters (O(n) time, O(1) space)
func (s Solution) Optimal(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(str[left])) && !unicode.IsDigit(rune(str[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(str[right])) && !unicode.IsDigit(rune(str[right])) {
			right--
		}
		if left < right && unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

// Best: same as optimal but with continue statements (cleaner)
func (s Solution) Best(str string) bool {
	for l, r := 0, len(str)-1; l < r; {
		if !unicode.IsLetter(rune(str[l])) && !unicode.IsDigit(rune(str[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(str[r])) && !unicode.IsDigit(rune(str[r])) {
			r--
			continue
		}
		if unicode.ToLower(rune(str[l])) != unicode.ToLower(rune(str[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	sol := Solution{}

	str1 := "A man, a plan, a canal: Panama"
	str2 := "race a car"
	str3 := ""

	fmt.Println("Brute:   ", sol.Brute(str1))
	fmt.Println("Optimal: ", sol.Optimal(str2))
	fmt.Println("Best:    ", sol.Best(str3))
}
