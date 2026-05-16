package datastructure

import "fmt"

// ─────────────────────────────────────────────────────────────────────────────
// Types
// ─────────────────────────────────────────────────────────────────────────────

// ListNode is a generic singly-linked list node.
type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

// SinglyLinkedList manages a linked list via a head pointer.
// T is constrained to comparable so that Search and DeleteByValue work.
type SinglyLinkedList[T comparable] struct {
	Head *ListNode[T]
}

// ─────────────────────────────────────────────────────────────────────────────
// Standalone helper functions – operate directly on *ListNode[T]
// ─────────────────────────────────────────────────────────────────────────────

// CreateList builds a linked list from a slice and returns the head.
func CreateList[T any](vals []T) *ListNode[T] {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode[T]{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode[T]{Val: v}
		curr = curr.Next
	}
	return head
}

// ListToSlice converts a linked list to a slice of values.
func ListToSlice[T any](head *ListNode[T]) []T {
	var result []T
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// PrintList prints the list in [1, 2, 3] format.
func PrintList[T any](head *ListNode[T]) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

// ListLength returns the number of nodes.
func ListLength[T any](head *ListNode[T]) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// GetNode returns the node at the given 0-based index, or nil if out of range.
func GetNode[T any](head *ListNode[T], index int) *ListNode[T] {
	for i := 0; head != nil; i++ {
		if i == index {
			return head
		}
		head = head.Next
	}
	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
// SinglyLinkedList[T] methods
// ─────────────────────────────────────────────────────────────────────────────

// IsEmpty reports whether the list has no nodes.
func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.Head == nil
}

// Length returns the number of nodes.
func (s *SinglyLinkedList[T]) Length() int {
	return ListLength(s.Head)
}

// InsertAtEnd appends a value at the tail.
func (s *SinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &ListNode[T]{Val: val}
	if s.Head == nil {
		s.Head = newNode
		return
	}
	curr := s.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

// InsertAtHead prepends a value at the head.
func (s *SinglyLinkedList[T]) InsertAtHead(val T) {
	s.Head = &ListNode[T]{Val: val, Next: s.Head}
}

// InsertAtIndex inserts a value at the given 0-based index.
// Inserts at the end if index >= length.
func (s *SinglyLinkedList[T]) InsertAtIndex(index int, val T) {
	if index <= 0 {
		s.InsertAtHead(val)
		return
	}
	dummy := &ListNode[T]{Next: s.Head}
	curr := dummy
	for i := 0; i < index && curr.Next != nil; i++ {
		curr = curr.Next
	}
	curr.Next = &ListNode[T]{Val: val, Next: curr.Next}
	s.Head = dummy.Next
}

// DeleteAtHead removes the first node.
func (s *SinglyLinkedList[T]) DeleteAtHead() {
	if s.Head != nil {
		s.Head = s.Head.Next
	}
}

// DeleteAtTail removes the last node.
func (s *SinglyLinkedList[T]) DeleteAtTail() {
	if s.Head == nil {
		return
	}
	if s.Head.Next == nil {
		s.Head = nil
		return
	}
	curr := s.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

// DeleteByValue removes the first node whose Val equals val.
func (s *SinglyLinkedList[T]) DeleteByValue(val T) {
	dummy := &ListNode[T]{Next: s.Head}
	prev, curr := dummy, s.Head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
	s.Head = dummy.Next
}

// DeleteAtIndex removes the node at the given 0-based index.
func (s *SinglyLinkedList[T]) DeleteAtIndex(index int) {
	if s.Head == nil {
		return
	}
	if index == 0 {
		s.DeleteAtHead()
		return
	}
	curr := s.Head
	for i := 0; i < index-1 && curr.Next != nil; i++ {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

// Search returns the first node whose Val equals val, or nil.
func (s *SinglyLinkedList[T]) Search(val T) *ListNode[T] {
	curr := s.Head
	for curr != nil {
		if curr.Val == val {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

// Get returns the node at the given 0-based index, or nil.
func (s *SinglyLinkedList[T]) Get(index int) *ListNode[T] {
	return GetNode(s.Head, index)
}

// ToSlice returns all values as a slice.
func (s *SinglyLinkedList[T]) ToSlice() []T {
	return ListToSlice(s.Head)
}

// FromSlice replaces the list contents with values from the slice.
func (s *SinglyLinkedList[T]) FromSlice(vals []T) {
	s.Head = CreateList(vals)
}

// Print prints the list in [1, 2, 3] format.
func (s *SinglyLinkedList[T]) Print() {
	PrintList(s.Head)
}
