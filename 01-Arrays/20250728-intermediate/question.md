# Sliding Window Maximum Problem

## Problem Description

Given an array of integers and a sliding window of size k, find the maximum element in each window as the window slides from left to right across the array. Return an array containing the maximum values for each window position.

This problem combines array traversal with efficient data structure usage to maintain maximum values across sliding windows, demonstrating the practical application of deques (double-ended queues) for optimization.

## Requirements

- Implement a function that takes an array of integers and window size k as input
- Return an array of maximum values for each sliding window position
- The solution should be more efficient than recalculating the maximum for each window (O(n*k))
- Handle edge cases including arrays smaller than or equal to window size
- The function should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func MaxSlidingWindow(nums []int, k int) []int
```

**TypeScript:**
```typescript
function maxSlidingWindow(nums: number[], k: number): number[]
```

## Example Usage

**Go:**
```go
result := MaxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3)  // returns [3,3,5,5,6,7]
result := MaxSlidingWindow([]int{1}, 1)                  // returns [1]
result := MaxSlidingWindow([]int{1,2,3,4,5}, 3)         // returns [3,4,5]
```

**TypeScript:**
```typescript
const result = maxSlidingWindow([1,3,-1,-3,5,3,6,7], 3);  // returns [3,3,5,5,6,7]
const result = maxSlidingWindow([1], 1);                  // returns [1]
const result = maxSlidingWindow([1,2,3,4,5], 3);         // returns [3,4,5]
```

## Detailed Example

For array `[1,3,-1,-3,5,3,6,7]` with window size `k=3`:

```
Window position 1: [1  3  -1] -3  5  3  6  7  → max = 3
Window position 2:  1 [3  -1  -3] 5  3  6  7  → max = 3
Window position 3:  1  3 [-1  -3  5] 3  6  7  → max = 5
Window position 4:  1  3  -1 [-3  5  3] 6  7  → max = 5
Window position 5:  1  3  -1  -3 [5  3  6] 7  → max = 6
Window position 6:  1  3  -1  -3  5 [3  6  7] → max = 7
```

Result: `[3, 3, 5, 5, 6, 7]`

## Constraints

- Array length: 1 ≤ length ≤ 10^5  
- Element values: -10^4 ≤ element ≤ 10^4
- Window size: 1 ≤ k ≤ array length
- Must return exactly (n-k+1) elements where n is array length

## Test Cases

1. **Basic case**: `[1,3,-1,-3,5,3,6,7], k=3` → `[3,3,5,5,6,7]`
2. **Single element**: `[1], k=1` → `[1]`
3. **Window equals array**: `[1,-1,0,5,4], k=5` → `[5]`
4. **All same elements**: `[4,4,4,4], k=2` → `[4,4,4]`
5. **Decreasing array**: `[9,8,7,6,5], k=3` → `[9,8,7]`
6. **Increasing array**: `[1,2,3,4,5], k=3` → `[3,4,5]`
7. **Negative numbers**: `[-7,-8,7,5,7,1,6,0], k=4` → `[7,7,7,7,7]`
8. **Large window**: `[1,2,3,4], k=3` → `[3,4]`
9. **Mixed signs**: `[-1,2,-3,4,-5,6], k=2` → `[2,2,4,4,6]`
10. **Peak in middle**: `[1,3,1,2,0,5], k=3` → `[3,3,2,5]`

## Learning Objectives

- Understand sliding window technique with efficient maximum tracking
- Practice using deque (double-ended queue) data structure for optimization
- Learn to maintain relevant candidates while discarding obsolete ones
- Experience with amortized time complexity analysis
- Apply monotonic deque pattern for range query optimization
- Understand trade-offs between time and space complexity

## Algorithm Hint

Use a deque to store indices of array elements. Maintain the deque such that:
1. Indices are stored in decreasing order of their corresponding values
2. The front of the deque always contains the index of the maximum element in the current window
3. Remove indices that are outside the current window from the front
4. Remove indices from the back whose corresponding values are smaller than the current element

This achieves O(n) time complexity as each element is added and removed from the deque at most once.

## Difficulty Level

**Intermediate** - This problem combines array operations with deque data structure usage. It requires understanding of efficient window management and the monotonic deque optimization pattern, representing a step up from basic array traversal problems.