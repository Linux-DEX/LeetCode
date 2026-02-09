/*
50. Pow(x, n)

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000

Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100

Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
*/

package main

import "fmt"

// Brute Solution : O(n) time, O(1) space
func myPowBrute(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	result := 1.0
	for i := 0; i < n; i++ {
		result *= x
	}

	if isNegative {
		return 1 / result
	}
	return result
}

// Better Solution : O(log n) time, O(log n) space
func myPowBetter(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPowBetter(x, -n)
	}

	half := myPowBetter(x, n/2)

	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// Optimal Solution : O(log n) time, O(1) space
func myPowOptimal(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0

	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}

	return result
}

func main() {
	x1 := 2.00000
	n1 := 10
	x2 := 2.10000
	n2 := 3
	x3 := 2.00000
	n3 := -2

	fmt.Println("Brute Solution :")
	fmt.Println(myPowBrute(x1, n1))
	fmt.Println(myPowBrute(x2, n2))
	fmt.Println(myPowBrute(x3, n3))

	fmt.Println("Better Solution :")
	fmt.Println(myPowBetter(x1, n1))
	fmt.Println(myPowBetter(x2, n2))
	fmt.Println(myPowBetter(x3, n3))

	fmt.Println("Optimal Solution :")
	fmt.Println(myPowOptimal(x1, n1))
	fmt.Println(myPowOptimal(x2, n2))
	fmt.Println(myPowOptimal(x3, n3))
}
