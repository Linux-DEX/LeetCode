package algorithms

// time complexity: O(d*(n + k))
// RadixSort sorts a slice of non-negative integers using radix sort (LSD)
func RadixSort(arr []int) []int {
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

	exp := 1
	for maxVal/exp > 0 {
		countingSortByDigit(arr, exp)
		exp *= 10
	}

	return arr
}

// countingSortByDigit sorts arr based on the digit represented by exp
func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	copy(arr, output)
}

