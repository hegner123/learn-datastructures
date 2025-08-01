# Binary Trees

## Overview

A binary tree is a hierarchical data structure where each node has at most two children, commonly referred to as the left child and right child. This comprehensive guide covers all aspects of binary trees through focused modules.

## Module Structure

This documentation is organized into focused modules for better learning:

1. **[Binary Tree Structure](binary_tree_structure.md)** - Basic definitions, node structure, tree types, and fundamental concepts
2. **[Binary Tree Insertion](binary_tree_insertion.md)** - Various insertion methods including level-order, manual positioning, and specialized construction
3. **[Binary Tree Deletion](binary_tree_deletion.md)** - Deletion strategies for different scenarios including leaf deletion, subtree removal, and restructuring
4. **[Binary Tree Traversal](binary_tree_traversal.md)** - Complete coverage of all traversal methods (DFS, BFS, and specialized traversals)
5. **[Binary Tree Operations](binary_tree_operations.md)** - Utility methods for tree analysis, search operations, and transformations

## Quick Reference

### Key Properties
- **Root**: The topmost node in the tree
- **Leaf**: A node with no children  
- **Height**: The length of the longest path from root to leaf
- **Depth**: The length of the path from root to a particular node
- **Level**: All nodes at the same depth are at the same level

### Tree Types
- **Full Binary Tree**: Every node has either 0 or 2 children
- **Complete Binary Tree**: All levels are filled except possibly the last level
- **Perfect Binary Tree**: All internal nodes have 2 children and all leaves are at the same level
- **Balanced Binary Tree**: Height difference between left and right subtrees is at most 1

### Traversal Methods
- **Inorder**: Left → Root → Right
- **Preorder**: Root → Left → Right  
- **Postorder**: Left → Right → Root
- **Level Order**: Breadth-first traversal

### Time Complexity
- **Search**: O(n) - worst case, O(log n) - balanced tree
- **Insertion**: O(n) - worst case, O(log n) - balanced tree  
- **Deletion**: O(n) - worst case, O(log n) - balanced tree
- **Traversal**: O(n)

## Implementation Examples

For complete implementations with detailed code examples, see the individual module files:

- **[Binary Tree Structure](binary_tree_structure.md)** - Contains basic Go and TypeScript implementations with node structures and constructors
- **[Binary Tree Insertion](binary_tree_insertion.md)** - Multiple insertion strategies with complete working examples  
- **[Binary Tree Deletion](binary_tree_deletion.md)** - Various deletion approaches with full implementations
- **[Binary Tree Traversal](binary_tree_traversal.md)** - All traversal methods including recursive, iterative, and specialized variants
- **[Binary Tree Operations](binary_tree_operations.md)** - Utility methods and advanced operations with complete code

### Quick Example

Here's a minimal binary tree example in both languages:

```go
// Go - Basic Usage
bt := NewBinaryTree()
bt.Insert(1)
bt.Insert(2) 
bt.Insert(3)
fmt.Println("Level Order:", bt.LevelOrderTraversal()) // [1 2 3]
```

```typescript
// TypeScript - Basic Usage  
const bt = new BinaryTree<number>();
bt.insert(1);
bt.insert(2);
bt.insert(3);
console.log("Level Order:", bt.levelOrderTraversal()); // [1, 2, 3]
```

## Common Use Cases

- **Expression parsing**: Mathematical and logical expressions
- **File systems**: Directory structure representation
- **Decision trees**: AI and machine learning algorithms
- **Syntax trees**: Compilers and interpreters
- **Huffman coding**: Data compression algorithms
- **Game trees**: Game AI for decision making
- **Organization charts**: Hierarchical structures
- **Family trees**: Genealogical data representation

## Advantages

- Hierarchical data representation
- Efficient searching in balanced trees
- Natural recursive structure
- Flexible insertion and deletion
- Memory efficient for sparse data

## Disadvantages

- No constant time access to arbitrary elements
- Can become unbalanced leading to poor performance
- More complex than linear data structures
- Requires more memory for storing pointers
- Tree traversal can be memory intensive for large trees