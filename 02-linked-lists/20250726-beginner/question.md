# Linked List Reverse

## Problem Description

Implement a function to reverse a singly linked list in-place. 
You need to create a basic linked list structure and implement the reverse operation 
that modifies the existing list structure rather than creating a new one.

## Requirements

### Node Structure
- Each node should contain:
  - A data field to store integer values
  - A next field to reference the next node in the sequence

### LinkedList Structure
- Should maintain a reference to the head node
- Handle empty list cases appropriately
- Provide basic functionality to build and inspect the list

### Operations to Implement

#### `append(data)`
- Adds a new node with the given data at the end of the list
- If list is empty, the new node becomes the head
- Time complexity: O(n)

#### `reverse()`
- Reverses the linked list in-place by modifying the existing node connections
- Should change the direction of all next pointers
- The original tail becomes the new head
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

**Example 1: Basic Reverse**
```
list = new LinkedList()
list.append(1)         // List: [1]
list.append(2)         // List: [1, 2]
list.append(3)         // List: [1, 2, 3]
list.reverse()         // List: [3, 2, 1]
list.toArray()         // Returns: [3, 2, 1]
```

**Example 2: Single Element**
```
list = new LinkedList()
list.append(42)        // List: [42]
list.reverse()         // List: [42]
list.toArray()         // Returns: [42]
```

**Example 3: Empty List**
```
list = new LinkedList()
list.reverse()         // List: []
list.toArray()         // Returns: []
list.size()            // Returns: 0
```

**Example 4: Two Elements**
```
list = new LinkedList()
list.append(10)        // List: [10]
list.append(20)        // List: [10, 20]
list.reverse()         // List: [20, 10]
list.toArray()         // Returns: [20, 10]
```

## Edge Cases to Handle

1. **Empty list**: Reverse operation should work correctly on an empty list
2. **Single element**: List with one element should remain unchanged after reverse
3. **Two elements**: Should properly swap the two elements
4. **Negative numbers**: Should handle negative integers correctly
5. **Duplicate elements**: Should preserve duplicate values in reversed order
6. **Large numbers**: Should handle large integer values

## Constraints

- Use only the linked list data structure (no additional data structures required for the reverse operation)
- The reverse operation must be performed in-place (modify existing nodes, don't create new ones)
- Handle memory management appropriately (especially important in Go)
- Maintain the integrity of the linked list structure throughout all operations
- Work with integer data types

## Implementation Notes

- This is a **beginner-level** problem focusing on pointer manipulation in linked lists
- The goal is to understand how to traverse and modify link connections
- Pay special attention to updating the head pointer correctly
- Consider using three pointers (previous, current, next) for the reverse algorithm
- Ensure that no nodes are lost during the reversal process

## Test Cases

The solution should pass all of the following test cases:

1. **Empty list**: `[]` → `[]`
2. **Single element**: `[5]` → `[5]`
3. **Two elements**: `[1, 2]` → `[2, 1]`
4. **Three elements**: `[1, 2, 3]` → `[3, 2, 1]`
5. **Multiple elements**: `[1, 2, 3, 4, 5]` → `[5, 4, 3, 2, 1]`
6. **Negative numbers**: `[-1, -2, -3]` → `[-3, -2, -1]`
7. **Mixed positive and negative**: `[-2, 0, 3, -1]` → `[-1, 3, 0, -2]`
8. **Duplicate elements**: `[1, 2, 2, 1]` → `[1, 2, 2, 1]`
9. **All same elements**: `[7, 7, 7, 7]` → `[7, 7, 7, 7]`
10. **Large numbers**: `[1000000, 2000000, 3000000]` → `[3000000, 2000000, 1000000]`
