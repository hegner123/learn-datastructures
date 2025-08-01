package main

import "testing"

func TestArraySum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty array", []int{}, 0},
		{"single element", []int{5}, 5},
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-5, 10, -3, 8}, 10},
		{"all zeros", []int{0, 0, 0, 0}, 0},
		{"large numbers", []int{1000, 2000, 3000}, 6000},
		{"single negative", []int{-100}, -100},
		{"alternating signs", []int{1, -1, 2, -2, 3}, 3},
		{"duplicate values", []int{7, 7, 7, 7}, 28},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ArraySum(tc.input)
			if result != tc.expected {
				t.Errorf("ArraySum(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}