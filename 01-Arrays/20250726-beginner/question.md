# Array Reverse Problem

## Problem Description

Write a function that reverses the elements of an integer array in-place. This is a fundamental array operation that demonstrates basic array manipulation using two-pointer technique.

## Requirements

- Implement a function that takes an array of integers as input
- Reverse the elements of the array in-place (modify the original array)
- Return the reversed array (same reference as input)
- The function should work with positive numbers, negative numbers, and zero
- Handle edge cases including empty arrays and single-element arrays

## Function Signatures

**Go:**
```go
func ReverseArray(arr []int) []int
```

**TypeScript:**
```typescript
function reverseArray(arr: number[]): number[]
```

## Example Usage

**Go:**
```go
result := ReverseArray([]int{1, 2, 3, 4, 5})  // returns [5, 4, 3, 2, 1]
result := ReverseArray([]int{})               // returns []
result := ReverseArray([]int{42})             // returns [42]
```

**TypeScript:**
```typescript
const result = reverseArray([1, 2, 3, 4, 5]);  // returns [5, 4, 3, 2, 1]
const result = reverseArray([]);               // returns []
const result = reverseArray([42]);             // returns [42]
```

## Constraints

- Array length: 0 ≤ length ≤ 1000
- Element values: -1000 ≤ element ≤ 1000
- The function must modify the array in-place
- Return the same array reference that was passed in

## Test Cases

1. **Empty array**: `[]` → `[]`
2. **Single element**: `[42]` → `[42]`
3. **Two elements**: `[1, 2]` → `[2, 1]`
4. **Odd length array**: `[1, 2, 3, 4, 5]` → `[5, 4, 3, 2, 1]`
5. **Even length array**: `[10, 20, 30, 40]` → `[40, 30, 20, 10]`
6. **Array with negative numbers**: `[-1, -2, -3, -4]` → `[-4, -3, -2, -1]`
7. **Array with mixed positive and negative**: `[-5, 10, -3, 8, 2]` → `[2, 8, -3, 10, -5]`
8. **Array with zeros**: `[0, 0, 1, 0, 0]` → `[0, 0, 1, 0, 0]`
9. **Array with duplicate values**: `[7, 7, 7, 7]` → `[7, 7, 7, 7]`
10. **Large numbers**: `[1000, 2000, 3000]` → `[3000, 2000, 1000]`

## Learning Objectives

- Understand the two-pointer technique for array manipulation
- Practice in-place array modification
- Learn to handle edge cases (empty arrays, single elements)
- Experience with array indexing and swapping elements
- Familiarize with testing strategies for array operations

## Algorithm Hint

Use two pointers - one starting from the beginning (left) and one from the end (right) of the array. Swap the elements at these positions and move the pointers towards the center until they meet.

## Difficulty Level

**Beginner** - This problem focuses solely on array operations without requiring additional data structures or complex algorithms. It demonstrates fundamental array manipulation techniques.
