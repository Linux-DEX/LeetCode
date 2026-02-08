package algorithms

// time complexity: O(n log n)
// HeapSort sorts a slice of integers using the heap sort algorithm.
func HeapSort(arr []int) []int {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

// heapify ensures the subtree rooted at i is a max heap
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
