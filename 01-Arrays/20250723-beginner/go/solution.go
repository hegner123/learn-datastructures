package main

// TODO: Implement the solution
// This function should find and return the maximum element in the given array
// Return 0 for empty arrays
func ArrayMax(arr []int) int {
	// TODO: Implement this function
	if len(arr) == 0 {
		return 0
	}
    var max int
    for i,item:= range arr {
        if i ==0{
            max = item
        } else {
            if item > max {
                max = item
            }
        }
    }
	return max
}

