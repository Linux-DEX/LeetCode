/*
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.

According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.

Example 1:

Input: citations = [3,0,6,1,5]

Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:

Input: citations = [1,3,1]
Output: 1
*/
package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

// -----------------------------------
// Brute Force (Sort Descending)
// Time: O(n log n), Space: O(1)
// -----------------------------------
func (s Solution) brute(citations []int) int {
	n := len(citations)

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for h := 0; h < n; h++ {
		if citations[h] < h+1 {
			return h
		}
	}

	return n
}

// -----------------------------------
// Best (Binary Search after sorting)
// Time: O(n log n), Space: O(1)
// -----------------------------------
func (s Solution) best(citations []int) int {
	n := len(citations)

	sort.Ints(citations)

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] >= n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return n - left
}

// -----------------------------------
// Optimal (Counting / Bucket Sort)
// Time: O(n), Space: O(n)
// -----------------------------------
func (s Solution) optimal(citations []int) int {
	n := len(citations)
	freq := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			freq[n]++
		} else {
			freq[c]++
		}
	}

	count := 0
	for h := n; h >= 0; h-- {
		count += freq[h]
		if count >= h {
			return h
		}
	}

	return 0
}

func main() {
	sol := Solution{}

	citations := []int{3, 0, 6, 1, 5}
	fmt.Println("Brute sol :", sol.brute(citations))

	citations = []int{3, 0, 6, 1, 5}
	fmt.Println("Best sol :", sol.best(citations))

	citations = []int{3, 0, 6, 1, 5}
	fmt.Println("Optimal sol :", sol.optimal(citations))
}
