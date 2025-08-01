package main

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Basic case",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Window size equals array length",
			nums:     []int{1, -1, 0, 5, 4},
			k:        5,
			expected: []int{5},
		},
		{
			name:     "All same elements",
			nums:     []int{4, 4, 4, 4},
			k:        2,
			expected: []int{4, 4, 4},
		},
		{
			name:     "Decreasing array",
			nums:     []int{9, 8, 7, 6, 5},
			k:        3,
			expected: []int{9, 8, 7},
		},
		{
			name:     "Increasing array",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 4, 5},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-7, -8, 7, 5, 7, 1, 6, 0},
			k:        4,
			expected: []int{7, 7, 7, 7, 7},
		},
		{
			name:     "Large window size",
			nums:     []int{1, 2, 3, 4},
			k:        3,
			expected: []int{3, 4},
		},
		{
			name:     "Mixed positive negative",
			nums:     []int{-1, 2, -3, 4, -5, 6},
			k:        2,
			expected: []int{2, 2, 4, 4, 6},
		},
		{
			name:     "Peak in middle",
			nums:     []int{1, 3, 1, 2, 0, 5},
			k:        3,
			expected: []int{3, 3, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}