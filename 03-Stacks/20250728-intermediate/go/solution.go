package main

import (
	"errors"
)

// MinStack represents a stack with minimum tracking and undo history
type MinStack struct {
	// TODO: Define the fields needed for your MinStack implementation
	// Consider what data structures you need:
	// - A stack for storing values
	// - A stack for tracking minimums  
	// - An array for storing operation history
}

// Operation represents a recorded operation for undo functionality
type Operation struct {
	// TODO: Define fields to store operation information
	// Consider what you need to store to enable undo:
	// - Operation type (push/pop)
	// - Value involved in the operation
}

// NewMinStack creates and returns a new MinStack instance
func NewMinStack() *MinStack {
	// TODO: Initialize and return a new MinStack
	// Initialize all necessary data structures
	return nil
}

// Push adds an element to the top of the stack and records the operation
func (ms *MinStack) Push(value int) {
	// TODO: Implement push operation
	// 1. Add value to the main stack
	// 2. Update minimum tracking
	// 3. Record the operation in history
}

// Pop removes and returns the top element, recording the operation
// Returns error if stack is empty
func (ms *MinStack) Pop() (int, error) {
	// TODO: Implement pop operation
	// 1. Check if stack is empty
	// 2. Remove from main stack and update minimum tracking
	// 3. Record the operation in history (store popped value for undo)
	// 4. Return the popped value
	return 0, errors.New("stack is empty")
}

// Peek returns the top element without removing it
// Returns error if stack is empty
func (ms *MinStack) Peek() (int, error) {
	// TODO: Implement peek operation
	// Return the top element without modifying the stack
	return 0, errors.New("stack is empty")
}

// GetMin returns the minimum element in the stack
// Returns error if stack is empty  
func (ms *MinStack) GetMin() (int, error) {
	// TODO: Implement get minimum operation
	// Return the current minimum value efficiently (O(1))
	return 0, errors.New("stack is empty")
}

// IsEmpty checks if the stack is empty
func (ms *MinStack) IsEmpty() bool {
	// TODO: Implement empty check
	return false
}

// Size returns the number of elements in the stack
func (ms *MinStack) Size() int {
	// TODO: Implement size operation
	return 0
}

// Undo reverses the last operation (push or pop)
// Returns true if successful, false if no history available
func (ms *MinStack) Undo() bool {
	// TODO: Implement undo operation
	// 1. Check if there's any operation to undo
	// 2. Get the last operation from history
	// 3. Reverse the operation:
	//    - If last was Push: remove the top element
	//    - If last was Pop: restore the popped element  
	// 4. Remove the operation from history
	// 5. Return true if successful, false if no history
	return false
}

// GetHistory returns an array of operation strings for debugging/testing
func (ms *MinStack) GetHistory() []string {
	// TODO: Implement get history operation
	// Return an array of strings describing each recorded operation
	// Format: "Push(value)" for push operations, "Pop()" for pop operations
	return nil
}

// ClearHistory clears the operation history but leaves stack contents unchanged
func (ms *MinStack) ClearHistory() {
	// TODO: Implement clear history operation
	// Clear the history array but don't modify the stack contents
}