package main

import "testing"

func TestFindMaxElement(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Basic positive array", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9},
		{"All negative numbers", []int{-5, -2, -8, -1}, -1},
		{"Single element", []int{42}, 42},
		{"Two elements ascending", []int{1, 3}, 3},
		{"Two elements descending", []int{5, 2}, 5},
		{"Array with duplicates", []int{7, 7, 7, 7}, 7},
		{"Mixed positive and negative", []int{-3, 8, -1, 4, -9, 2}, 8},
		{"Max at beginning", []int{10, 3, 7, 1, 5}, 10},
		{"Max at end", []int{1, 3, 7, 5, 15}, 15},
		{"Large numbers", []int{1000000, 999999, 1000001, 500000}, 1000001},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := FindMaxElement(tc.input)
			if err != nil {
				t.Errorf("FindMaxElement(%v) = %d; expected %d\n %v", tc.input, result, tc.expected, err)
			}
			if result != tc.expected {
				t.Errorf("FindMaxElement(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

