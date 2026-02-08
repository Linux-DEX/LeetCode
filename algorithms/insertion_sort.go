package algorithms

// time complexity: O(n^2)
// InsertionSort sorts a slice of integers using the insertion sort algorithm.
func InsertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1] that are greater than key
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Place key at its correct position
		arr[j+1] = key
	}

	return arr
}
