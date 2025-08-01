# Merge Sorted Arrays Problem

## Problem Description

Given two pre-sorted arrays of integers, merge them into a single sorted array containing all elements from both input arrays. The merged array should maintain the sorted order and preserve all duplicate elements from both arrays.

This problem demonstrates efficient array merging techniques using the two-pointer approach, which is fundamental for understanding merge sort algorithms and working with pre-sorted data.

## Requirements

- Implement a function that takes two sorted arrays as input
- Return a new array containing all elements in sorted order
- Handle arrays of different lengths efficiently
- Preserve all duplicate elements from both arrays
- Use O(m+n) time complexity where m and n are the lengths of the input arrays
- The solution should work with positive numbers, negative numbers, and zero

## Function Signatures

**Go:**
```go
func MergeSortedArrays(arr1 []int, arr2 []int) []int
```

**TypeScript:**
```typescript
function mergeSortedArrays(arr1: number[], arr2: number[]): number[]
```

## Example Usage

**Go:**
```go
result := MergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})  // returns [1, 2, 3, 4, 5, 6]
result := MergeSortedArrays([]int{}, []int{1, 2, 3})         // returns [1, 2, 3]
result := MergeSortedArrays([]int{1, 3, 5}, []int{1, 3, 5}) // returns [1, 1, 3, 3, 5, 5]
```

**TypeScript:**
```typescript
const result = mergeSortedArrays([1, 3, 5], [2, 4, 6]);  // returns [1, 2, 3, 4, 5, 6]
const result = mergeSortedArrays([], [1, 2, 3]);         // returns [1, 2, 3]
const result = mergeSortedArrays([1, 3, 5], [1, 3, 5]); // returns [1, 1, 3, 3, 5, 5]
```

## Detailed Example

For arrays `[1, 5, 9, 10, 15, 20]` and `[2, 3, 8, 13]`:

```
Step 1: Compare 1 vs 2 → add 1: [1]
Step 2: Compare 5 vs 2 → add 2: [1, 2]
Step 3: Compare 5 vs 3 → add 3: [1, 2, 3]
Step 4: Compare 5 vs 8 → add 5: [1, 2, 3, 5]
Step 5: Compare 9 vs 8 → add 8: [1, 2, 3, 5, 8]
Step 6: Compare 9 vs 13 → add 9: [1, 2, 3, 5, 8, 9]
Step 7: Compare 10 vs 13 → add 10: [1, 2, 3, 5, 8, 9, 10]
Step 8: Compare 15 vs 13 → add 13: [1, 2, 3, 5, 8, 9, 10, 13]
Step 9: Remaining from arr1: [1, 2, 3, 5, 8, 9, 10, 13, 15, 20]
```

Result: `[1, 2, 3, 5, 8, 9, 10, 13, 15, 20]`

## Constraints

- Array lengths: 0 ≤ length ≤ 1000 for each array
- Element values: -1000 ≤ element ≤ 1000
- Both input arrays are guaranteed to be sorted in ascending order
- The merged array must contain exactly (len(arr1) + len(arr2)) elements
- Time complexity must be O(m+n), space complexity O(m+n)

## Test Cases

1. **Basic merge**: `[1, 3, 5], [2, 4, 6]` → `[1, 2, 3, 4, 5, 6]`
2. **Empty first array**: `[], [1, 2, 3]` → `[1, 2, 3]`
3. **Empty second array**: `[1, 2, 3], []` → `[1, 2, 3]`
4. **Both arrays empty**: `[], []` → `[]`
5. **Different lengths**: `[1, 5, 9, 10, 15, 20], [2, 3, 8, 13]` → `[1, 2, 3, 5, 8, 9, 10, 13, 15, 20]`
6. **First array larger elements**: `[4, 5, 6], [1, 2, 3]` → `[1, 2, 3, 4, 5, 6]`
7. **Duplicate elements**: `[1, 3, 5], [1, 3, 5]` → `[1, 1, 3, 3, 5, 5]`
8. **Negative numbers**: `[-5, -1, 2], [-3, 0, 4]` → `[-5, -3, -1, 0, 2, 4]`
9. **Single elements**: `[5], [3]` → `[3, 5]`
10. **Large ranges**: `[100, 200, 300], [150, 250, 350, 400, 500]` → `[100, 150, 200, 250, 300, 350, 400, 500]`

## Learning Objectives

- Master the two-pointer technique for array operations
- Understand efficient merging algorithms used in merge sort
- Practice handling edge cases with empty arrays and different lengths
- Learn optimal time and space complexity analysis
- Experience with maintaining sorted order while merging data
- Develop skills in array traversal with multiple indices

## Algorithm Hint

Use two pointers to traverse both arrays simultaneously:
1. Initialize pointers at the beginning of each array
2. Compare elements at current positions
3. Add the smaller element to the result array and advance that pointer
4. When one array is exhausted, append all remaining elements from the other array
5. This approach ensures O(m+n) time complexity with a single pass through both arrays

## Difficulty Level

**Intermediate** - This problem combines fundamental array operations with the two-pointer technique. While it uses only arrays as the primary data structure, it requires understanding of efficient merging algorithms and careful handling of multiple array indices, making it more complex than basic array traversal problems.