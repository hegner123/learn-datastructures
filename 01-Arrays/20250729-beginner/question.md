# Array Contains Element Problem

## Problem Description

Write a function that determines whether a given target element exists in an integer array. This is a fundamental array operation that demonstrates basic linear search patterns and element comparison.

## Requirements

- Implement a function that takes an array of integers and a target value as input
- Return true if the target element is found in the array, false otherwise
- Handle edge cases including empty arrays
- The function should work with positive numbers, negative numbers, and zero
- Use linear search to traverse the array sequentially

## Function Signatures

**Go:**
```go
func ArrayContains(arr []int, target int) bool
```

**TypeScript:**
```typescript
function arrayContains(arr: number[], target: number): boolean
```

## Example Usage

**Go:**
```go
result := ArrayContains([]int{1, 2, 3, 4, 5}, 3)  // returns true
result := ArrayContains([]int{1, 2, 3, 4, 5}, 10) // returns false
result := ArrayContains([]int{}, 5)                // returns false
```

**TypeScript:**
```typescript
const result = arrayContains([1, 2, 3, 4, 5], 3);  // returns true
const result = arrayContains([1, 2, 3, 4, 5], 10); // returns false
const result = arrayContains([], 5);                // returns false
```

## Constraints

- Array length: 0 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- Target value: -1000 ≤ target ≤ 1000
- Return true if element is found, false if not found

## Test Cases

1. **Empty array**: `arr: [], target: 5` → `false`
2. **Single element - found**: `arr: [42], target: 42` → `true`
3. **Single element - not found**: `arr: [42], target: 10` → `false`
4. **Multiple elements - target at beginning**: `arr: [1, 2, 3, 4, 5], target: 1` → `true`
5. **Multiple elements - target at end**: `arr: [1, 2, 3, 4, 5], target: 5` → `true`
6. **Multiple elements - target in middle**: `arr: [1, 2, 3, 4, 5], target: 3` → `true`
7. **Multiple elements - target not found**: `arr: [1, 2, 3, 4, 5], target: 10` → `false`
8. **Array with negative numbers - found**: `arr: [-5, -2, -8, -1], target: -2` → `true`
9. **Array with negative numbers - not found**: `arr: [-5, -2, -8, -1], target: 2` → `false`
10. **Array with duplicates - found**: `arr: [7, 7, 7, 7], target: 7` → `true`

## Learning Objectives

- Understand basic linear search algorithms
- Practice array traversal and element comparison
- Learn to handle edge cases (empty arrays)
- Experience with boolean return values and conditional logic
- Familiarize with testing strategies for search operations

## Algorithm Hint

Iterate through each element in the array sequentially and compare it with the target value. If a match is found, return true immediately. If the loop completes without finding a match, return false.

## Difficulty Level

**Beginner** - This problem focuses solely on array operations without requiring additional data structures or complex algorithms. It demonstrates fundamental search patterns and comparison operations.