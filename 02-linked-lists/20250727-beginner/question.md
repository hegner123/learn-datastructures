# Linked List Middle Element

## Problem Description

Implement a singly linked list with a method to find the middle element. You need to create a basic linked list structure and implement operations to build the list and find its middle element efficiently.

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

#### `findMiddle()`
- Returns the data of the middle node in the linked list
- For lists with an odd number of elements, return the exact middle element
- For lists with an even number of elements, return the **first** of the two middle elements
- Returns appropriate success/failure indication
- Time complexity: O(n)
- Space complexity: O(1)

#### `toArray()` / `toSlice()`
- Returns all elements in the list as an array (TypeScript) or slice (Go)
- Elements should be in order from head to tail
- Returns empty array/slice if list is empty
- Time complexity: O(n)

#### `size()`
- Returns the total number of nodes in the list
- Time complexity: O(n)

## Example Usage

### Input/Output Examples

**Example 1: Odd Number of Elements**
```
list = new LinkedList()
list.append(1)         // List: [1]
list.append(2)         // List: [1, 2]
list.append(3)         // List: [1, 2, 3]
list.findMiddle()      // Returns: 2 (middle element)
list.toArray()         // Returns: [1, 2, 3]
list.size()            // Returns: 3
```

**Example 2: Even Number of Elements**
```
list = new LinkedList()
list.append(10)        // List: [10]
list.append(20)        // List: [10, 20]
list.append(30)        // List: [10, 20, 30]
list.append(40)        // List: [10, 20, 30, 40]
list.findMiddle()      // Returns: 20 (first of two middle elements)
list.toArray()         // Returns: [10, 20, 30, 40]
list.size()            // Returns: 4
```

**Example 3: Single Element**
```
list = new LinkedList()
list.append(42)        // List: [42]
list.findMiddle()      // Returns: 42
list.toArray()         // Returns: [42]
list.size()            // Returns: 1
```

**Example 4: Empty List**
```
list = new LinkedList()
list.findMiddle()      // Returns: failure indication (0, false)
list.toArray()         // Returns: []
list.size()            // Returns: 0
```

## Edge Cases to Handle

1. **Empty list**: findMiddle should handle empty lists gracefully
2. **Single element**: Should return that element as the middle
3. **Two elements**: Should return the first element (position 0)
4. **Negative numbers**: Should handle negative integers correctly
5. **Duplicate elements**: Should work correctly with duplicate values
6. **Large numbers**: Should handle large integer values
7. **Odd vs Even length**: Correct middle calculation for both cases

## Constraints

- Use only the linked list data structure (no additional data structures required)
- Work with integer data types
- Handle memory management appropriately (especially important in Go)
- Maintain the integrity of the linked list structure throughout all operations
- The middle element calculation should follow the specified rule for even-length lists

## Implementation Notes

- This is a **beginner-level** problem focusing on basic linked list traversal and index calculation
- The goal is to understand how to navigate a linked list and determine position-based elements
- Consider different approaches: two-pass (count then find) vs one-pass solutions
- Pay attention to the definition of "middle" for even-length lists
- Ensure proper handling of edge cases like empty and single-element lists

## Test Cases

The solution should pass all of the following test cases:

1. **Empty list**: `[]` → middle: not found
2. **Single element**: `[5]` → middle: `5`
3. **Two elements**: `[1, 2]` → middle: `1` (first of two middle)
4. **Three elements**: `[1, 2, 3]` → middle: `2`
5. **Four elements**: `[10, 20, 30, 40]` → middle: `20` (first of two middle)
6. **Five elements**: `[1, 2, 3, 4, 5]` → middle: `3`
7. **Six elements**: `[100, 200, 300, 400, 500, 600]` → middle: `300` (first of two middle)
8. **Seven elements with negatives**: `[-3, -2, -1, 0, 1, 2, 3]` → middle: `0`
9. **Large numbers**: `[1000000, 2000000, 3000000, 4000000, 5000000]` → middle: `3000000`
10. **Duplicate elements**: `[7, 7, 7, 7, 7, 7]` → middle: `7` (first of two middle)