package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

// -----------------------------------
// Time: O(1)  (max 3999 → constant ops)
// Space: O(1)
// -----------------------------------
func (s Solution) intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV", "I",
	}

	var result strings.Builder

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result.WriteString(symbols[i])
			num -= values[i]
		}
	}

	return result.String()
}

// -----------------------------------
// Main
// -----------------------------------
func main() {
	sol := Solution{}

	fmt.Println(sol.intToRoman(3))    // III
	fmt.Println(sol.intToRoman(58))   // LVIII
	fmt.Println(sol.intToRoman(1994)) // MCMXCIV
}
