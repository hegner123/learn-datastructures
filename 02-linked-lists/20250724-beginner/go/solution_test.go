package main

import (
	"reflect"
	"testing"
)

func TestLinkedListOperations(t *testing.T) {
	// Test Case 1: Insert at beginning
	t.Run("Insert at beginning", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Insert(10)
		ll.Insert(20)
		ll.Insert(30)
		
		expected := []int{30, 20, 10}
		result := ll.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 2: Append at end
	t.Run("Append at end", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		
		expected := []int{1, 2, 3}
		result := ll.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 3: Mixed insert and append
	t.Run("Mixed insert and append", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Insert(10)
		ll.Append(20)
		ll.Insert(5)
		ll.Append(30)
		
		expected := []int{5, 10, 20, 30}
		result := ll.ToSlice()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test Case 4: Delete existing element
	t.Run("Delete existing element", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)
		
		result := ll.Delete(2)
		if !result {
			t.Errorf("Expected true when deleting existing element")
		}
		
		expected := []int{1, 3, 4}
		slice := ll.ToSlice()
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v after deletion, got %v", expected, slice)
		}
	})

	// Test Case 5: Delete non-existent element
	t.Run("Delete non-existent element", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(1)
		ll.Append(2)
		
		result := ll.Delete(99)
		if result {
			t.Errorf("Expected false when deleting non-existent element")
		}
		
		expected := []int{1, 2}
		slice := ll.ToSlice()
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("List should remain unchanged after failed deletion")
		}
	})

	// Test Case 6: Delete from single element list
	t.Run("Delete from single element list", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(42)
		
		result := ll.Delete(42)
		if !result {
			t.Errorf("Expected true when deleting from single element list")
		}
		
		if ll.Size() != 0 {
			t.Errorf("Expected empty list after deleting single element")
		}
		
		expected := []int{}
		slice := ll.ToSlice()
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected empty slice, got %v", slice)
		}
	})

	// Test Case 7: Search existing element
	t.Run("Search existing element", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(5)
		ll.Append(10)
		ll.Append(15)
		
		if !ll.Search(10) {
			t.Errorf("Expected true when searching for existing element")
		}
		
		if !ll.Search(5) {
			t.Errorf("Expected true when searching for first element")
		}
		
		if !ll.Search(15) {
			t.Errorf("Expected true when searching for last element")
		}
	})

	// Test Case 8: Search non-existent element
	t.Run("Search non-existent element", func(t *testing.T) {
		ll := &LinkedList{}
		ll.Append(1)
		ll.Append(2)
		
		if ll.Search(99) {
			t.Errorf("Expected false when searching for non-existent element")
		}
	})

	// Test Case 9: Size operations
	t.Run("Size operations", func(t *testing.T) {
		ll := &LinkedList{}
		
		if ll.Size() != 0 {
			t.Errorf("Expected size 0 for empty list")
		}
		
		ll.Append(1)
		if ll.Size() != 1 {
			t.Errorf("Expected size 1 after adding one element")
		}
		
		ll.Insert(2)
		ll.Append(3)
		if ll.Size() != 3 {
			t.Errorf("Expected size 3 after adding three elements")
		}
		
		ll.Delete(2)
		if ll.Size() != 2 {
			t.Errorf("Expected size 2 after deleting one element")
		}
	})

	// Test Case 10: Empty list operations
	t.Run("Empty list operations", func(t *testing.T) {
		ll := &LinkedList{}
		
		if ll.Search(1) {
			t.Errorf("Expected false when searching empty list")
		}
		
		if ll.Delete(1) {
			t.Errorf("Expected false when deleting from empty list")
		}
		
		if ll.Size() != 0 {
			t.Errorf("Expected size 0 for empty list")
		}
		
		expected := []int{}
		slice := ll.ToSlice()
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected empty slice for empty list, got %v", slice)
		}
	})
}