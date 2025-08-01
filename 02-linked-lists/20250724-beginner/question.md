# Linked List Implementation

## Problem Description

Implement a basic singly linked list data structure with the following core operations:

1. **Insert at Beginning**: Insert a new element at the head of the list
2. **Append**: Add a new element at the end of the list  
3. **Delete**: Remove the first occurrence of a specific value
4. **Search**: Check if a value exists in the list
5. **Size**: Return the number of elements in the list
6. **Display**: Return all elements as an array/slice

## Requirements

### Node Structure
- Each node should contain:
  - A data field to store the value
  - A next field to reference the next node in the sequence

### LinkedList Structure
- Should maintain a reference to the head node
- Handle empty list cases appropriately

### Operations to Implement

#### `insert(data)` 
- Adds a new node with the given data at the beginning of the list
- The new node becomes the new head
- Time complexity: O(1)

#### `append(data)`
- Adds a new node with the given data at the end of the list
- If list is empty, the new node becomes the head
- Time complexity: O(n)

#### `delete(data)`
- Removes the first node that contains the specified data
- Returns `true` if element was found and removed, `false` otherwise
- Time complexity: O(n)

#### `search(data)`
- Searches for a node containing the specified data
- Returns `true` if found, `false` otherwise
- Time complexity: O(n)

#### `size()`
- Returns the total number of nodes in the list
- Time complexity: O(n)

#### `toArray()` / `toSlice()`
- Returns all elements in the list as an array (TypeScript) or slice (Go)
- Elements should be in order from head to tail
- Returns empty array/slice if list is empty
- Time complexity: O(n)

## Example Usage

### Input/Output Examples

**Example 1: Basic Operations**
```
list = new LinkedList()
list.insert(10)        // List: [10]
list.insert(20)        // List: [20, 10]
list.append(5)         // List: [20, 10, 5]
list.search(10)        // Returns: true
list.search(99)        // Returns: false
list.size()            // Returns: 3
list.toArray()         // Returns: [20, 10, 5]
```

**Example 2: Deletion**
```
list = new LinkedList()
list.append(1)         // List: [1]
list.append(2)         // List: [1, 2]
list.append(3)         // List: [1, 2, 3]
list.delete(2)         // Returns: true, List: [1, 3]
list.delete(99)        // Returns: false, List: [1, 3]
list.toArray()         // Returns: [1, 3]
```

**Example 3: Empty List**
```
list = new LinkedList()
list.search(1)         // Returns: false
list.size()            // Returns: 0
list.delete(1)         // Returns: false
list.toArray()         // Returns: []
```

## Edge Cases to Handle

1. **Empty list operations**: All operations should work correctly on an empty list
2. **Single element list**: Ensure operations work when list has only one element
3. **Delete non-existent element**: Should return false without modifying the list
4. **Insert/append with various data types**: Handle integers as the primary data type

## Constraints

- Use only the linked list data structure (no additional data structures required)
- Do not use built-in list/array operations for the core logic
- Handle memory management appropriately (especially important in Go)
- Maintain the integrity of the linked list structure throughout all operations

## Implementation Notes

- This is a **beginner-level** problem focusing on basic linked list operations
- The goal is to understand and implement the fundamental concepts of linked lists
- Pay attention to edge cases like empty lists and single-node lists
- Ensure proper handling of the head pointer in all operations