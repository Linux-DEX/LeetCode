package algorithms

import "sort"

// time complexity: O(n + k)
// BucketSort sorts a slice of floats using bucket sort algorithm.
func BucketSort(arr []float64) []float64 {
	n := len(arr)
	if n == 0 {
		return arr
	}

	buckets := make([][]float64, n)
	for i := range arr {
		idx := int(float64(n) * arr[i]) // bucket index
		buckets[idx] = append(buckets[idx], arr[i])
	}

	for i := range buckets {
		sort.Float64s(buckets[i]) // you can also use insertion sort
	}

	idx := 0
	for i := range buckets {
		for _, v := range buckets[i] {
			arr[idx] = v
			idx++
		}
	}

	return arr
}

