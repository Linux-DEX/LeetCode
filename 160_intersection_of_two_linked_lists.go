/*
# Intersection of Two Linked Lists

Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

The test cases are generated such that there are no cycles anywhere in the entire linked structure.
The linked lists must retain their original structure after the function returns.

Example 1:

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'

Example 2:

Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'

Example 3:

Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection

Constraints:

1 <= m, n <= 3 * 10^4
1 <= Node.val <= 10^5
0 <= skipA <= m
0 <= skipB <= n

Follow up: O(m + n) time and O(1) memory.
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

// Brute Force
// Time  : O(m x n)
// Space : O(1)
func (s Solution) brute(headA, headB *ListNode) *ListNode {
	tempA := headA

	for tempA != nil {
		tempB := headB
		for tempB != nil {
			if tempA == tempB {
				return tempA
			}
			tempB = tempB.Next
		}
		tempA = tempA.Next
	}

	return nil
}

// Best
// Time  : O(m + n)
// Space : O(m)
func (s Solution) best(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	tempA := headA
	for tempA != nil {
		visited[tempA] = true
		tempA = tempA.Next
	}

	tempB := headB
	for tempB != nil {
		if visited[tempB] {
			return tempB
		}
		tempB = tempB.Next
	}

	return nil
}

// Optimal
// Time  : O(m + n)
// Space : O(1)
func (s Solution) optimal(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	for pA != pB {

		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
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

// Helper: Build two lists sharing the tail listA[skipA:] == listB[skipB:].
// skipA == len(listA) means no intersection; the shared values come from listA.
// Returns both heads and the expected intersection node (nil if none).
func buildIntersecting(listA, listB []int, skipA, skipB int) (*ListNode, *ListNode, *ListNode) {
	headA := buildList(listA)
	headB := buildList(listB)

	if skipA >= len(listA) || skipB >= len(listB) {
		return headA, headB, nil
	}

	// Node before the intersection in A, and the intersection node itself.
	var shared *ListNode = headA
	for i := 0; i < skipA; i++ {
		shared = shared.Next
	}

	if skipB == 0 {
		return headA, shared, shared
	}

	prevB := headB
	for i := 0; i < skipB-1; i++ {
		prevB = prevB.Next
	}
	prevB.Next = shared

	return headA, headB, shared
}

// Helper: Print result
func PrintResult(label string, got, want *ListNode) {
	switch {
	case got == nil && want == nil:
		fmt.Println(label, ": No intersection (OK)")
	case got == nil || want == nil:
		fmt.Println(label, ": MISMATCH")
	case got == want:
		fmt.Printf("%s : Intersected at '%d' (OK)\n", label, got.Val)
	default:
		fmt.Printf("%s : MISMATCH, got node with val %d\n", label, got.Val)
	}
}

func main() {
	sol := Solution{}

	tests := []struct {
		listA, listB []int
		skipA, skipB int
	}{
		{[]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3},
		{[]int{1, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1},
		{[]int{2, 6, 4}, []int{1, 5}, 3, 2},
	}

	for i, t := range tests {
		fmt.Println("Test", i+1)

		a, b, want := buildIntersecting(t.listA, t.listB, t.skipA, t.skipB)
		PrintResult("Brute", sol.brute(a, b), want)

		a, b, want = buildIntersecting(t.listA, t.listB, t.skipA, t.skipB)
		PrintResult("Best", sol.best(a, b), want)

		a, b, want = buildIntersecting(t.listA, t.listB, t.skipA, t.skipB)
		PrintResult("Optimal", sol.optimal(a, b), want)
	}
}
