package main

import (
	"testing"
)

func TestStack_NewStack(t *testing.T) {
	stack := NewStack()
	if stack == nil {
		t.Error("NewStack() should return a non-nil stack")
	}
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if stack.Size() != 0 {
		t.Errorf("New stack size should be 0, got %d", stack.Size())
	}
}

func TestStack_PushSingleElement(t *testing.T) {
	stack := NewStack()
	stack.Push(10)
	
	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}
	if stack.Size() != 1 {
		t.Errorf("Stack size should be 1 after one push, got %d", stack.Size())
	}
	
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek should not return error on non-empty stack: %v", err)
	}
	if top != 10 {
		t.Errorf("Top element should be 10, got %d", top)
	}
}

func TestStack_PushMultipleElements(t *testing.T) {
	stack := NewStack()
	elements := []int{5, 15, 25, 35}
	
	for i, elem := range elements {
		stack.Push(elem)
		if stack.Size() != i+1 {
			t.Errorf("Stack size should be %d after pushing %d elements, got %d", i+1, i+1, stack.Size())
		}
	}
	
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek should not return error: %v", err)
	}
	if top != 35 {
		t.Errorf("Top element should be 35 (last pushed), got %d", top)
	}
}

func TestStack_PopSingleElement(t *testing.T) {
	stack := NewStack()
	stack.Push(42)
	
	popped, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop should not return error on non-empty stack: %v", err)
	}
	if popped != 42 {
		t.Errorf("Popped element should be 42, got %d", popped)
	}
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping the only element")
	}
	if stack.Size() != 0 {
		t.Errorf("Stack size should be 0 after popping all elements, got %d", stack.Size())
	}
}

func TestStack_PopMultipleElements(t *testing.T) {
	stack := NewStack()
	elements := []int{1, 2, 3, 4, 5}
	
	for _, elem := range elements {
		stack.Push(elem)
	}
	
	// Pop elements and verify LIFO order
	expectedOrder := []int{5, 4, 3, 2, 1}
	for i, expected := range expectedOrder {
		popped, err := stack.Pop()
		if err != nil {
			t.Errorf("Pop %d should not return error: %v", i+1, err)
		}
		if popped != expected {
			t.Errorf("Pop %d: expected %d, got %d", i+1, expected, popped)
		}
		if stack.Size() != len(elements)-i-1 {
			t.Errorf("Stack size should be %d after %d pops, got %d", len(elements)-i-1, i+1, stack.Size())
		}
	}
}

func TestStack_PopEmptyStack(t *testing.T) {
	stack := NewStack()
	
	_, err := stack.Pop()
	if err == nil {
		t.Error("Pop on empty stack should return an error")
	}
}

func TestStack_PeekEmptyStack(t *testing.T) {
	stack := NewStack()
	
	_, err := stack.Peek()
	if err == nil {
		t.Error("Peek on empty stack should return an error")
	}
}

func TestStack_PeekDoesNotModifyStack(t *testing.T) {
	stack := NewStack()
	elements := []int{100, 200, 300}
	
	for _, elem := range elements {
		stack.Push(elem)
	}
	
	initialSize := stack.Size()
	
	// Peek multiple times
	for i := 0; i < 3; i++ {
		top, err := stack.Peek()
		if err != nil {
			t.Errorf("Peek %d should not return error: %v", i+1, err)
		}
		if top != 300 {
			t.Errorf("Peek %d: expected 300, got %d", i+1, top)
		}
		if stack.Size() != initialSize {
			t.Errorf("Peek %d should not change stack size: expected %d, got %d", i+1, initialSize, stack.Size())
		}
	}
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack()
	elements := []int{7, 14, 21, 28, 35}
	
	for _, elem := range elements {
		stack.Push(elem)
	}
	
	if stack.IsEmpty() {
		t.Error("Stack should not be empty before clear")
	}
	
	stack.Clear()
	
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after clear")
	}
	if stack.Size() != 0 {
		t.Errorf("Stack size should be 0 after clear, got %d", stack.Size())
	}
	
	// Verify operations work correctly after clear
	stack.Push(99)
	if stack.Size() != 1 {
		t.Errorf("Stack size should be 1 after push following clear, got %d", stack.Size())
	}
}

func TestStack_LIFOBehavior(t *testing.T) {
	stack := NewStack()
	
	// Test LIFO with mixed operations
	stack.Push(1)
	stack.Push(2)
	
	first, _ := stack.Pop()
	if first != 2 {
		t.Errorf("First pop should return 2, got %d", first)
	}
	
	stack.Push(3)
	stack.Push(4)
	
	// Stack should now contain [1, 3, 4] (bottom to top)
	top, _ := stack.Peek()
	if top != 4 {
		t.Errorf("Top should be 4, got %d", top)
	}
	
	second, _ := stack.Pop()
	third, _ := stack.Pop()
	fourth, _ := stack.Pop()
	
	if second != 4 || third != 3 || fourth != 1 {
		t.Errorf("LIFO order violated: got %d, %d, %d, expected 4, 3, 1", second, third, fourth)
	}
	
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}