# Binary Search Tree - Basic Operations

## Problem Description

Implement the fundamental operations of a Binary Search Tree (BST): insertion, search, and in-order traversal.

A Binary Search Tree is a binary tree data structure where:
- Each node has at most two children (left and right)
- For any node, all values in the left subtree are less than the node's value
- For any node, all values in the right subtree are greater than the node's value
- Duplicate values are typically not inserted (ignored in this implementation)

## Requirements

Implement the following three functions:

### 1. Insert Function
- **Go**: `Insert(root *TreeNode, val int) *TreeNode`
- **TypeScript**: `insert(root: TreeNode | null, val: number): TreeNode | null`

Insert a new value into the BST while maintaining the BST property. Return the root of the modified tree.

### 2. Search Function  
- **Go**: `Search(root *TreeNode, val int) bool`
- **TypeScript**: `search(root: TreeNode | null, val: number): boolean`

Search for a value in the BST. Return true if the value exists, false otherwise.

### 3. In-order Traversal Function
- **Go**: `InorderTraversal(root *TreeNode) []int`
- **TypeScript**: `inorderTraversal(root: TreeNode | null): number[]`

Perform an in-order traversal of the BST, returning values in sorted ascending order.

## TreeNode Structure

**Go:**
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

**TypeScript:**
```typescript
class TreeNode {
    val: number;
    left: TreeNode | null;
    right: TreeNode | null;
}
```

## Example Usage

```typescript
// Create empty tree
let root = null;

// Insert values: 10, 5, 15, 3, 7
root = insert(root, 10);
root = insert(root, 5);
root = insert(root, 15);
root = insert(root, 3);
root = insert(root, 7);

// Search for values
search(root, 7);  // returns true
search(root, 12); // returns false

// Get sorted values
inorderTraversal(root); // returns [3, 5, 7, 10, 15]
```

## Constraints

- Node values are integers
- Duplicate values should be ignored (not inserted)
- Tree can be empty (null/nil root)
- Values can be negative
- Maximum tree size is limited by available memory

## Test Cases

1. **Insert into empty tree**: Insert value 5 into empty tree → [5]
2. **Insert left child**: Insert 5 into tree with root 10 → [5, 10]  
3. **Insert right child**: Insert 15 into tree with root 10 → [10, 15]
4. **Multiple insertions**: Insert 15 into tree [10, 5] → [5, 10, 15]
5. **Left subtree growth**: Insert 3 into tree [10, 5, 15] → [3, 5, 10, 15]
6. **Right subtree growth**: Insert 18 into tree [10, 5, 15] → [5, 10, 15, 18]  
7. **Duplicate handling**: Insert 10 into tree [10, 5, 15] → [5, 10, 15]
8. **Balanced insertion**: Insert 7 into tree [10, 5, 15, 3] → [3, 5, 7, 10, 15]
9. **Negative values**: Insert -2 into tree [10, 5] → [-2, 5, 10]
10. **Large values**: Insert 100 into tree [10, 5, 15] → [5, 10, 15, 100]

Expected search behaviors:
- Search in empty tree returns false
- Search for existing values returns true
- Search for non-existent values returns false
- Search works for values at any depth in the tree