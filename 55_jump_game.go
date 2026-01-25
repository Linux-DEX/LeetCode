package main

import (
	"fmt"
)

type Solution struct{}

func (Solution) best(nums []int) bool {
	maxJump := 0
	size := len(nums)

	for i := 0; i < size; i++ {
		if i > maxJump {
			return false
		}

		// update farthest reach
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
	}

	return true
}

func main() {
	sol := Solution{}

	nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}

	fmt.Println("Best sol :", sol.best(nums))
}

