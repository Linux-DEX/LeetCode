package algorithms

// time complexity: O(n^2)
// SelectionSort sorts a slice of integers using the selection sort algorithm.
func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		mini := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}
		arr[mini], arr[i] = arr[i], arr[mini]
	}

	return arr
}
