package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the linked list structure
type LinkedList struct {
	Head *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	// TODO: Implement constructor
	newList := new(LinkedList)
	return newList
}

// Append adds a new node with the given data at the end of the list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// ToSlice returns all elements in the list as a slice
func (ll *LinkedList) ToSlice() []int {
	// TODO: Implement conversion to slice
	s := make([]int, 0)
	if ll.Head == nil {
		fmt.Println("early slice exit")
		return s
	}
	current := ll.Head
	s = append(s, current.Data)
	for current.Next != nil {
		current = current.Next
		s = append(s, current.Data)
	}
	return s
}

// Reverse reverses the linked list in-place
func (ll *LinkedList) Reverse() {
	// TODO: Implement in-place reversal of the linked list
	if ll.Head == nil {
		return
	}
	var previous *Node = nil
	var next *Node
	current := ll.Head
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next

	}
	ll.Head = previous

}

// Size returns the number of elements in the list
func (ll *LinkedList) Size() int {
	if ll.Head == nil {
		return 0
	}
	count := 1
	current := ll.Head
	for current.Next != nil {
		count++
		current = current.Next
	}
	return count
}
