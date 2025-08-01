package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		insert   int
		expected []int
	}{
		{
			name:     "Insert into empty tree",
			initial:  []int{},
			insert:   5,
			expected: []int{5},
		},
		{
			name:     "Insert smaller value as left child",
			initial:  []int{10},
			insert:   5,
			expected: []int{5, 10},
		},
		{
			name:     "Insert larger value as right child",
			initial:  []int{10},
			insert:   15,
			expected: []int{10, 15},
		},
		{
			name:     "Insert multiple values in order",
			initial:  []int{10, 5},
			insert:   15,
			expected: []int{5, 10, 15},
		},
		{
			name:     "Insert value creating left subtree",
			initial:  []int{10, 5, 15},
			insert:   3,
			expected: []int{3, 5, 10, 15},
		},
		{
			name:     "Insert value creating right subtree",
			initial:  []int{10, 5, 15},
			insert:   18,
			expected: []int{5, 10, 15, 18},
		},
		{
			name:     "Insert duplicate value",
			initial:  []int{10, 5, 15},
			insert:   10,
			expected: []int{5, 10, 15},
		},
		{
			name:     "Insert creating balanced tree",
			initial:  []int{10, 5, 15, 3},
			insert:   7,
			expected: []int{3, 5, 7, 10, 15},
		},
		{
			name:     "Insert negative value",
			initial:  []int{10, 5},
			insert:   -2,
			expected: []int{-2, 5, 10},
		},
		{
			name:     "Insert large value",
			initial:  []int{10, 5, 15},
			insert:   100,
			expected: []int{5, 10, 15, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			
			// Build initial tree
			for _, val := range tt.initial {
				root = Insert(root, val)
			}
			
			// Insert new value
			root = Insert(root, tt.insert)
			
			// Get inorder traversal
			result := InorderTraversal(root)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Insert() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		search   int
		expected bool
	}{
		{
			name:     "Search in empty tree",
			values:   []int{},
			search:   5,
			expected: false,
		},
		{
			name:     "Search root value",
			values:   []int{10},
			search:   10,
			expected: true,
		},
		{
			name:     "Search non-existent value in single node",
			values:   []int{10},
			search:   5,
			expected: false,
		},
		{
			name:     "Search left child",
			values:   []int{10, 5, 15},
			search:   5,
			expected: true,
		},
		{
			name:     "Search right child",
			values:   []int{10, 5, 15},
			search:   15,
			expected: true,
		},
		{
			name:     "Search deep left value",
			values:   []int{10, 5, 15, 3, 7},
			search:   3,
			expected: true,
		},
		{
			name:     "Search deep right value",
			values:   []int{10, 5, 15, 3, 7, 12, 18},
			search:   18,
			expected: true,
		},
		{
			name:     "Search non-existent small value",
			values:   []int{10, 5, 15},
			search:   1,
			expected: false,
		},
		{
			name:     "Search non-existent large value",
			values:   []int{10, 5, 15},
			search:   20,
			expected: false,
		},
		{
			name:     "Search negative value",
			values:   []int{10, 5, 15, -2},
			search:   -2,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			
			// Build tree
			for _, val := range tt.values {
				root = Insert(root, val)
			}
			
			// Search for value
			result := Search(root, tt.search)
			
			if result != tt.expected {
				t.Errorf("Search() = %v, want %v", result, tt.expected)
			}
		})
	}
}