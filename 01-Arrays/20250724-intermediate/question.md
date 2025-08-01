# Two Sum Problem

## Problem Description

Given an array of integers and a target sum, find two numbers in the array that add up to the target sum. Return the indices of the two numbers. You may assume that each input has exactly one solution, and you cannot use the same element twice.

## Requirements

- Implement a function that takes an array of integers and a target sum as input
- Return an array containing the indices of the two numbers that add up to the target
- The solution should be more efficient than checking all possible pairs (O(n²))
- Handle edge cases including arrays with exactly two elements
- The function should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func TwoSum(nums []int, target int) []int
```

**TypeScript:**
```typescript
function twoSum(nums: number[], target: number): number[]
```

## Example Usage

**Go:**
```go
result := TwoSum([]int{2, 7, 11, 15}, 9)  // returns [0, 1] (2 + 7 = 9)
result := TwoSum([]int{3, 2, 4}, 6)       // returns [1, 2] (2 + 4 = 6)
result := TwoSum([]int{3, 3}, 6)          // returns [0, 1] (3 + 3 = 6)
```

**TypeScript:**
```typescript
const result = twoSum([2, 7, 11, 15], 9);  // returns [0, 1] (2 + 7 = 9)
const result = twoSum([3, 2, 4], 6);       // returns [1, 2] (2 + 4 = 6)
const result = twoSum([3, 3], 6);          // returns [0, 1] (3 + 3 = 6)
```

## Constraints

- Array length: 2 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- Target sum: -2000 ≤ target ≤ 2000
- Exactly one solution exists for each test case
- Return indices in ascending order [smaller_index, larger_index]

## Test Cases

1. **Basic case**: `[2, 7, 11, 15], target=9` → `[0, 1]`
2. **Reverse order**: `[15, 11, 7, 2], target=9` → `[2, 3]`
3. **Same elements**: `[3, 3], target=6` → `[0, 1]`
4. **Negative numbers**: `[-1, -2, -3, -4, -5], target=-8` → `[2, 4]`
5. **Mixed signs**: `[-3, 4, 3, 90], target=0` → `[0, 2]`
6. **Zero target**: `[-2, 0, 1, 1], target=2` → `[2, 3]`
7. **Large numbers**: `[999, 1, 2], target=1000` → `[0, 1]`
8. **Adjacent elements**: `[1, 2, 3, 4], target=7` → `[2, 3]`
9. **First and last**: `[5, 1, 2, 8], target=13` → `[0, 3]`
10. **Middle elements**: `[10, 20, 30, 40, 50], target=70` → `[2, 3]`

## Learning Objectives

- Understand hash table/map usage for efficient lookups
- Practice time complexity optimization (from O(n²) to O(n))
- Learn complement calculation techniques (target - current_value)
- Experience with index tracking while iterating
- Familiarize with space-time trade-offs in algorithm design

## Algorithm Hint

Consider using a hash map to store values you've seen along with their indices. For each element, check if its complement (target - element) exists in the hash map.

## Difficulty Level

**Intermediate** - This problem combines basic array operations with hash map concepts and requires understanding of complement calculation and efficient lookup strategies.
