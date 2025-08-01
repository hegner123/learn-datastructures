package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Reverse order",
			nums:     []int{15, 11, 7, 2},
			target:   9,
			expected: []int{2, 3},
		},
		{
			name:     "Same elements",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "Mixed signs",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Zero target",
			nums:     []int{-2, 0, 1, 1},
			target:   2,
			expected: []int{2, 3},
		},
		{
			name:     "Large numbers",
			nums:     []int{999, 1, 2},
			target:   1000,
			expected: []int{0, 1},
		},
		{
			name:     "Adjacent elements",
			nums:     []int{1, 2, 3, 4},
			target:   7,
			expected: []int{2, 3},
		},
		{
			name:     "First and last",
			nums:     []int{5, 1, 2, 8},
			target:   13,
			expected: []int{0, 3},
		},
		{
			name:     "Middle elements",
			nums:     []int{10, 20, 30, 40, 50},
			target:   70,
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}