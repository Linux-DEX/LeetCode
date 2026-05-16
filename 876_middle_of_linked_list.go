/*
876. Middle of the Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

---

Example 1:

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/
package main

import (
	"fmt"

	"linux-dex/leetcode/datastructure"
)

// Alias the generic node
type ListNode = datastructure.ListNode[int]

type Solution struct{}

// Brute-force: O(n) time, O(n) space
func (s *Solution) Brute(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodes := []*ListNode{}
	for temp := head; temp != nil; temp = temp.Next {
		nodes = append(nodes, temp)
	}
	return nodes[len(nodes)/2]
}

// Better: O(n) time, O(1) space (Two Pass)
func (s *Solution) Better(head *ListNode) *ListNode {
	length := datastructure.ListLength(head)
	return datastructure.GetNode(head, length/2)
}

// Optimal: O(n) time, O(1) space (Slow & Fast Pointer)
func (s *Solution) Optimal(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	sol := Solution{}

	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7},
		{},
	}

	for i, tc := range testCases {
		fmt.Printf("========== Test Case %d ==========\n", i+1)
		fmt.Print("Input   : ")
		datastructure.PrintList(datastructure.CreateList(tc))

		fmt.Print("Brute   : ")
		datastructure.PrintList(sol.Brute(datastructure.CreateList(tc)))

		fmt.Print("Better  : ")
		datastructure.PrintList(sol.Better(datastructure.CreateList(tc)))

		fmt.Print("Optimal : ")
		datastructure.PrintList(sol.Optimal(datastructure.CreateList(tc)))

		fmt.Println()
	}
}
