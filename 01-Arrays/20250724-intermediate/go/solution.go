package main

// TwoSum finds two numbers in the array that add up to the target sum
// and returns their indices.
//
// Parameters:
//   - nums: array of integers
//   - target: target sum to find
//
// Returns:
//   - []int: array containing the indices of the two numbers [index1, index2]
//
// TODO: Implement the TwoSum function
// Hint: Use a map to store values and their indices for O(n) time complexity
func TwoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for index, current := range nums {
		complement := target - current
		if storedIndex, exists := store[complement]; exists {
			return []int{storedIndex, index}
		}
		store[current] = index
	}
    return []int{}
}
