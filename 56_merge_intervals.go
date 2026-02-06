package main

import (
	"fmt"
	"sort"
)

// Brute Solution: O(n^2) time
func mergeBrute(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	merged := make([]bool, n)
	result := [][]int{}

	for i := 0; i < n; i++ {
		if merged[i] {
			continue
		}

		start := intervals[i][0]
		end := intervals[i][1]

		for j := i + 1; j < n; j++ {
			if merged[j] {
				continue
			}

			// Check overlap
			if intervals[j][0] <= end && intervals[j][1] >= start {
				start = min(start, intervals[j][0])
				end = max(end, intervals[j][1])
				merged[j] = true
			}
		}

		result = append(result, []int{start, end})
	}

	return result
}

// Optimal Solution: O(n log n) time
func mergeOptimal(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// Sort by start time (if same start, sort by end)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		// overlap
		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			result = append(result, curr)
		}
	}

	return result
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals2 := [][]int{{1, 4}, {4, 5}}
	intervals3 := [][]int{{4, 7}, {1, 4}}

	fmt.Println("Brute solution :")
	fmt.Println(mergeBrute(intervals1))
	fmt.Println(mergeBrute(intervals2))
	fmt.Println(mergeBrute(intervals3))

	fmt.Println("Optimal solution:")
	fmt.Println(mergeOptimal(intervals1))
	fmt.Println(mergeOptimal(intervals2))
	fmt.Println(mergeOptimal(intervals3))
}
