# Binary Search Trees

## What is a Binary Search Tree?

A Binary Search Tree (BST) is a hierarchical data structure that extends the concept of a binary tree with an ordering property. In a BST, for every node, all values in the left subtree are less than the node's value, and all values in the right subtree are greater than the node's value. This ordering property enables efficient searching, insertion, and deletion operations.

## Key Properties

- **Binary Tree Structure**: Each node has at most two children (left and right)
- **Ordering Property**: Left subtree values < node value < right subtree values
- **Recursive Structure**: Each subtree is also a valid BST
- **Unique Values**: Typically no duplicate values are allowed (implementation dependent)
- **Dynamic Size**: Can grow and shrink during runtime
- **Self-Organizing**: Structure adapts based on insertion order

## Time Complexity

- **Search**: O(log n) average, O(n) worst case
- **Insertion**: O(log n) average, O(n) worst case
- **Deletion**: O(log n) average, O(n) worst case
- **Traversal**: O(n) for all traversal methods
- **Find Min/Max**: O(log n) average, O(n) worst case

## Go Implementation

### Basic Binary Search Tree

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// BST represents the binary search tree
type BST struct {
    Root *TreeNode
}

// NewBST creates a new empty binary search tree
func NewBST() *BST {
    return &BST{Root: nil}
}

// Insert adds a new value to the BST
func (bst *BST) Insert(value int) {
    bst.Root = bst.insertRecursive(bst.Root, value)
}

func (bst *BST) insertRecursive(node *TreeNode, value int) *TreeNode {
    if node == nil {
        return &TreeNode{Value: value}
    }
    
    if value < node.Value {
        node.Left = bst.insertRecursive(node.Left, value)
    } else if value > node.Value {
        node.Right = bst.insertRecursive(node.Right, value)
    }
    
    return node
}

// Search finds a value in the BST
func (bst *BST) Search(value int) bool {
    return bst.searchRecursive(bst.Root, value)
}

func (bst *BST) searchRecursive(node *TreeNode, value int) bool {
    if node == nil {
        return false
    }
    
    if value == node.Value {
        return true
    } else if value < node.Value {
        return bst.searchRecursive(node.Left, value)
    } else {
        return bst.searchRecursive(node.Right, value)
    }
}

// Delete removes a value from the BST
func (bst *BST) Delete(value int) {
    bst.Root = bst.deleteRecursive(bst.Root, value)
}

func (bst *BST) deleteRecursive(node *TreeNode, value int) *TreeNode {
    if node == nil {
        return nil
    }
    
    if value < node.Value {
        node.Left = bst.deleteRecursive(node.Left, value)
    } else if value > node.Value {
        node.Right = bst.deleteRecursive(node.Right, value)
    } else {
        // Node to delete found
        if node.Left == nil {
            return node.Right
        } else if node.Right == nil {
            return node.Left
        }
        
        // Node has two children
        minNode := bst.findMin(node.Right)
        node.Value = minNode.Value
        node.Right = bst.deleteRecursive(node.Right, minNode.Value)
    }
    
    return node
}

func (bst *BST) findMin(node *TreeNode) *TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}

// InOrderTraversal returns values in sorted order
func (bst *BST) InOrderTraversal() []int {
    var result []int
    bst.inOrderRecursive(bst.Root, &result)
    return result
}

func (bst *BST) inOrderRecursive(node *TreeNode, result *[]int) {
    if node != nil {
        bst.inOrderRecursive(node.Left, result)
        *result = append(*result, node.Value)
        bst.inOrderRecursive(node.Right, result)
    }
}

// Example usage
func main() {
    bst := NewBST()
    
    // Insert values
    values := []int{50, 30, 70, 20, 40, 60, 80}
    for _, value := range values {
        bst.Insert(value)
    }
    
    // Search for values
    fmt.Println("Search 40:", bst.Search(40)) // true
    fmt.Println("Search 25:", bst.Search(25)) // false
    
    // Print sorted values
    fmt.Println("Sorted values:", bst.InOrderTraversal())
    
    // Delete a value
    bst.Delete(30)
    fmt.Println("After deleting 30:", bst.InOrderTraversal())
}
```

## TypeScript Implementation

### Basic Binary Search Tree

```typescript
// TreeNode class represents a node in the binary search tree
class TreeNode<T> {
    value: T;
    left: TreeNode<T> | null;
    right: TreeNode<T> | null;
    
    constructor(value: T) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}

// BST class represents the binary search tree
class BST<T> {
    private root: TreeNode<T> | null;
    private compare: (a: T, b: T) => number;
    
    constructor(compareFunction?: (a: T, b: T) => number) {
        this.root = null;
        this.compare = compareFunction || ((a, b) => {
            if (a < b) return -1;
            if (a > b) return 1;
            return 0;
        });
    }
    
    // Insert a new value into the BST
    insert(value: T): void {
        this.root = this.insertRecursive(this.root, value);
    }
    
    private insertRecursive(node: TreeNode<T> | null, value: T): TreeNode<T> {
        if (node === null) {
            return new TreeNode(value);
        }
        
        const comparison = this.compare(value, node.value);
        if (comparison < 0) {
            node.left = this.insertRecursive(node.left, value);
        } else if (comparison > 0) {
            node.right = this.insertRecursive(node.right, value);
        }
        
        return node;
    }
    
    // Search for a value in the BST
    search(value: T): boolean {
        return this.searchRecursive(this.root, value);
    }
    
    private searchRecursive(node: TreeNode<T> | null, value: T): boolean {
        if (node === null) {
            return false;
        }
        
        const comparison = this.compare(value, node.value);
        if (comparison === 0) {
            return true;
        } else if (comparison < 0) {
            return this.searchRecursive(node.left, value);
        } else {
            return this.searchRecursive(node.right, value);
        }
    }
    
    // Delete a value from the BST
    delete(value: T): void {
        this.root = this.deleteRecursive(this.root, value);
    }
    
    private deleteRecursive(node: TreeNode<T> | null, value: T): TreeNode<T> | null {
        if (node === null) {
            return null;
        }
        
        const comparison = this.compare(value, node.value);
        if (comparison < 0) {
            node.left = this.deleteRecursive(node.left, value);
        } else if (comparison > 0) {
            node.right = this.deleteRecursive(node.right, value);
        } else {
            // Node to delete found
            if (node.left === null) {
                return node.right;
            } else if (node.right === null) {
                return node.left;
            }
            
            // Node has two children
            const minNode = this.findMin(node.right);
            node.value = minNode.value;
            node.right = this.deleteRecursive(node.right, minNode.value);
        }
        
        return node;
    }
    
    private findMin(node: TreeNode<T>): TreeNode<T> {
        while (node.left !== null) {
            node = node.left;
        }
        return node;
    }
    
    // Get all values in sorted order (in-order traversal)
    inOrderTraversal(): T[] {
        const result: T[] = [];
        this.inOrderRecursive(this.root, result);
        return result;
    }
    
    private inOrderRecursive(node: TreeNode<T> | null, result: T[]): void {
        if (node !== null) {
            this.inOrderRecursive(node.left, result);
            result.push(node.value);
            this.inOrderRecursive(node.right, result);
        }
    }
    
    // Check if the tree is empty
    isEmpty(): boolean {
        return this.root === null;
    }
}

// Example usage with type safety
const bst = new BST<number>();

// Insert values
const values = [50, 30, 70, 20, 40, 60, 80];
values.forEach(value => bst.insert(value));

// Search for values
console.log("Search 40:", bst.search(40)); // true
console.log("Search 25:", bst.search(25)); // false

// Print sorted values
console.log("Sorted values:", bst.inOrderTraversal());

// Delete a value
bst.delete(30);
console.log("After deleting 30:", bst.inOrderTraversal());
```

## Common Use Cases

- **Database Indexing**: Fast lookup and range queries in database systems
- **Expression Parsing**: Building and evaluating mathematical expressions
- **File System Organization**: Hierarchical file and directory structures
- **Auto-complete Systems**: Efficient prefix matching and suggestions
- **Priority Systems**: Managing tasks or processes with different priorities
- **Range Queries**: Finding all values within a specific range efficiently
- **Sorted Data Maintenance**: Keeping data sorted while allowing dynamic insertions/deletions
- **Decision Trees**: AI and machine learning applications for classification

## Advantages

- **Efficient Search**: O(log n) average case search time
- **Sorted Output**: In-order traversal gives sorted sequence
- **Dynamic Operations**: Supports efficient insertion and deletion
- **Memory Efficient**: Only stores necessary nodes, no wasted space
- **Range Queries**: Efficient range-based operations
- **Flexible Ordering**: Can work with any comparable data type

## Disadvantages

- **Worst-Case Performance**: Can degrade to O(n) if tree becomes unbalanced
- **No Random Access**: Cannot directly access elements by index
- **Balancing Issues**: May require rebalancing for optimal performance
- **Memory Overhead**: Each node requires additional pointer storage
- **Cache Performance**: Poor spatial locality compared to arrays
- **Complexity**: More complex to implement than linear data structures