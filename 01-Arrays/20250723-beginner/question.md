# Array Maximum Element Problem

## Problem Description

Write a function that finds and returns the maximum element in a given integer array. This is a fundamental array operation that demonstrates basic array traversal and comparison patterns.

## Requirements

- Implement a function that takes an array of integers as input
- Return the maximum element in the array
- Handle edge cases including empty arrays (return 0 for empty arrays)
- The function should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func ArrayMax(arr []int) int
```

**TypeScript:**
```typescript
function arrayMax(arr: number[]): number
```

## Example Usage

**Go:**
```go
result := ArrayMax([]int{1, 2, 3, 4, 5})  // returns 5
result := ArrayMax([]int{})               // returns 0
result := ArrayMax([]int{-1, -2, -3})     // returns -1
```

**TypeScript:**
```typescript
const result = arrayMax([1, 2, 3, 4, 5]);  // returns 5
const result = arrayMax([]);               // returns 0
const result = arrayMax([-1, -2, -3]);     // returns -1
```

## Constraints

- Array length: 0 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- Return 0 for empty arrays
- Return the exact maximum value found

## Test Cases

1. **Empty array**: `[]` → `0`
2. **Single element**: `[42]` → `42`
3. **Ascending order**: `[1, 2, 3, 4, 5]` → `5`
4. **Descending order**: `[10, 8, 6, 4, 2]` → `10`
5. **All negative**: `[-5, -2, -10, -1]` → `-1`
6. **Mixed numbers**: `[-5, 10, -3, 8, 2]` → `10`
7. **All same**: `[7, 7, 7, 7]` → `7`
8. **Single negative**: `[-100]` → `-100`
9. **Max at beginning**: `[50, 1, 2, 3]` → `50`
10. **Max in middle**: `[1, 2, 100, 3, 4]` → `100`

## Learning Objectives

- Understand basic array traversal patterns
- Practice handling edge cases (empty arrays)
- Learn comparison operations in array contexts
- Experience with finding extrema in datasets
- Familiarize with testing strategies for array operations

## Difficulty Level

**Beginner** - This problem focuses solely on array operations without requiring additional data structures or complex algorithms.