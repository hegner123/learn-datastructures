package main

import "testing"

func TestFindMinElement(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			arr:      []int{42},
			expected: 42,
		},
		{
			name:     "Ascending order",
			arr:      []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Descending order",
			arr:      []int{10, 8, 6, 4, 2},
			expected: 2,
		},
		{
			name:     "All negative",
			arr:      []int{-5, -2, -10, -1},
			expected: -10,
		},
		{
			name:     "Mixed numbers",
			arr:      []int{-5, 10, -3, 8, 2},
			expected: -5,
		},
		{
			name:     "All same",
			arr:      []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "Single negative",
			arr:      []int{-100},
			expected: -100,
		},
		{
			name:     "Min at end",
			arr:      []int{50, 30, 20, 1},
			expected: 1,
		},
		{
			name:     "Min in middle",
			arr:      []int{10, 5, -10, 15, 20},
			expected: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMinElement(tt.arr)
			if result != tt.expected {
				t.Errorf("FindMinElement(%v) = %d; expected %d", tt.arr, result, tt.expected)
			}
		})
	}
}