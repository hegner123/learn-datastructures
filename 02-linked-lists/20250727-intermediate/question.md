# Merge Two Sorted Linked Lists

## Problem Description

Implement a function to merge two sorted singly linked lists into a single sorted linked list. You need to create a basic linked list structure and implement the merge operation that combines two pre-sorted lists while maintaining the sorted order.

This problem combines linked list manipulation with array-based sorting concepts, requiring you to understand both data structures to effectively merge sorted sequences.

## Requirements

### Node Structure
- Each node should contain:
  - A data field to store integer values  
  - A next field to reference the next node in the sequence

### LinkedList Structure
- Should maintain a reference to the head node
- Handle empty list cases appropriately
- Provide functionality to build and inspect the list

### Operations to Implement

#### `append(data)`
- Adds a new node with the given data at the end of the list
- If list is empty, the new node becomes the head
- Time complexity: O(n)

#### `toArray()` / `toSlice()`
- Returns all elements in the list as an array (TypeScript) or slice (Go)
- Elements should be in order from head to tail
- Returns empty array/slice if list is empty
- Time complexity: O(n)

#### `createFromArray(data)` / `CreateFromSlice(data)` (static/standalone function)
- Creates a new LinkedList from an array/slice of sorted integers
- Elements should be added in the same order as provided
- Time complexity: O(n)

#### `mergeSortedLists(list1, list2)` (main function)
- Merges two sorted linked lists into a single sorted linked list
- Both input lists must already be sorted in ascending order
- The result should be a new list containing all elements in sorted order
- Should not modify the original lists
- Time complexity: O(n + m) where n and m are the lengths of the lists
- Space complexity: O(n + m) for the new result list

## Example Usage

### Input/Output Examples

**Example 1: Basic Merge**
```
list1 = createFromArray([1, 3, 5])      // List1: [1, 3, 5]
list2 = createFromArray([2, 4, 6])      // List2: [2, 4, 6]
result = mergeSortedLists(list1, list2)
result.toArray()                         // Returns: [1, 2, 3, 4, 5, 6]
```

**Example 2: Different Lengths**
```
list1 = createFromArray([1, 5, 9, 15])  // List1: [1, 5, 9, 15]
list2 = createFromArray([2, 3])         // List2: [2, 3]
result = mergeSortedLists(list1, list2)
result.toArray()                         // Returns: [1, 2, 3, 5, 9, 15]
```

**Example 3: One Empty List**
```
list1 = new LinkedList()                // List1: []
list2 = createFromArray([1, 3, 5])      // List2: [1, 3, 5]
result = mergeSortedLists(list1, list2)
result.toArray()                         // Returns: [1, 3, 5]
```

**Example 4: Both Empty Lists**
```
list1 = new LinkedList()                // List1: []
list2 = new LinkedList()                // List2: []
result = mergeSortedLists(list1, list2)
result.toArray()                         // Returns: []
```

## Edge Cases to Handle

1. **Both lists empty**: Should return an empty list
2. **One list empty**: Should return a copy of the non-empty list
3. **Lists with different lengths**: Should handle gracefully
4. **Duplicate elements**: Should preserve all duplicates in sorted order
5. **Single element lists**: Should work correctly with single-node lists
6. **Negative numbers**: Should handle negative integers correctly
7. **All elements from one list come first**: Handle cases where all elements of one list are smaller than the other
8. **Large number ranges**: Should work with large integer values

## Constraints

- Both input lists are guaranteed to be sorted in ascending order
- Work with integer data types
- Handle memory management appropriately (especially important in Go)
- Create a new result list rather than modifying input lists
- Maintain the integrity of all linked list structures throughout operations
- The merge should be stable (preserve relative order of equal elements)

## Implementation Notes

- This is an **intermediate-level** problem that combines linked list traversal with merge sort concepts
- The goal is to understand how to merge sorted sequences using pointer manipulation
- Consider using two pointers to traverse both lists simultaneously
- Pay attention to proper handling of remaining elements when one list is exhausted
- Ensure that the original lists remain unchanged after the merge operation
- The solution should demonstrate understanding of both linked list operations and array-based merging logic

## Test Cases

The solution should pass all of the following test cases:

1. **Both lists empty**: `[] + [] = []`
2. **First list empty**: `[] + [1, 3, 5] = [1, 3, 5]`
3. **Second list empty**: `[2, 4, 6] + [] = [2, 4, 6]`
4. **Simple merge**: `[1, 3, 5] + [2, 4, 6] = [1, 2, 3, 4, 5, 6]`
5. **Different lengths**: `[1, 5, 9, 15] + [2, 3] = [1, 2, 3, 5, 9, 15]`
6. **First list completely smaller**: `[1, 2, 3, 4] + [5, 6, 7, 8] = [1, 2, 3, 4, 5, 6, 7, 8]`
7. **Second list completely smaller**: `[10, 20, 30] + [1, 2, 3, 4, 5] = [1, 2, 3, 4, 5, 10, 20, 30]`
8. **Duplicate elements**: `[1, 3, 5, 5] + [2, 3, 4, 5] = [1, 2, 3, 3, 4, 5, 5, 5]`
9. **Single element lists**: `[5] + [3] = [3, 5]`
10. **Negative numbers**: `[-5, -1, 3] + [-3, 0, 2, 4] = [-5, -3, -1, 0, 2, 3, 4]`