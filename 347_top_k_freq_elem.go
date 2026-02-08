/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]

Example 3:

Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
Output: [1,2]
*/
package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements a min-heap for pairs [frequency, number]
type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // compare by frequency
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Solution struct{}

// Best approach: bucket sort O(n)
func (s Solution) Best(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Create buckets where index = frequency
	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}

	return result
}

// Optimal approach: min-heap O(n log k)
func (s Solution) Optimal(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, [2]int{freq, num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := []int{}
	for h.Len() > 0 {
		top := heap.Pop(h).([2]int)
		result = append(result, top[1])
	}

	return result
}

func main() {
	sol := Solution{}

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Print("Best sol    : ")
	ans := sol.Best(nums, k)
	for _, n := range ans {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

	fmt.Print("Optimal sol : ")
	optAns := sol.Optimal(nums, k)
	for _, n := range optAns {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
