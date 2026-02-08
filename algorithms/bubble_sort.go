package algorithms

// time complexity: O(n^2)
// BubbleSort sorts a slice of integers using the bubble sort algorithm.
func BubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// After each pass, the largest element bubbles to the end
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
