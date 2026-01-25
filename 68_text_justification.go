package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

// Brute approach: greedy line building with space distribution
func (s Solution) Brute(words []string, maxWidth int) []string {
	var res []string
	n := len(words)
	i := 0

	for i < n {
		j := i + 1
		lineLen := len(words[i])

		// Determine how many words fit in the current line
		for j < n && lineLen+len(words[j])+(j-i) <= maxWidth {
			lineLen += len(words[j])
			j++
		}

		spaceSlots := j - i - 1
		line := words[i]

		// Last line or single-word line -> left justified
		if j == n || spaceSlots == 0 {
			for k := i + 1; k < j; k++ {
				line += " " + words[k]
			}
			line += strings.Repeat(" ", maxWidth-len(line))
		} else {
			// Fully justify the line
			totalSpaces := maxWidth - lineLen
			space := totalSpaces / spaceSlots
			extra := totalSpaces % spaceSlots

			for k := i + 1; k < j; k++ {
				spaces := space
				if extra > 0 {
					spaces++
					extra--
				}
				line += strings.Repeat(" ", spaces) + words[k]
			}
		}

		res = append(res, line)
		i = j
	}

	return res
}

func main() {
	sol := Solution{}

	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16

	result := sol.Brute(words, maxWidth)
	for _, line := range result {
		fmt.Printf("\"%s\"\n", line)
	}
}
