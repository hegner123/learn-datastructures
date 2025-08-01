# Binary Tree Insertion Methods

## Overview

Binary trees support multiple insertion strategies depending on the desired tree structure and use case. Unlike Binary Search Trees, binary trees don't have inherent ordering rules, so insertion methods focus on structural requirements.

## Insertion Methods

### 1. Level-Order Insertion (Complete Tree)

Creates a complete binary tree by filling levels left-to-right.

```go
// Insert adds a new node to the binary tree (level order insertion)
func (bt *BinaryTree) Insert(data int) {
    newNode := &TreeNode{data: data}
    
    if bt.root == nil {
        bt.root = newNode
        return
    }
    
    // Use queue for level order insertion
    queue := []*TreeNode{bt.root}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        if current.left == nil {
            current.left = newNode
            return
        } else {
            queue = append(queue, current.left)
        }
        
        if current.right == nil {
            current.right = newNode
            return
        } else {
            queue = append(queue, current.right)
        }
    }
}
```

### 2. Manual Position Insertion

Insert at specific parent node positions.

```go
// InsertAtPosition inserts a node as left or right child of a specific parent
func (bt *BinaryTree) InsertAtPosition(parentData int, data int, isLeft bool) bool {
    if bt.root == nil {
        return false
    }
    
    parent := bt.findNode(bt.root, parentData)
    if parent == nil {
        return false
    }
    
    newNode := &TreeNode{data: data}
    
    if isLeft {
        if parent.left == nil {
            parent.left = newNode
            return true
        }
    } else {
        if parent.right == nil {
            parent.right = newNode
            return true
        }
    }
    
    return false // Position already occupied
}

// Helper function to find a node
func (bt *BinaryTree) findNode(node *TreeNode, data int) *TreeNode {
    if node == nil {
        return nil
    }
    
    if node.data == data {
        return node
    }
    
    left := bt.findNode(node.left, data)
    if left != nil {
        return left
    }
    
    return bt.findNode(node.right, data)
}
```

### 3. Array to Tree Construction

Build tree from array representation (with null values).

```go
// BuildFromArray constructs a binary tree from array representation
// nil values in array represent missing nodes
func (bt *BinaryTree) BuildFromArray(arr []*int) {
    if len(arr) == 0 || arr[0] == nil {
        bt.root = nil
        return
    }
    
    bt.root = &TreeNode{data: *arr[0]}
    queue := []*TreeNode{bt.root}
    i := 1
    
    for len(queue) > 0 && i < len(arr) {
        current := queue[0]
        queue = queue[1:]
        
        // Left child
        if i < len(arr) && arr[i] != nil {
            current.left = &TreeNode{data: *arr[i]}
            queue = append(queue, current.left)
        }
        i++
        
        // Right child
        if i < len(arr) && arr[i] != nil {
            current.right = &TreeNode{data: *arr[i]}
            queue = append(queue, current.right)
        }
        i++
    }
}

// Helper function to create int pointer
func intPtr(x int) *int {
    return &x
}
```

### 4. Expression Tree Construction

Build tree from postfix expressions.

```go
// ExpressionNode for mathematical expressions
type ExpressionNode struct {
    value string
    left  *ExpressionNode
    right *ExpressionNode
}

// BuildExpressionTree creates a binary tree from postfix expression
func BuildExpressionTree(postfix []string) *ExpressionNode {
    if len(postfix) == 0 {
        return nil
    }
    
    stack := []*ExpressionNode{}
    
    for _, token := range postfix {
        node := &ExpressionNode{value: token}
        
        if isOperator(token) {
            if len(stack) >= 2 {
                node.right = stack[len(stack)-1]
                node.left = stack[len(stack)-2]
                stack = stack[:len(stack)-2]
            }
        }
        
        stack = append(stack, node)
    }
    
    if len(stack) > 0 {
        return stack[0]
    }
    return nil
}

func isOperator(token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/"
}
```

## TypeScript Implementation

### Level-Order Insertion

```typescript
class BinaryTree<T> {
    private root: TreeNode<T> | null;
    
    constructor() {
        this.root = null;
    }
    
    // Insert node using level order insertion
    public insert(data: T): void {
        const newNode = new TreeNode<T>(data);
        
        if (!this.root) {
            this.root = newNode;
            return;
        }
        
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const current = queue.shift()!;
            
            if (!current.left) {
                current.left = newNode;
                return;
            } else {
                queue.push(current.left);
            }
            
            if (!current.right) {
                current.right = newNode;
                return;
            } else {
                queue.push(current.right);
            }
        }
    }
    
    // Insert at specific position
    public insertAtPosition(parentData: T, data: T, isLeft: boolean): boolean {
        if (!this.root) return false;
        
        const parent = this.findNode(this.root, parentData);
        if (!parent) return false;
        
        const newNode = new TreeNode<T>(data);
        
        if (isLeft) {
            if (!parent.left) {
                parent.left = newNode;
                return true;
            }
        } else {
            if (!parent.right) {
                parent.right = newNode;
                return true;
            }
        }
        
        return false; // Position already occupied
    }
    
    private findNode(node: TreeNode<T> | null, data: T): TreeNode<T> | null {
        if (!node) return null;
        if (node.data === data) return node;
        
        const left = this.findNode(node.left, data);
        if (left) return left;
        
        return this.findNode(node.right, data);
    }
    
    // Build from array representation
    public buildFromArray(arr: (T | null)[]): void {
        if (arr.length === 0 || arr[0] === null) {
            this.root = null;
            return;
        }
        
        this.root = new TreeNode<T>(arr[0]!);
        const queue: TreeNode<T>[] = [this.root];
        let i = 1;
        
        while (queue.length > 0 && i < arr.length) {
            const current = queue.shift()!;
            
            // Left child
            if (i < arr.length && arr[i] !== null) {
                current.left = new TreeNode<T>(arr[i]!);
                queue.push(current.left);
            }
            i++;
            
            // Right child
            if (i < arr.length && arr[i] !== null) {
                current.right = new TreeNode<T>(arr[i]!);
                queue.push(current.right);
            }
            i++;
        }
    }
    
    // Insert multiple nodes
    public insertMany(data: T[]): void {
        data.forEach(item => this.insert(item));
    }
    
    // Replace node value
    public replaceNode(oldData: T, newData: T): boolean {
        const node = this.findNode(this.root, oldData);
        if (node) {
            node.data = newData;
            return true;
        }
        return false;
    }
}
```

### Expression Tree (TypeScript)

```typescript
class ExpressionNode {
    public value: string;
    public left: ExpressionNode | null;
    public right: ExpressionNode | null;
    
    constructor(value: string) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}

class ExpressionTree {
    private root: ExpressionNode | null;
    
    constructor() {
        this.root = null;
    }
    
    // Build expression tree from postfix
    public buildFromPostfix(postfix: string[]): void {
        const stack: ExpressionNode[] = [];
        
        for (const token of postfix) {
            const node = new ExpressionNode(token);
            
            if (this.isOperator(token)) {
                if (stack.length >= 2) {
                    node.right = stack.pop()!;
                    node.left = stack.pop()!;
                }
            }
            
            stack.push(node);
        }
        
        this.root = stack.length > 0 ? stack[0] : null;
    }
    
    private isOperator(token: string): boolean {
        return ['+', '-', '*', '/'].includes(token);
    }
    
    public getRoot(): ExpressionNode | null {
        return this.root;
    }
}
```

## Usage Examples

### Go Examples

```go
func main() {
    bt := NewBinaryTree()
    
    // Level order insertion
    elements := []int{1, 2, 3, 4, 5, 6, 7}
    for _, element := range elements {
        bt.Insert(element)
    }
    
    // Manual position insertion
    bt2 := NewBinaryTree()
    bt2.Insert(1) // Root
    bt2.InsertAtPosition(1, 2, true)  // Left child of 1
    bt2.InsertAtPosition(1, 3, false) // Right child of 1
    
    // Array construction
    bt3 := NewBinaryTree()
    arr := []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(4), nil, intPtr(5)}
    bt3.BuildFromArray(arr)
    
    // Expression tree
    postfix := []string{"3", "4", "+", "2", "*", "7", "/"}
    expTree := BuildExpressionTree(postfix)
    
    fmt.Println("Trees created successfully")
}
```

### TypeScript Examples

```typescript
// Level order insertion
const bt1 = new BinaryTree<number>();
const elements: number[] = [1, 2, 3, 4, 5, 6, 7];
bt1.insertMany(elements);

// Manual position insertion
const bt2 = new BinaryTree<number>();
bt2.insert(1); // Root
bt2.insertAtPosition(1, 2, true);  // Left child of 1
bt2.insertAtPosition(1, 3, false); // Right child of 1

// Array construction
const bt3 = new BinaryTree<number>();
const arr: (number | null)[] = [1, 2, 3, null, 4, null, 5];
bt3.buildFromArray(arr);

// Expression tree
const expTree = new ExpressionTree();
const postfix: string[] = ["3", "4", "+", "2", "*", "7", "/"];
expTree.buildFromPostfix(postfix);

console.log("All trees created successfully");
```

## Insertion Strategy Comparison

| Method | Use Case | Time Complexity | Maintains Structure |
|--------|----------|----------------|-------------------|
| Level-Order | Complete trees, heaps | O(n) | Yes - Complete |
| Manual Position | Custom structures | O(n) | No - Custom |
| Array Construction | Known structure | O(n) | Yes - As specified |
| Expression Tree | Math expressions | O(n) | Yes - Expression rules |

## Best Practices

1. **Choose insertion method based on use case**:
   - Level-order for balanced, complete trees
   - Manual position for specific structures
   - Array construction for known layouts

2. **Validate insertion constraints**:
   - Check if position is available
   - Handle null/empty cases
   - Maintain tree invariants

3. **Consider memory efficiency**:
   - Level-order creates compact trees
   - Manual insertion can create sparse trees

4. **Error handling**:
   - Return success/failure indicators
   - Handle edge cases gracefully