# Binary Search Tree Implementation

## Problem Description

Implement a complete Binary Search Tree (BST) data structure with the fundamental operations: insert, search, delete, and in-order traversal. This is a direct implementation of the BST data structure to practice the core concepts and algorithms.

A Binary Search Tree is a hierarchical data structure where each node has a value and up to two children (left and right). The BST property states that for every node:
- All values in the left subtree are less than the node's value
- All values in the right subtree are greater than the node's value
- No duplicate values are allowed

## Requirements

### Data Structures
Create the following structures:

**TreeNode**: Represents a single node in the tree
- `value`: integer value stored in the node
- `left`: pointer/reference to left child node (or null)
- `right`: pointer/reference to right child node (or null)

**BST**: Represents the binary search tree
- `root`: pointer/reference to the root node (or null for empty tree)

### Methods to Implement

1. **Insert(value)**: Add a new value to the tree while maintaining BST property
2. **Search(value)**: Return true if value exists in tree, false otherwise
3. **InOrderTraversal()**: Return array/slice of all values in ascending order
4. **Delete(value)**: Remove a value from the tree while maintaining BST property

## Example Usage

```go
// Go example
bst := &BST{}
bst.Insert(5)
bst.Insert(3)
bst.Insert(7)
bst.Insert(2)
bst.Insert(4)

fmt.Println(bst.Search(3))        // true
fmt.Println(bst.Search(6))        // false
fmt.Println(bst.InOrderTraversal()) // [2, 3, 4, 5, 7]

bst.Delete(3)
fmt.Println(bst.InOrderTraversal()) // [2, 4, 5, 7]
```

```typescript
// TypeScript example
const bst = new BST();
bst.insert(5);
bst.insert(3);
bst.insert(7);
bst.insert(2);
bst.insert(4);

console.log(bst.search(3));        // true
console.log(bst.search(6));        // false
console.log(bst.inOrderTraversal()); // [2, 3, 4, 5, 7]

bst.delete(3);
console.log(bst.inOrderTraversal()); // [2, 4, 5, 7]
```

## Constraints

- Values are integers
- No duplicate values allowed in the BST
- Tree can be empty (no nodes)
- Handle all edge cases (empty tree, single node, etc.)
- Maintain BST property after every operation

## Implementation Notes

### Insert Algorithm
- If tree is empty, create root node
- Compare value with current node
- Go left if value is smaller, right if larger
- Insert at first null position found
- Ignore duplicates

### Search Algorithm
- Start from root
- Compare target with current node value
- Navigate left if target is smaller, right if larger
- Return true if found, false if reach null

### Delete Algorithm (Three Cases)
1. **Leaf node**: Simply remove the node
2. **One child**: Replace node with its child
3. **Two children**: Replace with inorder successor (smallest value in right subtree)

### In-Order Traversal
- Visit left subtree
- Process current node
- Visit right subtree
- Results in ascending order for BST

## Test Cases

Your implementation should pass these 10 test cases:

1. **Single Element**: Insert one value, verify search works
2. **Multiple Elements**: Insert several values, verify all can be found
3. **In-Order Traversal**: Verify traversal returns sorted values
4. **Delete Leaf**: Delete a node with no children
5. **Delete One Child**: Delete a node with exactly one child
6. **Delete Two Children**: Delete a node with two children
7. **Delete Root**: Delete the root node
8. **Empty Tree**: Operations on empty tree should not crash
9. **Duplicates**: Inserting duplicates should be ignored
10. **Single Node Tree**: All operations on tree with one node

## Time Complexity Goals

- **Insert**: O(log n) average, O(n) worst case
- **Search**: O(log n) average, O(n) worst case  
- **Delete**: O(log n) average, O(n) worst case
- **In-Order Traversal**: O(n)

## Space Complexity

- O(n) for storing n nodes
- O(h) for recursion stack where h is tree height