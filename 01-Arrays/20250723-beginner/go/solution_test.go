package main

import "testing"

func TestArrayMax(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty array", []int{}, 0},
		{"single element", []int{42}, 42},
		{"ascending order", []int{1, 2, 3, 4, 5}, 5},
		{"descending order", []int{10, 8, 6, 4, 2}, 10},
		{"all negative", []int{-5, -2, -10, -1}, -1},
		{"mixed numbers", []int{-5, 10, -3, 8, 2}, 10},
		{"all same", []int{7, 7, 7, 7}, 7},
		{"single negative", []int{-100}, -100},
		{"max at beginning", []int{50, 1, 2, 3}, 50},
		{"max in middle", []int{1, 2, 100, 3, 4}, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ArrayMax(tc.input)
			if result != tc.expected {
				t.Errorf("ArrayMax(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}