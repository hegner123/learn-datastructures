package main

import (
	"testing"
)

func TestBasicMinStackOperations(t *testing.T) {
	minStack := NewMinStack()
	
	// Push operations
	minStack.Push(3)
	minStack.Push(1)  
	minStack.Push(4)
	
	// Test GetMin
	min, err := minStack.GetMin()
	if err != nil {
		t.Errorf("GetMin() failed: %v", err)
	}
	if min != 1 {
		t.Errorf("GetMin() = %d, want 1", min)
	}
	
	// Test Peek
	top, err := minStack.Peek()
	if err != nil {
		t.Errorf("Peek() failed: %v", err)
	}
	if top != 4 {
		t.Errorf("Peek() = %d, want 4", top)
	}
	
	// Test Size
	if size := minStack.Size(); size != 3 {
		t.Errorf("Size() = %d, want 3", size)
	}
}

func TestPopAndMinUpdate(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(5)
	minStack.Push(2)
	minStack.Push(7)
	
	// Pop operation
	popped, err := minStack.Pop()
	if err != nil {
		t.Errorf("Pop() failed: %v", err)
	}
	if popped != 7 {
		t.Errorf("Pop() = %d, want 7", popped)
	}
	
	// Check min after pop
	min, err := minStack.GetMin()
	if err != nil {
		t.Errorf("GetMin() failed: %v", err)
	}
	if min != 2 {
		t.Errorf("GetMin() after pop = %d, want 2", min)
	}
}

func TestMultipleSameMinimumValues(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(2)
	minStack.Push(3)
	minStack.Push(2)
	minStack.Push(1)
	minStack.Push(2)
	
	// Pop the top 2
	popped, err := minStack.Pop()
	if err != nil {
		t.Errorf("Pop() failed: %v", err)
	}
	if popped != 2 {
		t.Errorf("Pop() = %d, want 2", popped)
	}
	
	// Min should still be 1
	min, err := minStack.GetMin()
	if err != nil {
		t.Errorf("GetMin() failed: %v", err)
	}
	if min != 1 {
		t.Errorf("GetMin() after pop = %d, want 1", min)
	}
}

func TestBasicUndoPushOperation(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(10)
	minStack.Push(5)
	
	// Undo the last push
	success := minStack.Undo()
	if !success {
		t.Error("Undo() should return true")
	}
	
	// Check result
	top, err := minStack.Peek()
	if err != nil {
		t.Errorf("Peek() failed: %v", err)
	}
	if top != 10 {
		t.Errorf("Peek() after undo = %d, want 10", top)
	}
	
	if size := minStack.Size(); size != 1 {
		t.Errorf("Size() after undo = %d, want 1", size)
	}
}

func TestBasicUndoPopOperation(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(1)
	minStack.Push(2)
	minStack.Pop()
	
	// Undo the pop
	success := minStack.Undo()
	if !success {
		t.Error("Undo() should return true")
	}
	
	// Check result
	top, err := minStack.Peek()
	if err != nil {
		t.Errorf("Peek() failed: %v", err)
	}
	if top != 2 {
		t.Errorf("Peek() after undo = %d, want 2", top)
	}
	
	if size := minStack.Size(); size != 2 {
		t.Errorf("Size() after undo = %d, want 2", size)
	}
}

func TestMultipleUndoOperations(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)
	
	// Undo twice
	minStack.Undo()
	minStack.Undo()
	
	// Check result
	top, err := minStack.Peek()
	if err != nil {
		t.Errorf("Peek() failed: %v", err)
	}
	if top != 1 {
		t.Errorf("Peek() after undos = %d, want 1", top)
	}
	
	min, err := minStack.GetMin()
	if err != nil {
		t.Errorf("GetMin() failed: %v", err)
	}
	if min != 1 {
		t.Errorf("GetMin() after undos = %d, want 1", min)
	}
}

func TestUndoOnEmptyHistory(t *testing.T) {
	minStack := NewMinStack()
	
	// Try to undo with no history
	success := minStack.Undo()
	if success {
		t.Error("Undo() should return false when no history")
	}
	
	// Stack should remain empty
	if !minStack.IsEmpty() {
		t.Error("Stack should remain empty after failed undo")
	}
}

func TestHistoryTracking(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(5)
	minStack.Push(3)
	minStack.Pop()
	minStack.Push(7)
	
	history := minStack.GetHistory()
	expected := []string{"Push(5)", "Push(3)", "Pop()", "Push(7)"}
	
	if len(history) != len(expected) {
		t.Errorf("History length = %d, want %d", len(history), len(expected))
	}
	
	for i, op := range expected {
		if i >= len(history) || history[i] != op {
			t.Errorf("History[%d] = %s, want %s", i, history[i], op)
		}
	}
}

func TestClearHistoryFunctionality(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(1)
	minStack.Push(2)
	minStack.Pop()
	
	// Clear history
	minStack.ClearHistory()
	
	// History should be empty
	history := minStack.GetHistory()
	if len(history) != 0 {
		t.Errorf("History length after clear = %d, want 0", len(history))
	}
	
	// But stack content should be unchanged
	if size := minStack.Size(); size != 1 {
		t.Errorf("Size() after clear history = %d, want 1", size)
	}
	
	top, err := minStack.Peek()
	if err != nil {
		t.Errorf("Peek() failed: %v", err)
	}
	if top != 1 {
		t.Errorf("Peek() after clear history = %d, want 1", top)
	}
}

func TestComplexMixedOperationsWithUndo(t *testing.T) {
	minStack := NewMinStack()
	
	minStack.Push(4)
	minStack.Push(1)
	minStack.Push(6)
	
	// Pop
	minStack.Pop()
	// Stack: [4, 1], min: 1
	min1, _ := minStack.GetMin()
	if min1 != 1 {
		t.Errorf("GetMin() after pop = %d, want 1", min1)
	}
	
	// Undo the pop
	minStack.Undo()
	// Stack: [4, 1, 6], min: 1
	min2, _ := minStack.GetMin()
	if min2 != 1 {
		t.Errorf("GetMin() after undo = %d, want 1", min2)
	}
	
	// Pop again
	minStack.Pop()
	// Stack: [4, 1], min: 1
	min3, _ := minStack.GetMin()
	if min3 != 1 {
		t.Errorf("GetMin() after second pop = %d, want 1", min3)
	}
	
	// Undo the pop again
	minStack.Undo()
	// Stack: [4, 1, 6], min: 1
	min4, _ := minStack.GetMin()
	if min4 != 1 {
		t.Errorf("GetMin() after second undo = %d, want 1", min4)
	}
}