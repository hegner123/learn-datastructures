package main

import (
	"testing"
)

func TestBinaryTreeHeight(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: -1,
		},
		{
			name:     "Single node",
			values:   []int{1},
			expected: 0,
		},
		{
			name:     "Two nodes",
			values:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Three nodes",
			values:   []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "Complete binary tree - height 2",
			values:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: 2,
		},
		{
			name:     "Balanced tree - height 3",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: 3,
		},
		{
			name:     "Partial level tree",
			values:   []int{5, 3, 8, 2, 4, 7, 9, 1},
			expected: 3,
		},
		{
			name:     "Four nodes - height 2",
			values:   []int{1, 2, 3, 4},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree()
			
			// Build tree using level-order insertion
			for _, value := range tt.values {
				bt.Insert(value)
			}
			
			result := bt.Height()
			if result != tt.expected {
				t.Errorf("Height() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBinaryTreeSize(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: 0,
		},
		{
			name:     "Single node",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "Complete tree with 7 nodes",
			values:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: 7,
		},
		{
			name:     "Tree with 4 nodes",
			values:   []int{1, 2, 4, 5},
			expected: 4,
		},
		{
			name:     "Large complete tree",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree()
			
			// Build tree using level-order insertion
			for _, value := range tt.values {
				bt.Insert(value)
			}
			
			result := bt.Size()
			if result != tt.expected {
				t.Errorf("Size() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBinaryTreeIsEmpty(t *testing.T) {
	// Test empty tree
	bt := NewBinaryTree()
	if !bt.IsEmpty() {
		t.Errorf("IsEmpty() = false, want true for empty tree")
	}

	// Test non-empty tree
	bt.Insert(1)
	if bt.IsEmpty() {
		t.Errorf("IsEmpty() = true, want false for non-empty tree")
	}
}

