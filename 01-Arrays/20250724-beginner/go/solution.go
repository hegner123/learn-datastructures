package main

// FindMinElement finds the minimum element in an array of integers.
// Returns the minimum value found in the array.
// If the array is empty, returns 0.
func FindMinElement(arr []int) int {
	// TODO: Implement the function to find the minimum element in the array
	// This is a beginner-level problem that focuses on basic array traversal
	// and comparison operations.
	if len(arr) == 0 {
		return 0
	}
	min := 0
	for i, item := range arr {
		if i == 0 {
			min = item
		} else {
			if item < min {
				min = item
			}
		}

	}
	return min
}
