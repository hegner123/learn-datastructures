package main

import "errors"

// FindMaxElement finds the maximum element in an array of integers
// Returns an error if the array is empty
func FindMaxElement(arr []int) (int, error) {
	// TODO: Implement your solution here
	if len(arr) < 1 {
		return 0, errors.New("method not implemented")
	}
	var max int
	for index, item := range arr {
		if index == 0 {
			max = item
		}
		if item > max {
			max = item
		}
	}
	return max, nil
}
