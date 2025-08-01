package main

import (
	"reflect"
	"testing"
)

func TestLinkedListReverse(t *testing.T) {
	// Test Case 1: Empty list
	t.Run("Empty list", func(t *testing.T) {
		list := NewLinkedList()
		list.Reverse()
		expected := []int{}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 2: Single element
	t.Run("Single element", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(5)
		list.Reverse()
		expected := []int{5}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 3: Two elements
	t.Run("Two elements", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(1)
		list.Append(2)
		list.Reverse()
		expected := []int{2, 1}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 4: Three elements
	t.Run("Three elements", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.Reverse()
		expected := []int{3, 2, 1}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 5: Multiple elements
	t.Run("Multiple elements", func(t *testing.T) {
		list := NewLinkedList()
		for i := 1; i <= 5; i++ {
			list.Append(i)
		}
		list.Reverse()
		expected := []int{5, 4, 3, 2, 1}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 6: Negative numbers
	t.Run("Negative numbers", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(-1)
		list.Append(-2)
		list.Append(-3)
		list.Reverse()
		expected := []int{-3, -2, -1}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 7: Mixed positive and negative
	t.Run("Mixed positive and negative", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(-2)
		list.Append(0)
		list.Append(3)
		list.Append(-1)
		list.Reverse()
		expected := []int{-1, 3, 0, -2}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 8: Duplicate elements
	t.Run("Duplicate elements", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(1)
		list.Append(2)
		list.Append(2)
		list.Append(1)
		list.Reverse()
		expected := []int{1, 2, 2, 1}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 9: All same elements
	t.Run("All same elements", func(t *testing.T) {
		list := NewLinkedList()
		for i := 0; i < 4; i++ {
			list.Append(7)
		}
		list.Reverse()
		expected := []int{7, 7, 7, 7}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 10: Large numbers
	t.Run("Large numbers", func(t *testing.T) {
		list := NewLinkedList()
		list.Append(1000000)
		list.Append(2000000)
		list.Append(3000000)
		list.Reverse()
		expected := []int{3000000, 2000000, 1000000}
		result := list.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestLinkedListSize(t *testing.T) {
	// Test size function with different scenarios
	t.Run("Size after reverse operations", func(t *testing.T) {
		list := NewLinkedList()
		
		// Empty list
		if list.Size() != 0 {
			t.Errorf("Expected size 0, got %d", list.Size())
		}
		
		// Add elements
		list.Append(1)
		list.Append(2)
		list.Append(3)
		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}
		
		// Reverse should not change size
		list.Reverse()
		if list.Size() != 3 {
			t.Errorf("Expected size 3 after reverse, got %d", list.Size())
		}
	})
}
