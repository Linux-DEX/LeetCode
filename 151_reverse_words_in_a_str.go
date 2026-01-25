package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

// Brute approach: split, reverse, join
func (s Solution) Brute(input string) string {
	words := strings.Fields(input) // splits by whitespace and removes extra spaces
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// Best approach: use stack
func (s Solution) Best(input string) string {
	words := strings.Fields(input)
	stack := make([]string, 0, len(words))
	for _, word := range words {
		stack = append(stack, word)
	}

	result := ""
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(result) > 0 {
			result += " "
		}
		result += top
	}
	return result
}

// Optimal approach: trim spaces, reverse string, reverse each word
func (s Solution) Optimal(input string) string {
	// Trim leading and trailing spaces, reduce multiple spaces
	trimmed := strings.Join(strings.Fields(input), " ")

	// Reverse entire string
	runes := []rune(trimmed)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Reverse each word
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			for l, r := start, i-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	sol := Solution{}

	// Test input
	input := "a good   example"

	fmt.Println("Brute sol   :", sol.Brute(input))
	fmt.Println("Best sol    :", sol.Best(input))
	fmt.Println("Optimal sol :", sol.Optimal(input))
}
