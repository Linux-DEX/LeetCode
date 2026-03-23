/*
# Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:

Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []
*/
package main

import (
	"fmt"
)

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

// Brute Force (O(n) space)
func (s Solution) brute(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Store values in array
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	// Rewrite linked list in reverse
	curr = head
	for i := len(values) - 1; i >= 0; i-- {
		curr.Val = values[i]
		curr = curr.Next
	}

	return head
}

// Best (Recursive)
func (s Solution) best(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := s.best(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}

// Optimal (Iterative)
func (s Solution) optimal(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}

// Helper: Build List
func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	curr := head

	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}

	return head
}

// Helper: Print List
func PrintList(label string, head *ListNode) {
	fmt.Print(label, " : [")
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func main() {
	sol := Solution{}

	input := []int{1, 2, 3, 4, 5}

	head1 := buildList(input)
	head2 := buildList(input)
	head3 := buildList(input)

	PrintList("Brute", sol.brute(head1))
	PrintList("Best (Recursive)", sol.best(head2))
	PrintList("Optimal (Iterative)", sol.optimal(head3))
}
