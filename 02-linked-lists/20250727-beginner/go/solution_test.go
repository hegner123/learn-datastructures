package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLinkedListMethods(t *testing.T) {
	var failures []string

	// Test empty list
	ll := &LinkedList{}

	// Test size on empty list
	if size := ll.size(); size != 0 {
		failures = append(failures, "size() on empty list should return 0")
	}
	fmt.Printf("Testing size() on empty list - Input: empty list, Output: %d\n", ll.size())

	// Test toSlice on empty list
	if slice := ll.toSlice(); len(slice) != 0 {
		failures = append(failures, "toSlice() on empty list should return empty slice")
	}
	fmt.Printf("Testing toSlice() on empty list - Input: empty list, Output: %v\n", ll.toSlice())

	// Test append single element
	ll.append(10)
	fmt.Printf("Testing append(10) - Input: 10, List after append: %v\n", ll.toSlice())
	if size := ll.size(); size != 1 {
		failures = append(failures, "size() after appending one element should return 1")
	}
	fmt.Printf("Testing size() after append(10) - Input: list [10], Output: %d\n", ll.size())

	if slice := ll.toSlice(); !reflect.DeepEqual(slice, []int{10}) {
		failures = append(failures, "toSlice() after appending one element should return [10]")
	}
	fmt.Printf("Testing toSlice() after append(10) - Input: list [10], Output: %v\n", ll.toSlice())

	// Test append multiple elements
	ll.append(20)
	ll.append(30)
	fmt.Printf("Testing append(20), append(30) - Input: 20, 30, List after appends: %v\n", ll.toSlice())

	if size := ll.size(); size != 3 {
		failures = append(failures, "size() after appending three elements should return 3")
	}
	fmt.Printf("Testing size() after multiple appends - Input: list [10, 20, 30], Output: %d\n", ll.size())

	if slice := ll.toSlice(); !reflect.DeepEqual(slice, []int{10, 20, 30}) {
		failures = append(failures, "toSlice() after appending elements should return [10, 20, 30]")
	}
	fmt.Printf("Testing toSlice() after multiple appends - Input: list [10, 20, 30], Output: %v\n", ll.toSlice())

	// Test with even number of elements
	ll.append(40)
	fmt.Printf("Testing append(40) - Input: 40, List after append: %v\n", ll.toSlice())

	if size := ll.size(); size != 4 {
		failures = append(failures, "size() after appending four elements should return 4")
	}
	fmt.Printf("Testing size() with four elements - Input: list [10, 20, 30, 40], Output: %d\n", ll.size())

	if slice := ll.toSlice(); !reflect.DeepEqual(slice, []int{10, 20, 30, 40}) {
		failures = append(failures, "toSlice() after appending elements should return [10, 20, 30, 40]")
	}
	fmt.Printf("Testing toSlice() with four elements - Input: list [10, 20, 30, 40], Output: %v\n", ll.toSlice())

	// If any failures occurred, report them and fail the test
	if len(failures) > 0 {
		for _, failure := range failures {
			t.Error(failure)
		}
		t.Fatal("Linked list method verification failed - cannot proceed with solution tests")
		os.Exit(0)
	}

	tests := []struct {
		name           string
		elements       []int
		expectedMiddle int
		expectedFound  bool
		expectedSlice  []int
		expectedSize   int
	}{
		{
			name:           "Empty list",
			elements:       []int{},
			expectedMiddle: 0,
			expectedFound:  false,
			expectedSlice:  []int{},
			expectedSize:   0,
		},
		{
			name:           "Single element",
			elements:       []int{5},
			expectedMiddle: 5,
			expectedFound:  true,
			expectedSlice:  []int{5},
			expectedSize:   1,
		},
		{
			name:           "Two elements",
			elements:       []int{1, 2},
			expectedMiddle: 1,
			expectedFound:  true,
			expectedSlice:  []int{1, 2},
			expectedSize:   2,
		},
		{
			name:           "Three elements",
			elements:       []int{1, 2, 3},
			expectedMiddle: 2,
			expectedFound:  true,
			expectedSlice:  []int{1, 2, 3},
			expectedSize:   3,
		},
		{
			name:           "Four elements",
			elements:       []int{10, 20, 30, 40},
			expectedMiddle: 20,
			expectedFound:  true,
			expectedSlice:  []int{10, 20, 30, 40},
			expectedSize:   4,
		},
		{
			name:           "Five elements",
			elements:       []int{1, 2, 3, 4, 5},
			expectedMiddle: 3,
			expectedFound:  true,
			expectedSlice:  []int{1, 2, 3, 4, 5},
			expectedSize:   5,
		},
		{
			name:           "Six elements",
			elements:       []int{100, 200, 300, 400, 500, 600},
			expectedMiddle: 300,
			expectedFound:  true,
			expectedSlice:  []int{100, 200, 300, 400, 500, 600},
			expectedSize:   6,
		},
		{
			name:           "Seven elements with negatives",
			elements:       []int{-3, -2, -1, 0, 1, 2, 3},
			expectedMiddle: 0,
			expectedFound:  true,
			expectedSlice:  []int{-3, -2, -1, 0, 1, 2, 3},
			expectedSize:   7,
		},
		{
			name:           "Large numbers",
			elements:       []int{1000000, 2000000, 3000000, 4000000, 5000000},
			expectedMiddle: 3000000,
			expectedFound:  true,
			expectedSlice:  []int{1000000, 2000000, 3000000, 4000000, 5000000},
			expectedSize:   5,
		},
		{
			name:           "Duplicate elements",
			elements:       []int{7, 7, 7, 7, 7, 7},
			expectedMiddle: 7,
			expectedFound:  true,
			expectedSlice:  []int{7, 7, 7, 7, 7, 7},
			expectedSize:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}

			fmt.Printf("\n--- Running test case: %s ---\n", tt.name)

			// Build the list
			for _, element := range tt.elements {
				ll.append(element)
			}
			fmt.Printf("Built list with elements: %v\n", tt.elements)

			// Test findMiddle
			middle, found := ll.findMiddle()
			fmt.Printf("Testing findMiddle() - Input: %v, Output: (%d, %t), Expected: (%d, %t)\n",
				tt.elements, middle, found, tt.expectedMiddle, tt.expectedFound)
			if middle != tt.expectedMiddle || found != tt.expectedFound {
				t.Errorf("findMiddle() = (%d, %t), want (%d, %t)",
					middle, found, tt.expectedMiddle, tt.expectedFound)
			}

			// Test toSlice
			slice := ll.toSlice()
			fmt.Printf("Testing toSlice() - Input: list with elements %v, Output: %v, Expected: %v\n",
				tt.elements, slice, tt.expectedSlice)
			if !reflect.DeepEqual(slice, tt.expectedSlice) {
				t.Errorf("toSlice() = %v, want %v", slice, tt.expectedSlice)
			}

			// Test size
			size := ll.size()
			fmt.Printf("Testing size() - Input: list with elements %v, Output: %d, Expected: %d\n",
				tt.elements, size, tt.expectedSize)
			if size != tt.expectedSize {
				t.Errorf("size() = %d, want %d", size, tt.expectedSize)
			}
		})
	}
}
