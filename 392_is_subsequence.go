package main

import (
	"fmt"
)

type Solution struct{}

func (Solution) isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}

func main() {
	solution := Solution{}

	s := "abc"
	t := "ahbgdc"

	if solution.isSubsequence(s, t) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
