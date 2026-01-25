package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

// -------------------------
// Brute Force: Using Split
// Time  : O(n)
// Space : O(n)
// -------------------------
func (s Solution) Brute(str string) int {
	words := strings.Fields(str) // splits by spaces automatically
	if len(words) == 0 {
		return 0
	}
	lastWord := words[len(words)-1]
	return len(lastWord)
}

// -------------------------
// Optimal: O(1) space
// Time  : O(n)
// Space : O(1)
// -------------------------
func (s Solution) Best(str string) int {
	i := len(str) - 1

	// skip trailing spaces
	for i >= 0 && str[i] == ' ' {
		i--
	}

	length := 0
	// count characters until space or start
	for i >= 0 && str[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	sol := Solution{}

	// Test cases
	input := "luffy is still joyboy"
	fmt.Println("Brute Sol :", sol.Brute(input))
	fmt.Println("Best Sol  :", sol.Best(input))

	// Example outputs
	// input = "   Hello World   "
	// input = "   fly me   to   the moon  "
}

