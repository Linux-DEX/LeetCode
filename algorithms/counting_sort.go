package algorithms

// time complexity: O(n + k)
// CountingSort sorts a slice of non-negative integers using the counting sort algorithm.
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find max value
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Count occurrences
	count := make([]int, maxVal+1)
	for _, v := range arr {
		count[v]++
	}

	// Reconstruct sorted array
	idx := 0
	for i, c := range count {
		for j := 0; j < c; j++ {
			arr[idx] = i
			idx++
		}
	}

	return arr
}
