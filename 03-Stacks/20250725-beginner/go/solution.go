package main

import "errors"

// Stack represents a stack data structure for integers
// TODO: Define the internal structure to hold stack elements
type Stack struct {
	items []int
}

// NewStack creates and returns a new empty stack
// TODO: Initialize and return a new Stack instance
// Make sure to properly initialize any internal data structures
func NewStack() *Stack {
	s := &Stack{
		items: make([]int, 0),
	}
	return s
}

// Push adds an element to the top of the stack
// TODO: Implement the push operation
// Add the item to the top of the stack
func (s *Stack) Push(data int) error {
	s.items = append(s.items, data)
	return nil
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
// TODO: Implement the pop operation
// Return the top element and remove it from the stack
// Return an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	index := len(s.items) - 1
	r := s.items[index]
	s.items = s.items[:index]

	return r, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
// TODO: Implement the peek operation
// Return the top element without modifying the stack
// Return an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	index := len(s.items) - 1
	return s.items[index], nil
}

// IsEmpty returns true if the stack is empty, false otherwise
// TODO: Implement the isEmpty check
// Return true if the stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
// TODO: Implement size calculation
// Return the current number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack
// TODO: Implement clear operation
// Remove all elements and reset the stack to empty state
func (s *Stack) Clear() error {
	e := make([]int, 0)
	s.items = e
	return nil
}
