# Binary Search Tree - Basic Operations

## Problem Description

Implement a Binary Search Tree (BST) that supports the fundamental operations: insertion, search, and in-order traversal. A Binary Search Tree is a hierarchical data structure where each node has at most two children, and for every node, all values in the left subtree are less than the node's value, and all values in the right subtree are greater than the node's value.

## Requirements

### Core Operations

1. **Insert(value)**: Add a new value to the BST while maintaining the BST property
2. **Search(value)**: Find if a value exists in the BST, return true if found, false otherwise  
3. **InOrderTraversal()**: Return all values in the BST in sorted order (ascending)

### Data Structure Requirements

- Each node should contain an integer value and pointers to left and right children
- The tree should maintain the BST property: left < node < right for all nodes
- Duplicate values should not be inserted (ignore duplicate insertion attempts)
- Handle empty tree cases appropriately

## Example Usage

### Go Example
```go
bst := NewBST()

// Insert values
bst.Insert(50)
bst.Insert(30)
bst.Insert(70)
bst.Insert(20)
bst.Insert(40)

// Search for values
found := bst.Search(30)  // returns true
missing := bst.Search(60) // returns false

// Get sorted values
sorted := bst.InOrderTraversal() // returns [20, 30, 40, 50, 70]
```

### TypeScript Example
```typescript
const bst = new BST();

// Insert values
bst.insert(50);
bst.insert(30);
bst.insert(70);
bst.insert(20);
bst.insert(40);

// Search for values
const found = bst.search(30);  // returns true
const missing = bst.search(60); // returns false

// Get sorted values
const sorted = bst.inOrderTraversal(); // returns [20, 30, 40, 50, 70]
```

## Constraints

- Values are integers in the range [-1000, 1000]
- Maximum tree size: 100 nodes
- All operations should handle empty tree cases
- Insertion of duplicate values should be ignored (no duplicate nodes)
- In-order traversal should return values in ascending order

## Test Cases

### Test Case 1: Single Value Operations
- **Input**: Insert(5), Search(5)
- **Expected**: Search returns true

### Test Case 2: Multiple Value Insertion and Search
- **Input**: Insert(10), Insert(5), Insert(15), Search(5), Search(10), Search(15)
- **Expected**: All searches return true

### Test Case 3: Search Non-existing Value
- **Input**: Insert(10), Insert(20), Search(15)
- **Expected**: Search returns false

### Test Case 4: Simple In-order Traversal
- **Input**: Insert(10), Insert(5), Insert(15), InOrderTraversal()
- **Expected**: Returns [5, 10, 15]

### Test Case 5: Complex Tree Traversal
- **Input**: Insert(50), Insert(30), Insert(20), Insert(40), Insert(70), InOrderTraversal()
- **Expected**: Returns [20, 30, 40, 50, 70]

### Test Case 6: Duplicate Value Handling
- **Input**: Insert(10), Insert(10), Insert(10), InOrderTraversal()
- **Expected**: Returns [10] (single occurrence)

### Test Case 7: Left-skewed Tree
- **Input**: Insert(30), Insert(20), Insert(10), Search(20), InOrderTraversal()
- **Expected**: Search returns true, Traversal returns [10, 20, 30]

### Test Case 8: Right-skewed Tree
- **Input**: Insert(10), Insert(20), Insert(30), Search(25), InOrderTraversal()
- **Expected**: Search returns false, Traversal returns [10, 20, 30]

### Test Case 9: Empty Tree Operations
- **Input**: Search(10), InOrderTraversal()
- **Expected**: Search returns false, Traversal returns []

### Test Case 10: Large Tree Operations
- **Input**: Insert(50,25,75,12,37,62,87), Search(37), Search(100), InOrderTraversal()
- **Expected**: First search true, second false, Traversal returns [12,25,37,50,62,75,87]

## Implementation Notes

- Use recursive approaches for cleaner code structure
- Consider implementing helper methods for recursive operations
- Ensure proper handling of null/nil nodes
- The in-order traversal (Left-Root-Right) naturally produces sorted output
- Focus on correctness first, optimization is not required for this beginner-level problem

## Testing

### Go Testing
```bash
cd go/
make test
```

### TypeScript Testing
```bash
cd ts/
npm test
```

Both implementations should pass all 10 test cases to demonstrate correct BST functionality.