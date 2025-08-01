package main

// TODO: Implement the ArrayContains function
// The function should return true if the target element is found in the array, false otherwise
// Signature: func ArrayContains(arr []int, target int) bool
func ArrayContains(arr []int, target int) bool {
	// TODO: Implement linear search through the array
	// Check each element to see if it matches the target
	// Return true if found, false if not found after checking all elements
    if len(arr)==0{
        return false
    }
    for _,item:= range arr{
        if item == target{
            return true
        }
    }
	return false
}
