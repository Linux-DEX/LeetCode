/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false
*/
package main

import (
	"fmt"
)

type Solution struct{}

// ---------------- Brute Force ----------------
func (Solution) brute(s string, slen int, t string, tlen int) bool {
	if slen != tlen {
		return false
	}

	used := make([]bool, tlen)

	for i := 0; i < slen; i++ {
		found := false

		for j := 0; j < tlen; j++ {
			if !used[j] && s[i] == t[j] {
				used[j] = true
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// ---------------- Optimal ----------------
// Time: O(n)
// Space: O(1)
func (Solution) optimal(s string, slen int, t string, tlen int) bool {
	if slen != tlen {
		return false
	}

	var arr [26]int

	for i := 0; i < slen; i++ {
		arr[s[i]-'a']--
		arr[t[i]-'a']++
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

// ---------------- Main ----------------
func main() {
	sol := Solution{}

	s := "cat"
	t := "tax"

	slen := len(s)
	tlen := len(t)

	fmt.Print("Brute force : ")
	if sol.brute(s, slen, t, tlen) {
		fmt.Println("yes anagram")
	} else {
		fmt.Println("not anagram")
	}

	fmt.Print("Optimal solution : ")
	if sol.optimal(s, slen, t, tlen) {
		fmt.Println("yes anagram")
	} else {
		fmt.Println("not anagram")
	}
}
