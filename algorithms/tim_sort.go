package algorithms

const run = 32

// time complexity: O(n log n)
// TimSort sorts a slice of integers using a hybrid of insertion and merge sort.
func TimSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i += run {
		end := i + run
		if end > n {
			end = n
		}
		InsertionSort(arr[i:end])
	}

	size := run
	for size < n {
		for left := 0; left < n; left += 2 * size {
			mid := left + size
			right := left + 2*size
			if mid > n {
				mid = n
			}
			if right > n {
				right = n
			}
			arr[left:right] = merge(arr[left:mid], arr[mid:right])
		}
		size *= 2
	}

	return arr
}

