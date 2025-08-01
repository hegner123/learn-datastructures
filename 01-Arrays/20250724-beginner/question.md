# Array Minimum Element Problem

## Problem Description

Write a function that finds and returns the minimum element in a given integer array. This is a fundamental array operation that demonstrates basic array traversal and comparison patterns.

## Requirements

- Implement a function that takes an array of integers as input
- Return the minimum element in the array
- Handle edge cases including empty arrays (return 0 for empty arrays)
- The function should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func FindMinElement(arr []int) int
```

**TypeScript:**
```typescript
function findMinElement(arr: number[]): number
```

## Example Usage

**Go:**
```go
result := FindMinElement([]int{5, 2, 8, 1, 9})  // returns 1
result := FindMinElement([]int{})               // returns 0
result := FindMinElement([]int{-3, -1, -5})     // returns -5
```

**TypeScript:**
```typescript
const result = findMinElement([5, 2, 8, 1, 9]);  // returns 1
const result = findMinElement([]);               // returns 0
const result = findMinElement([-3, -1, -5]);     // returns -5
```

## Constraints

- Array length: 0 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- Return 0 for empty arrays
- Return the exact minimum value found

## Test Cases

1. **Empty array**: `[]` → `0`
2. **Single element**: `[42]` → `42`
3. **Ascending order**: `[1, 2, 3, 4, 5]` → `1`
4. **Descending order**: `[10, 8, 6, 4, 2]` → `2`
5. **All negative**: `[-5, -2, -10, -1]` → `-10`
6. **Mixed numbers**: `[-5, 10, -3, 8, 2]` → `-5`
7. **All same**: `[7, 7, 7, 7]` → `7`
8. **Single negative**: `[-100]` → `-100`
9. **Min at end**: `[50, 30, 20, 1]` → `1`
10. **Min in middle**: `[10, 5, -10, 15, 20]` → `-10`

## Learning Objectives

- Understand basic array traversal patterns
- Practice handling edge cases (empty arrays)
- Learn comparison operations in array contexts
- Experience with finding extrema in datasets
- Familiarize with testing strategies for array operations

## Difficulty Level

**Beginner** - This problem focuses solely on array operations without requiring additional data structures or complex algorithms.