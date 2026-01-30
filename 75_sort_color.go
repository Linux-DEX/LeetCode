package main

import (
	"fmt"
)

type Solution struct{}

// Brute-Sol: Merge Sort -> O(n log n) time, O(n) extra space
func (s *Solution) brute(nums []int) {
	if len(nums) <= 1 {
		return
	}

	temp := make([]int, len(nums))
	s.bruteHelper(nums, temp, 0, len(nums)-1)
}

func (s *Solution) bruteHelper(nums, temp []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	s.bruteHelper(nums, temp, left, mid)
	s.bruteHelper(nums, temp, mid+1, right)
	merge(nums, temp, left, mid, right)
}

func merge(nums, temp []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}

	i := left
	j := mid + 1
	k := left

	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
		}
		k++
	}

	for i <= mid {
		nums[k] = temp[i]
		i++
		k++
	}
}

// Better-Sol: O(2 n) time , O(1) space
func (s *Solution) better(nums []int) {
	count0, count1, count2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count0++
		} else if nums[i] == 1 {
			count1++
		} else {
			count2++
		}
	}

	for i := 0; i < count0; i++ {
		nums[i] = 0
	}
	for i := count0; i < count0+count1; i++ {
		nums[i] = 1
	}
	for i := count0 + count1; i < len(nums); i++ {
		nums[i] = 2
	}

}

// Optimal-Sol: O(n) time, O(1) space
func (s *Solution) optimal(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	sol := Solution{}

	nums1 := []int{2, 0, 2, 1, 1, 0}
	nums2 := []int{2, 0, 1}

	sol.brute(nums1)
	sol.brute(nums2)

	fmt.Println("Brute Sol: ")
	fmt.Println(nums1)
	fmt.Println(nums2)

	nums1 = []int{2, 0, 2, 1, 1, 0}
	nums2 = []int{2, 0, 1}

	sol.better(nums1)
	sol.better(nums2)

	fmt.Println("Better Sol: ")
	fmt.Println(nums1)
	fmt.Println(nums2)

	nums1 = []int{2, 0, 2, 1, 1, 0}
	nums2 = []int{2, 0, 1}

	sol.better(nums1)
	sol.better(nums2)

	fmt.Println("Optimal Sol: ")
	fmt.Println(nums1)
	fmt.Println(nums2)
}
