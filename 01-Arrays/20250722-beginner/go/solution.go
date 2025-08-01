package main

// TODO: Implement the solution
// This function should calculate the sum of all elements in the given array
func ArraySum(arr []int) int {
	var sum int
	if len(arr) < 1 {
        return 0

	}
	for i, item := range arr {
		if i == 0 {
			sum = item
		} else {
            sum += item
        }

	}

	return sum
}

