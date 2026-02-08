package algorithms

// time complexity: O(n log n) on average
// QuickSort sorts a slice of integers using the quick sort algorithm.
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, equal...), right...)
}
