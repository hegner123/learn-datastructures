package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Basic merge",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Empty first array",
			arr1:     []int{},
			arr2:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty second array",
			arr1:     []int{1, 2, 3},
			arr2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Both arrays empty",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "Different lengths",
			arr1:     []int{1, 5, 9, 10, 15, 20},
			arr2:     []int{2, 3, 8, 13},
			expected: []int{1, 2, 3, 5, 8, 9, 10, 13, 15, 20},
		},
		{
			name:     "First array larger elements",
			arr1:     []int{4, 5, 6},
			arr2:     []int{1, 2, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Duplicate elements",
			arr1:     []int{1, 3, 5},
			arr2:     []int{1, 3, 5},
			expected: []int{1, 1, 3, 3, 5, 5},
		},
		{
			name:     "Negative numbers",
			arr1:     []int{-5, -1, 2},
			arr2:     []int{-3, 0, 4},
			expected: []int{-5, -3, -1, 0, 2, 4},
		},
		{
			name:     "Single elements",
			arr1:     []int{5},
			arr2:     []int{3},
			expected: []int{3, 5},
		},
		{
			name:     "Large ranges",
			arr1:     []int{100, 200, 300},
			arr2:     []int{150, 250, 350, 400, 500},
			expected: []int{100, 150, 200, 250, 300, 350, 400, 500},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeSortedArrays(test.arr1, test.arr2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MergeSortedArrays(%v, %v) = %v; expected %v", test.arr1, test.arr2, result, test.expected)
			}
		})
	}
}