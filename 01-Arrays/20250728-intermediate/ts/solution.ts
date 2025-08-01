// TODO: Implement the maxSlidingWindow function
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
export function maxSlidingWindow(nums: number[], k: number): number[] {
    if (nums.length === 0 || k <= 0) {
        return [];
    }
    
    const result: number[] = [];
    const deque: number[] = [];
    
    for (let i = 0; i < nums.length; i++) {
        // Remove indices outside current window from front
        while (deque.length > 0 && deque[0] <= i - k) {
            deque.shift();
        }
        
        // Remove indices with smaller values from back
        while (deque.length > 0 && nums[deque[deque.length - 1]] <= nums[i]) {
            deque.pop();
        }
        
        // Add current index
        deque.push(i);
        
        // Add maximum to result when window is complete
        if (i >= k - 1) {
            result.push(nums[deque[0]]);
        }
    }
    
    return result;
}