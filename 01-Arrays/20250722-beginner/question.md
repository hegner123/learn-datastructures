# Array Sum Problem

## Problem Description

Write a function that calculates the sum of all elements in a given integer array. This is a fundamental array operation that demonstrates basic array traversal and accumulation patterns.

## Requirements

- Implement a function that takes an array of integers as input
- Return the sum of all elements in the array
- Handle edge cases including empty arrays
- The function should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func ArraySum(arr []int) int
```

**TypeScript:**
```typescript
function arraySum(arr: number[]): number
```

## Example Usage

**Go:**
```go
result := ArraySum([]int{1, 2, 3, 4, 5})  // returns 15
result := ArraySum([]int{})               // returns 0
result := ArraySum([]int{-1, -2, -3})     // returns -6
```

**TypeScript:**
```typescript
const result = arraySum([1, 2, 3, 4, 5]);  // returns 15
const result = arraySum([]);               // returns 0
const result = arraySum([-1, -2, -3]);     // returns -6
```

## Constraints

- Array length: 0 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- Return value should be the exact mathematical sum

## Test Cases

1. **Empty array**: `[]` → `0`
2. **Single element**: `[5]` → `5`
3. **Positive numbers**: `[1, 2, 3, 4, 5]` → `15`
4. **Negative numbers**: `[-1, -2, -3]` → `-6`
5. **Mixed numbers**: `[-5, 10, -3, 8]` → `10`
6. **All zeros**: `[0, 0, 0, 0]` → `0`
7. **Large numbers**: `[1000, 2000, 3000]` → `6000`
8. **Single negative**: `[-100]` → `-100`
9. **Alternating signs**: `[1, -1, 2, -2, 3]` → `3`
10. **Duplicate values**: `[7, 7, 7, 7]` → `28`

## Learning Objectives

- Understand basic array traversal patterns
- Practice handling edge cases (empty arrays)
- Learn accumulation patterns in programming
- Experience with integer arithmetic in array contexts
- Familiarize with testing strategies for array operations

## Difficulty Level

**Beginner** - This problem focuses solely on array operations without requiring additional data structures or complex algorithms.