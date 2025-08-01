package main

import "testing"

func TestArrayContains(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected bool
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: false,
		},
		{
			name:     "Single element - found",
			arr:      []int{42},
			target:   42,
			expected: true,
		},
		{
			name:     "Single element - not found",
			arr:      []int{42},
			target:   10,
			expected: false,
		},
		{
			name:     "Multiple elements - target at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: true,
		},
		{
			name:     "Multiple elements - target at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: true,
		},
		{
			name:     "Multiple elements - target in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "Multiple elements - target not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   10,
			expected: false,
		},
		{
			name:     "Array with negative numbers - found",
			arr:      []int{-5, -2, -8, -1},
			target:   -2,
			expected: true,
		},
		{
			name:     "Array with negative numbers - not found",
			arr:      []int{-5, -2, -8, -1},
			target:   2,
			expected: false,
		},
		{
			name:     "Array with duplicates - found",
			arr:      []int{7, 7, 7, 7},
			target:   7,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ArrayContains(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("ArrayContains(%v, %d) = %t, expected %t", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}