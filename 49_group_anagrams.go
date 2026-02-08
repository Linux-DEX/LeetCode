/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explanation:
There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Example 2:

Input: strs = [""]
Output: [[""]]

Example 3:

Input: strs = ["a"]
Output: [["a"]]
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

// Best (Sorting based) - O(n * k log k)
func (s Solution) best(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")

		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	result := [][]string{}
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

// Optimal (Frequency count) - O(n * k)
func (s Solution) optimal(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	ansMap := make(map[string][]string)

	for _, str := range strs {
		count := make([]int, 26)

		for _, c := range str {
			count[c-'a']++
		}

		var keyBuilder strings.Builder
		for i := 0; i < 26; i++ {
			keyBuilder.WriteString("#")
			keyBuilder.WriteString(strconv.Itoa(count[i]))
		}

		key := keyBuilder.String()
		ansMap[key] = append(ansMap[key], str)
	}

	result := [][]string{}
	for _, group := range ansMap {
		result = append(result, group)
	}

	return result
}

// Print helper
func (s Solution) printArr(result [][]string) {
	for _, group := range result {
		for _, word := range group {
			fmt.Print(word, ", ")
		}
		fmt.Println()
	}
}

func main() {
	sol := Solution{}

	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println("Best solution :")
	resultBest := sol.best(arr)
	sol.printArr(resultBest)

	fmt.Println("Optimal sol :")
	resultOpt := sol.optimal(arr)
	sol.printArr(resultOpt)
}
