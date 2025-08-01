package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Odd length array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Even length array",
			input:    []int{10, 20, 30, 40},
			expected: []int{40, 30, 20, 10},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-4, -3, -2, -1},
		},
		{
			name:     "Array with mixed positive and negative",
			input:    []int{-5, 10, -3, 8, 2},
			expected: []int{2, 8, -3, 10, -5},
		},
		{
			name:     "Array with zeros",
			input:    []int{0, 0, 1, 0, 0},
			expected: []int{0, 0, 1, 0, 0},
		},
		{
			name:     "Array with duplicate values",
			input:    []int{7, 7, 7, 7},
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "Large numbers",
			input:    []int{1000, 2000, 3000},
			expected: []int{3000, 2000, 1000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to preserve original for comparison
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)
			
			result := ReverseArray(inputCopy)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReverseArray(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}