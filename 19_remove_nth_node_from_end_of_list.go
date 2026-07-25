/*
19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

---

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]
*/

package main

import (
	"fmt"
	"linux-dex/leetcode/datastructure"
)

type Solution struct{}

// Brute Solution: Time O(n), Space O(n)
func (s *Solution) BruteForce(head *datastructure.ListNode[int], n int) *datastructure.ListNode[int] {
    if head == nil {
		return nil
	}
    
    dummy := &datastructure.ListNode[int]{Next: head}
	cur := dummy
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	prev := dummy
	for cur.Next != nil {
		prev = prev.Next
		cur = cur.Next
	}

	prev.Next = prev.Next.Next
	return dummy.Next
}

// Better Solution: Time O(n) [2 passes], Space O(1)
func (s *Solution) BetterSolution(head *datastructure.ListNode[int], n int) *datastructure.ListNode[int] {
	if head == nil {
		return nil
	}

	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	dummy := &datastructure.ListNode[int]{Next: head}
	cur = dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	
	return dummy.Next
}

// Optimal Solution(fast and slow pointers): Time O(n), Space O(1)
func (s *Solution) OptimalSolution(head *datastructure.ListNode[int], n int) *datastructure.ListNode[int] {
	dummy := &datastructure.ListNode[int]{Next: head}
	fast := dummy
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return dummy.Next
}

func main() {
	sol := Solution{}

	testCase := []struct {
		list []int
		n    int
	}{
		{list: []int{1, 2, 3, 4, 5}, n: 2},
		{list: []int{1, 2}, n: 1},
		{list: []int{1}, n: 1},
	}

	for i, tc := range testCase {
		fmt.Printf("\n===== Test Case %d ======\n", i)

		fmt.Printf("N : %d", tc.n)
		fmt.Print(" Input List: ")
		datastructure.PrintList(datastructure.CreateList(tc.list))

		fmt.Print("Brute Solution: ")
		datastructure.PrintList(sol.BruteForce(datastructure.CreateList(tc.list), tc.n))

		fmt.Print("Better Solution: ")
		datastructure.PrintList(sol.BetterSolution(datastructure.CreateList(tc.list), tc.n))

		fmt.Print("Optimal Solution: ")
		datastructure.PrintList(sol.OptimalSolution(datastructure.CreateList(tc.list), tc.n))
	}
}
