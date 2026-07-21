/*
21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

---
Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
*/
package main

import (
	"fmt"
	"sort"

	"linux-dex/leetcode/datastructure"
)

type Solution struct{}

// Brute Force (Collect + Sort)
// Time: O((n+m) log(n+m))
// Space: O(n+m)
func (s Solution) brute(
	list1 *datastructure.ListNode[int],
	list2 *datastructure.ListNode[int],
) *datastructure.ListNode[int] {

	values := []int{}

	for list1 != nil {
		values = append(values, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		values = append(values, list2.Val)
		list2 = list2.Next
	}

	sort.Ints(values)

	return datastructure.CreateList(values)
}

// Better (Create New List)
// Time: O(n+m)
// Space: O(n+m)
func (s Solution) better(
	list1 *datastructure.ListNode[int],
	list2 *datastructure.ListNode[int],
) *datastructure.ListNode[int] {

	dummy := &datastructure.ListNode[int]{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = &datastructure.ListNode[int]{Val: list1.Val}
			list1 = list1.Next
		} else {
			curr.Next = &datastructure.ListNode[int]{Val: list2.Val}
			list2 = list2.Next
		}
		curr = curr.Next
	}

	for list1 != nil {
		curr.Next = &datastructure.ListNode[int]{Val: list1.Val}
		curr = curr.Next
		list1 = list1.Next
	}

	for list2 != nil {
		curr.Next = &datastructure.ListNode[int]{Val: list2.Val}
		curr = curr.Next
		list2 = list2.Next
	}

	return dummy.Next
}

// Optimal (Reuse Existing Nodes)
// Time: O(n+m)
// Space: O(1)
func (s Solution) optimal(
	list1 *datastructure.ListNode[int],
	list2 *datastructure.ListNode[int],
) *datastructure.ListNode[int] {

	dummy := &datastructure.ListNode[int]{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

func main() {
	sol := Solution{}

	testCases := []struct {
		list1 []int
		list2 []int
	}{
		{
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
		},
		{
			list1: []int{},
			list2: []int{},
		},
		{
			list1: []int{},
			list2: []int{0},
		},
		{
			list1: []int{1, 3, 5},
			list2: []int{2, 4, 6},
		},
		{
			list1: []int{1, 2, 2},
			list2: []int{1, 1, 2},
		},
	}

	for i, tc := range testCases {
		fmt.Printf("\n========== Test Case %d ==========\n", i+1)

		fmt.Print("Input List1 : ")
		datastructure.PrintList(datastructure.CreateList(tc.list1))

		fmt.Print("Input List2 : ")
		datastructure.PrintList(datastructure.CreateList(tc.list2))

		fmt.Print("Brute       : ")
		datastructure.PrintList(
			sol.brute(
				datastructure.CreateList(tc.list1),
				datastructure.CreateList(tc.list2),
			),
		)

		fmt.Print("Better      : ")
		datastructure.PrintList(
			sol.better(
				datastructure.CreateList(tc.list1),
				datastructure.CreateList(tc.list2),
			),
		)

		fmt.Print("Optimal     : ")
		datastructure.PrintList(
			sol.optimal(
				datastructure.CreateList(tc.list1),
				datastructure.CreateList(tc.list2),
			),
		)
	}
}
