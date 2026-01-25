package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

// ThreeSum returns all unique triplets that sum to zero
func (s Solution) ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate values
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	sol := Solution{}
	nums := []int{-1, 0, 1, 2, -1, -4}

	result := sol.ThreeSum(nums)

	fmt.Println("Unique triplets:")
	for _, triplet := range result {
		fmt.Printf("[%d, %d, %d]\n", triplet[0], triplet[1], triplet[2])
	}
}
