package main

import "fmt"

type Solution struct{}

/*
Time: O(n^2)
Space: O(n)
*/
func (s Solution) brute(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	for i := 0; i < size; i++ {
		prod := 1
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			prod *= nums[j]
		}
		result[i] = prod
	}

	return result
}

/*
Time: O(3n) ~ O(n)
Space: O(n)
*/
func (s Solution) best(nums []int) []int {
	size := len(nums)

	prefix := make([]int, size)
	postfix := make([]int, size)
	result := make([]int, size)

	prefix[0] = 1
	postfix[size-1] = 1

	for i := 1; i < size; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := size - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}

	for i := 0; i < size; i++ {
		result[i] = prefix[i] * postfix[i]
	}

	return result
}

/*
Time: O(2n) ~ O(n)
Space: O(1) extra (output excluded)
*/
func (s Solution) optimal(nums []int) []int {
	size := len(nums)
	result := make([]int, size)

	prefix := 1
	for i := 0; i < size; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := size - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
}

func main() {
	sol := Solution{}
	arr := []int{1, 2, 3, 4}

	fmt.Print("Brute sol : ")
	printArr(sol.brute(arr))

	fmt.Print("Best sol : ")
	printArr(sol.best(arr))

	fmt.Print("Optimal sol : ")
	printArr(sol.optimal(arr))
}

