package algorithms

import "math"

// time complexity: O(n log n)
// IntroSort sorts a slice of integers using quick sort and switches to heap sort if recursion is too deep.
func IntroSort(arr []int) []int {
	maxDepth := 2 * int(math.Log2(float64(len(arr))))
	introSortHelper(arr, maxDepth)
	return arr
}

func introSortHelper(arr []int, depth int) {
	if len(arr) <= 16 {
		InsertionSort(arr)
		return
	}

	if depth == 0 {
		HeapSort(arr)
		return
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	equal := []int{}
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}

	introSortHelper(left, depth-1)
	introSortHelper(right, depth-1)

	copy(arr, append(append(left, equal...), right...))
}
