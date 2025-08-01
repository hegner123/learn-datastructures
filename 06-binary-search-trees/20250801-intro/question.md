# Binary Search Tree - Node Structure and Basic Properties

## Problem Description

Implement the fundamental building blocks of a Binary Search Tree (BST) by creating the node structure and implementing basic property validation and utility functions. This problem focuses on understanding the core concepts of BST nodes and their properties without implementing complex operations like insertion or deletion.

## Requirements

### Core Components

1. **Node Structure**: Create a node structure that can represent a BST node with:
   - An integer value
   - Left child pointer/reference
   - Right child pointer/reference

2. **CreateNode(value)**: Create a new BST node with the given value

3. **ValidateBST(root)**: Determine if a given binary tree satisfies the BST property:
   - For every node, all values in the left subtree must be less than the node's value
   - For every node, all values in the right subtree must be greater than the node's value
   - Empty trees are considered valid BSTs

4. **FindMin(root)**: Find the minimum value in a BST (leftmost node)

5. **FindMax(root)**: Find the maximum value in a BST (rightmost node)

### Data Structure Requirements

- Handle null/empty tree cases appropriately
- Return -1 for min/max operations on empty trees
- BST validation should work for any binary tree structure
- No duplicate values need to be considered for this intro problem

## Example Usage

### Go Example
```go
// Create nodes
node1 := CreateNode(5)
node2 := CreateNode(3)
node3 := CreateNode(7)

// Manually construct a small BST: 3 <- 5 -> 7
// (You'll implement the node linking in your solution)

// Validate BST property
isValid := ValidateBST(node1) // returns true

// Find min and max
min := FindMin(node1) // returns 3
max := FindMax(node1) // returns 7
```

### TypeScript Example
```typescript
// Create nodes
const node1 = createNode(5);
const node2 = createNode(3);
const node3 = createNode(7);

// Manually construct a small BST: 3 <- 5 -> 7
// (You'll implement the node linking in your solution)

// Validate BST property  
const isValid = validateBST(node1); // returns true

// Find min and max
const min = findMin(node1); // returns 3
const max = findMax(node1); // returns 7
```

## Constraints

- Values are integers in the range [-100, 100]
- Maximum tree depth: 5 levels
- Trees will have at most 10 nodes for testing
- Handle empty tree (null root) cases
- No duplicate values in test cases

## Test Cases

### Test Case 1: Single Node Creation
- **Input**: CreateNode(5)
- **Expected**: Node object with value 5

### Test Case 2: Single Node BST Validation
- **Input**: Tree with single node (value 5), ValidateBST()
- **Expected**: Returns true

### Test Case 3: Valid Small BST
- **Input**: Tree: 3 <- 5 -> 7, ValidateBST()
- **Expected**: Returns true

### Test Case 4: Invalid BST (Left Child Greater)
- **Input**: Tree: 8 <- 5 -> 7 (invalid), ValidateBST()
- **Expected**: Returns false

### Test Case 5: Empty Tree Validation
- **Input**: ValidateBST(null)
- **Expected**: Returns true

### Test Case 6: Find Min Single Node
- **Input**: Tree with single node (value 10), FindMin()
- **Expected**: Returns 10

### Test Case 7: Find Min Empty Tree
- **Input**: FindMin(null)
- **Expected**: Returns -1

### Test Case 8: Find Max Single Node
- **Input**: Tree with single node (value 15), FindMax()
- **Expected**: Returns 15

### Test Case 9: Find Max Empty Tree
- **Input**: FindMax(null)
- **Expected**: Returns -1

### Test Case 10: Find Min/Max Complex Tree
- **Input**: BST with structure: 1 <- 3 <- 5 -> 7 -> 9
- **Expected**: FindMin returns 1, FindMax returns 9

## Implementation Notes

- Focus on understanding the BST property rather than implementing insertion/deletion
- For validation, consider using recursive helper functions with min/max bounds
- The leftmost node contains the minimum value in a BST
- The rightmost node contains the maximum value in a BST
- Empty trees are valid BSTs by definition
- This is an intro-level problem focusing on fundamentals, not optimization

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

Both implementations should pass all 10 test cases to demonstrate correct understanding of BST node structure and basic properties.