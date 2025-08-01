package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedLists(t *testing.T) {
	// Test case 1: Both lists empty
	list1 := &LinkedList{}
	list2 := &LinkedList{}
	result := MergeSortedLists(list1, list2)
	expected := []int{}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 1 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 2: First list empty
	list1 = &LinkedList{}
	list2 = CreateFromSlice([]int{1, 3, 5})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 3, 5}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 2 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 3: Second list empty
	list1 = CreateFromSlice([]int{2, 4, 6})
	list2 = &LinkedList{}
	result = MergeSortedLists(list1, list2)
	expected = []int{2, 4, 6}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 3 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 4: Simple merge
	list1 = CreateFromSlice([]int{1, 3, 5})
	list2 = CreateFromSlice([]int{2, 4, 6})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 4 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 5: Lists with different lengths
	list1 = CreateFromSlice([]int{1, 5, 9, 15})
	list2 = CreateFromSlice([]int{2, 3})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 2, 3, 5, 9, 15}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 5 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 6: All elements from first list come first
	list1 = CreateFromSlice([]int{1, 2, 3, 4})
	list2 = CreateFromSlice([]int{5, 6, 7, 8})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 6 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 7: All elements from second list come first
	list1 = CreateFromSlice([]int{10, 20, 30})
	list2 = CreateFromSlice([]int{1, 2, 3, 4, 5})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 2, 3, 4, 5, 10, 20, 30}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 7 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 8: Duplicate elements
	list1 = CreateFromSlice([]int{1, 3, 5, 5})
	list2 = CreateFromSlice([]int{2, 3, 4, 5})
	result = MergeSortedLists(list1, list2)
	expected = []int{1, 2, 3, 3, 4, 5, 5, 5}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 8 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 9: Single element lists
	list1 = CreateFromSlice([]int{5})
	list2 = CreateFromSlice([]int{3})
	result = MergeSortedLists(list1, list2)
	expected = []int{3, 5}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 9 failed: expected %v, got %v", expected, result.ToSlice())
	}

	// Test case 10: Negative numbers
	list1 = CreateFromSlice([]int{-5, -1, 3})
	list2 = CreateFromSlice([]int{-3, 0, 2, 4})
	result = MergeSortedLists(list1, list2)
	expected = []int{-5, -3, -1, 0, 2, 3, 4}
	if !reflect.DeepEqual(result.ToSlice(), expected) {
		t.Errorf("Test 10 failed: expected %v, got %v", expected, result.ToSlice())
	}
}