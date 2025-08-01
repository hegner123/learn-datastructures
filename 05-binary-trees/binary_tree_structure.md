# Binary Tree Structure

## What is a Binary Tree?

A binary tree is a hierarchical data structure where each node has at most two children, commonly referred to as the left child and right child. It's a tree data structure where every node has zero, one, or two children.

## Key Properties

- **Root**: The topmost node in the tree
- **Leaf**: A node with no children
- **Height**: The length of the longest path from root to leaf
- **Depth**: The length of the path from root to a particular node
- **Level**: All nodes at the same depth are at the same level

## Types of Binary Trees

- **Full Binary Tree**: Every node has either 0 or 2 children
- **Complete Binary Tree**: All levels are filled except possibly the last level
- **Perfect Binary Tree**: All internal nodes have 2 children and all leaves are at the same level
- **Balanced Binary Tree**: Height difference between left and right subtrees is at most 1

## Time Complexity

- **Search**: O(n) - worst case, O(log n) - balanced tree
- **Insertion**: O(n) - worst case, O(log n) - balanced tree
- **Deletion**: O(n) - worst case, O(log n) - balanced tree
- **Traversal**: O(n)

## Go Implementation

### Basic Node Structure

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree
type TreeNode struct {
    data  int
    left  *TreeNode
    right *TreeNode
}

// BinaryTree represents the binary tree structure
type BinaryTree struct {
    root *TreeNode
}

// NewBinaryTree creates a new empty binary tree
func NewBinaryTree() *BinaryTree {
    return &BinaryTree{root: nil}
}

// NewTreeNode creates a new tree node
func NewTreeNode(data int) *TreeNode {
    return &TreeNode{
        data:  data,
        left:  nil,
        right: nil,
    }
}

// GetRoot returns the root node
func (bt *BinaryTree) GetRoot() *TreeNode {
    return bt.root
}

// SetRoot sets the root node
func (bt *BinaryTree) SetRoot(node *TreeNode) {
    bt.root = node
}

// IsEmpty checks if the tree is empty
func (bt *BinaryTree) IsEmpty() bool {
    return bt.root == nil
}

func main() {
    // Create a new binary tree
    bt := NewBinaryTree()
    
    // Create nodes manually
    root := NewTreeNode(1)
    root.left = NewTreeNode(2)
    root.right = NewTreeNode(3)
    root.left.left = NewTreeNode(4)
    root.left.right = NewTreeNode(5)
    
    bt.SetRoot(root)
    
    fmt.Printf("Tree is empty: %t\n", bt.IsEmpty())
    fmt.Printf("Root data: %d\n", bt.GetRoot().data)
}
```

## TypeScript Implementation

### Basic Node Structure

```typescript
// TreeNode class with generic type support
class TreeNode<T> {
    public data: T;
    public left: TreeNode<T> | null;
    public right: TreeNode<T> | null;
    
    constructor(data: T) {
        this.data = data;
        this.left = null;
        this.right = null;
    }
}

// BinaryTree class with generic type support
class BinaryTree<T> {
    private root: TreeNode<T> | null;
    
    constructor() {
        this.root = null;
    }
    
    // Get root node
    public getRoot(): TreeNode<T> | null {
        return this.root;
    }
    
    // Set root node
    public setRoot(node: TreeNode<T> | null): void {
        this.root = node;
    }
    
    // Check if tree is empty
    public isEmpty(): boolean {
        return this.root === null;
    }
}

// Example usage
const bt = new BinaryTree<number>();

// Create nodes manually
const root = new TreeNode<number>(1);
root.left = new TreeNode<number>(2);
root.right = new TreeNode<number>(3);
root.left.left = new TreeNode<number>(4);
root.left.right = new TreeNode<number>(5);

bt.setRoot(root);

console.log("Tree is empty:", bt.isEmpty());
console.log("Root data:", bt.getRoot()?.data);
```

## Tree Visualization

### Simple Tree Structure

```
        1
       / \
      2   3
     / \
    4   5
```

### Visual Tree Printer (TypeScript)

```typescript
class BinaryTree<T> {
    // ... previous code ...
    
    // Print tree structure
    public printTree(): void {
        if (!this.root) {
            console.log("Tree is empty");
            return;
        }
        
        this.printTreeHelper(this.root, "", true);
    }
    
    private printTreeHelper(node: TreeNode<T> | null, prefix: string, isLast: boolean): void {
        if (!node) return;
        
        console.log(prefix + (isLast ? "└── " : "├── ") + node.data);
        
        const children: TreeNode<T>[] = [];
        if (node.left) children.push(node.left);
        if (node.right) children.push(node.right);
        
        children.forEach((child, index) => {
            const isChildLast = index === children.length - 1;
            const childPrefix = prefix + (isLast ? "    " : "│   ");
            
            if (child === node.left && node.right) {
                this.printTreeHelper(child, childPrefix, false);
            } else {
                this.printTreeHelper(child, childPrefix, isChildLast);
            }
        });
    }
}
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