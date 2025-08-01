package main

// ReverseArray reverses the elements of an integer array in-place
// TODO: Implement this function to reverse the array elements
// The function should modify the original array and also return it
func ReverseArray(arr []int) []int {
	// TODO: Implement array reversal logic here
	// Hint: Use two pointers - one at the beginning and one at the end
	// Swap elements and move pointers towards center
	if len(arr) == 0 {
		return arr
	}
	p1 := 0
	p2 := len(arr) - 1
	for p1 < p2 {
		swap := arr[p1]
		arr[p1] = arr[p2]
		arr[p2] = swap
		p1++
		p2--
	}
	return arr
}

