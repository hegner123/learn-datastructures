package main

// TODO: Implement the MaxSlidingWindow function
// This function should find the maximum element in each sliding window of size k
// as the window moves from left to right through the array.
//
// Parameters:
//   - nums: array of integers
//   - k: size of the sliding window
//
// Returns:
//   - array of maximum values for each window position
//
// Example:
//   nums = [1,3,-1,-3,5,3,6,7], k = 3
//   Output: [3,3,5,5,6,7]
//
// Hint: Use a deque (double-ended queue) to efficiently track potential maximums
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	
	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0)
	
	for i := 0; i < len(nums); i++ {
		// Remove indices outside current window from front
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		
		// Remove indices with smaller values from back
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		
		// Add current index
		deque = append(deque, i)
		
		// Add maximum to result when window is complete
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	
	return result
}