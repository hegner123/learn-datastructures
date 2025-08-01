package main

import (
	"reflect"
	"testing"
)

func TestBSTBasicOperations(t *testing.T) {
	tests := []struct {
		name        string
		operations  []interface{}
		expected    []interface{}
	}{
		{
			name: "Test 1: Insert single value and search",
			operations: []interface{}{
				[]string{"insert", "search"},
				[]int{5, 5},
			},
			expected: []interface{}{true},
		},
		{
			name: "Test 2: Insert multiple values and search existing",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "search", "search", "search"},
				[]int{10, 5, 15, 5, 10, 15},
			},
			expected: []interface{}{true, true, true},
		},
		{
			name: "Test 3: Search for non-existing value",
			operations: []interface{}{
				[]string{"insert", "insert", "search"},
				[]int{10, 20, 15},
			},
			expected: []interface{}{false},
		},
		{
			name: "Test 4: In-order traversal of simple tree",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "traversal"},
				[]int{10, 5, 15, 0},
			},
			expected: []interface{}{[]int{5, 10, 15}},
		},
		{
			name: "Test 5: In-order traversal with more complex tree",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "insert", "insert", "traversal"},
				[]int{50, 30, 20, 40, 70, 0},
			},
			expected: []interface{}{[]int{20, 30, 40, 50, 70}},
		},
		{
			name: "Test 6: Insert duplicate values (should not create duplicates)",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "traversal"},
				[]int{10, 10, 10, 0},
			},
			expected: []interface{}{[]int{10}},
		},
		{
			name: "Test 7: Left-skewed tree operations",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "search", "traversal"},
				[]int{30, 20, 10, 20, 0},
			},
			expected: []interface{}{true, []int{10, 20, 30}},
		},
		{
			name: "Test 8: Right-skewed tree operations",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "search", "traversal"},
				[]int{10, 20, 30, 25, 0},
			},
			expected: []interface{}{false, []int{10, 20, 30}},
		},
		{
			name: "Test 9: Empty tree operations",
			operations: []interface{}{
				[]string{"search", "traversal"},
				[]int{10, 0},
			},
			expected: []interface{}{false, []int{}},
		},
		{
			name: "Test 10: Large set insertion and search",
			operations: []interface{}{
				[]string{"insert", "insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "traversal"},
				[]int{50, 25, 75, 12, 37, 62, 87, 37, 100, 0},
			},
			expected: []interface{}{true, false, []int{12, 25, 37, 50, 62, 75, 87}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			operations := tt.operations[0].([]string)
			values := tt.operations[1].([]int)
			expectedResults := tt.expected
			resultIndex := 0

			for i, op := range operations {
				switch op {
				case "insert":
					bst.Insert(values[i])
				case "search":
					result := bst.Search(values[i])
					if result != expectedResults[resultIndex] {
						t.Errorf("Search operation failed. Expected %v, got %v", expectedResults[resultIndex], result)
					}
					resultIndex++
				case "traversal":
					result := bst.InOrderTraversal()
					expectedSlice := expectedResults[resultIndex].([]int)
					if !reflect.DeepEqual(result, expectedSlice) {
						t.Errorf("Traversal operation failed. Expected %v, got %v", expectedSlice, result)
					}
					resultIndex++
				}
			}
		})
	}
}

func TestBSTEdgeCases(t *testing.T) {
	t.Run("Test nil BST operations", func(t *testing.T) {
		bst := NewBST()
		if bst == nil {
			t.Skip("BST constructor returns nil - implement NewBST() first")
		}
		
		// Test operations on empty tree
		if bst.Search(10) != false {
			t.Error("Search on empty tree should return false")
		}
		
		traversal := bst.InOrderTraversal()
		if traversal == nil {
			traversal = []int{}
		}
		if len(traversal) != 0 {
			t.Error("Traversal on empty tree should return empty slice")
		}
	})
}