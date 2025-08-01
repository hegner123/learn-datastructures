# Binary Tree Basic Operations

## Problem Description

Implement a basic binary tree with fundamental operations. This problem focuses on understanding the core structure and properties of binary trees without requiring additional data structures.

You need to implement a `BinaryTree` class with the following methods:

1. **Constructor/NewBinaryTree**: Create an empty binary tree
2. **Insert**: Add nodes using level-order insertion (breadth-first)
3. **Height**: Calculate the height of the tree
4. **Size**: Count the total number of nodes
5. **IsEmpty**: Check if the tree is empty

## Requirements

### Height Calculation
- Height is defined as the number of edges in the longest path from root to leaf
- An empty tree has height -1
- A single node tree has height 0
- Height should be calculated recursively

### Size Calculation
- Size is the total number of nodes in the tree
- An empty tree has size 0

### Level-Order Insertion
- New nodes should be inserted in level-order (breadth-first)
- This creates a complete binary tree structure
- Use a queue-based approach for insertion

## Example Usage

### Go
```go
bt := NewBinaryTree()

// Insert values: creates a complete binary tree
bt.Insert(1)
bt.Insert(2) 
bt.Insert(3)
bt.Insert(4)
bt.Insert(5)

// Tree structure:
//       1
//      / \
//     2   3
//    / \
//   4   5

fmt.Println(bt.Height()) // Output: 2
fmt.Println(bt.Size())   // Output: 5
fmt.Println(bt.IsEmpty()) // Output: false
```

### TypeScript
```typescript
const bt = new BinaryTree();

// Insert values: creates a complete binary tree
bt.insert(1);
bt.insert(2);
bt.insert(3);
bt.insert(4);
bt.insert(5);

// Tree structure:
//       1
//      / \
//     2   3
//    / \
//   4   5

console.log(bt.height()); // Output: 2
console.log(bt.size());   // Output: 5
console.log(bt.isEmpty()); // Output: false
```

## Constraints

- Values are integers in the range [-1000, 1000]
- Maximum tree size: 1000 nodes
- Height calculation must handle empty trees correctly
- All operations should be implemented recursively where appropriate

## Test Cases

### Test Case 1: Empty Tree
- **Input**: Empty tree
- **Expected Height**: -1
- **Expected Size**: 0
- **Expected IsEmpty**: true

### Test Case 2: Single Node
- **Input**: [1]
- **Expected Height**: 0
- **Expected Size**: 1
- **Expected IsEmpty**: false

### Test Case 3: Two Nodes (Left Child)
- **Input**: [1, 2]
- **Expected Height**: 1
- **Expected Size**: 2

### Test Case 4: Two Nodes (Right Child Only)
- **Input**: [1, null, 3] (represented as [1, -1, 3] in tests)
- **Expected Height**: 1
- **Expected Size**: 2

### Test Case 5: Complete Binary Tree
- **Input**: [1, 2, 3, 4, 5, 6, 7]
- **Expected Height**: 2
- **Expected Size**: 7

### Test Case 6: Left Skewed Tree
- **Input**: [1, 2, null, 3, null, null, null, 4] (deep left path)
- **Expected Height**: 3
- **Expected Size**: 4

### Test Case 7: Right Skewed Tree
- **Input**: [1, null, 2, null, null, null, 3, null, null, null, null, null, null, null, 4]
- **Expected Height**: 3
- **Expected Size**: 4

### Test Case 8: Balanced Tree (Height 3)
- **Input**: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]
- **Expected Height**: 3
- **Expected Size**: 15

### Test Case 9: Unbalanced Tree (Deep Left)
- **Input**: [1, 2, 3, 4, 5, null, null, 6, 7, null, null, null, null, null, null, 8]
- **Expected Height**: 4
- **Expected Size**: 8

### Test Case 10: Mixed Structure
- **Input**: [5, 3, 8, 2, 4, 7, 9, 1]
- **Expected Height**: 3
- **Expected Size**: 8

## Implementation Notes

- Use recursive approaches for height and size calculations
- For level-order insertion, maintain a queue of nodes to process
- Handle null/empty cases properly in all methods
- The tree structure should support any valid binary tree configuration
- Test cases use -1 to represent null nodes in array representations