package main

import (
	"fmt"
)

type Solution struct{}

// Brute approach: check each subtraction case explicitly
func (s Solution) Brute(roman string) int {
	result := 0
	n := len(roman)

	for i := 0; i < n; i++ {
		if i+1 < n && roman[i] == 'I' && (roman[i+1] == 'V' || roman[i+1] == 'X') {
			result -= 1
		} else if i+1 < n && roman[i] == 'X' && (roman[i+1] == 'L' || roman[i+1] == 'C') {
			result -= 10
		} else if i+1 < n && roman[i] == 'C' && (roman[i+1] == 'D' || roman[i+1] == 'M') {
			result -= 100
		} else {
			switch roman[i] {
			case 'I':
				result += 1
			case 'V':
				result += 5
			case 'X':
				result += 10
			case 'L':
				result += 50
			case 'C':
				result += 100
			case 'D':
				result += 500
			case 'M':
				result += 1000
			}
		}
	}
	return result
}

// Best approach: combine conditions
func (s Solution) Best(roman string) int {
	result := 0
	n := len(roman)

	for i := 0; i < n; i++ {
		if i+1 < n && ((roman[i] == 'I' && (roman[i+1] == 'V' || roman[i+1] == 'X')) ||
			(roman[i] == 'X' && (roman[i+1] == 'L' || roman[i+1] == 'C')) ||
			(roman[i] == 'C' && (roman[i+1] == 'D' || roman[i+1] == 'M'))) {
			if roman[i] == 'I' {
				result -= 1
			} else if roman[i] == 'X' {
				result -= 10
			} else {
				result -= 100
			}
		} else {
			switch roman[i] {
			case 'I':
				result += 1
			case 'V':
				result += 5
			case 'X':
				result += 10
			case 'L':
				result += 50
			case 'C':
				result += 100
			case 'D':
				result += 500
			case 'M':
				result += 1000
			}
		}
	}

	return result
}

// Optimal approach: use a map for direct lookup
func (s Solution) Optimal(roman string) int {
	romanMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	n := len(roman)
	for i := 0; i < n; i++ {
		if i+1 < n && romanMap[roman[i]] < romanMap[roman[i+1]] {
			result -= romanMap[roman[i]]
		} else {
			result += romanMap[roman[i]]
		}
	}

	return result
}

func main() {
	sol := Solution{}

	// Test cases
	// roman := "III"
	roman := "LVIII"
	// roman := "MCMXCIV"

	fmt.Println("Brute sol   :", sol.Brute(roman))
	fmt.Println("Best sol    :", sol.Best(roman))
	fmt.Println("Optimal sol :", sol.Optimal(roman))
}
