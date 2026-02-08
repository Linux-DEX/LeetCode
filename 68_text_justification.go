/*
Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.

Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[

	"This    is    an",
	"example  of text",
	"justification.  "

]

Example 2:

Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[

	"What   must   be",
	"acknowledgment  ",
	"shall be        "

]
Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
Note that the second line is also left-justified because it contains only one word.

Example 3:

Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[

	"Science  is  what we",
	"understand      well",
	"enough to explain to",
	"a  computer.  Art is",
	"everything  else  we",
	"do                  "

]
*/
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
