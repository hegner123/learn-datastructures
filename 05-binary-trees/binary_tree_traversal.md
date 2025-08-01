# Binary Tree Traversal Methods

## Overview

Tree traversal is the process of visiting all nodes in a tree data structure in a specific order. Binary trees support multiple traversal methods, each useful for different applications.

## Traversal Types

### 1. Depth-First Traversals (DFS)

These traversals use recursion or explicit stacks to visit nodes deeply before exploring siblings.

#### Inorder Traversal (Left → Root → Right)

Most common for Binary Search Trees as it visits nodes in sorted order.

```go
// InorderTraversal performs inorder traversal (Left → Root → Right)
func (bt *BinaryTree) InorderTraversal() []int {
    var result []int
    bt.inorderHelper(bt.root, &result)
    return result
}

func (bt *BinaryTree) inorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    
    bt.inorderHelper(node.left, result)   // Left
    *result = append(*result, node.data)  // Root
    bt.inorderHelper(node.right, result)  // Right
}

// Iterative inorder traversal using stack
func (bt *BinaryTree) InorderIterative() []int {
    var result []int
    var stack []*TreeNode
    current := bt.root
    
    for current != nil || len(stack) > 0 {
        // Go to leftmost node
        for current != nil {
            stack = append(stack, current)
            current = current.left
        }
        
        // Pop and process
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.data)
        
        // Move to right subtree
        current = current.right
    }
    
    return result
}
```

#### Preorder Traversal (Root → Left → Right)

Useful for copying/cloning trees and prefix expressions.

```go
// PreorderTraversal performs preorder traversal (Root → Left → Right)
func (bt *BinaryTree) PreorderTraversal() []int {
    var result []int
    bt.preorderHelper(bt.root, &result)
    return result
}

func (bt *BinaryTree) preorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    
    *result = append(*result, node.data)  // Root
    bt.preorderHelper(node.left, result)  // Left
    bt.preorderHelper(node.right, result) // Right
}

// Iterative preorder traversal
func (bt *BinaryTree) PreorderIterative() []int {
    if bt.root == nil {
        return []int{}
    }
    
    var result []int
    stack := []*TreeNode{bt.root}
    
    for len(stack) > 0 {
        // Pop node
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        result = append(result, node.data)
        
        // Push right first (so left is processed first)
        if node.right != nil {
            stack = append(stack, node.right)
        }
        if node.left != nil {
            stack = append(stack, node.left)
        }
    }
    
    return result
}
```

#### Postorder Traversal (Left → Right → Root)

Useful for deleting trees and postfix expressions.

```go
// PostorderTraversal performs postorder traversal (Left → Right → Root)
func (bt *BinaryTree) PostorderTraversal() []int {
    var result []int
    bt.postorderHelper(bt.root, &result)
    return result
}

func (bt *BinaryTree) postorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    
    bt.postorderHelper(node.left, result)  // Left
    bt.postorderHelper(node.right, result) // Right
    *result = append(*result, node.data)   // Root
}

// Iterative postorder traversal (more complex)
func (bt *BinaryTree) PostorderIterative() []int {
    if bt.root == nil {
        return []int{}
    }
    
    var result []int
    var stack []*TreeNode
    var lastVisited *TreeNode
    current := bt.root
    
    for len(stack) > 0 || current != nil {
        if current != nil {
            stack = append(stack, current)
            current = current.left
        } else {
            peekNode := stack[len(stack)-1]
            
            // If right child exists and hasn't been processed yet
            if peekNode.right != nil && lastVisited != peekNode.right {
                current = peekNode.right
            } else {
                result = append(result, peekNode.data)
                lastVisited = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            }
        }
    }
    
    return result
}
```

### 2. Breadth-First Traversal (BFS)

#### Level Order Traversal

Visits nodes level by level from left to right.

```go
// LevelOrderTraversal performs level order traversal (BFS)
func (bt *BinaryTree) LevelOrderTraversal() []int {
    if bt.root == nil {
        return []int{}
    }
    
    var result []int
    queue := []*TreeNode{bt.root}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        result = append(result, current.data)
        
        if current.left != nil {
            queue = append(queue, current.left)
        }
        
        if current.right != nil {
            queue = append(queue, current.right)
        }
    }
    
    return result
}

// LevelOrderByLevel returns traversal grouped by levels
func (bt *BinaryTree) LevelOrderByLevel() [][]int {
    if bt.root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{bt.root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        var currentLevel []int
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            currentLevel = append(currentLevel, node.data)
            
            if node.left != nil {
                queue = append(queue, node.left)
            }
            if node.right != nil {
                queue = append(queue, node.right)
            }
        }
        
        result = append(result, currentLevel)
    }
    
    return result
}
```

### 3. Specialized Traversals

#### Reverse Level Order

```go
// ReverseLevelOrder traverses from bottom to top
func (bt *BinaryTree) ReverseLevelOrder() []int {
    if bt.root == nil {
        return []int{}
    }
    
    var result []int
    queue := []*TreeNode{bt.root}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        result = append(result, current.data)
        
        // Add right first, then left (reverse of normal level order)
        if current.right != nil {
            queue = append(queue, current.right)
        }
        if current.left != nil {
            queue = append(queue, current.left)
        }
    }
    
    // Reverse the result
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return result
}
```

#### Zigzag Traversal

```go
// ZigzagTraversal alternates direction at each level
func (bt *BinaryTree) ZigzagTraversal() []int {
    if bt.root == nil {
        return []int{}
    }
    
    var result []int
    queue := []*TreeNode{bt.root}
    leftToRight := true
    
    for len(queue) > 0 {
        levelSize := len(queue)
        var currentLevel []int
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            currentLevel = append(currentLevel, node.data)
            
            if node.left != nil {
                queue = append(queue, node.left)
            }
            if node.right != nil {
                queue = append(queue, node.right)
            }
        }
        
        // Reverse level if going right to left
        if !leftToRight {
            for i, j := 0, len(currentLevel)-1; i < j; i, j = i+1, j-1 {
                currentLevel[i], currentLevel[j] = currentLevel[j], currentLevel[i]
            }
        }
        
        result = append(result, currentLevel...)
        leftToRight = !leftToRight
    }
    
    return result
}
```

## TypeScript Implementation

### Complete Traversal Implementation

```typescript
class BinaryTree<T> {
    private root: TreeNode<T> | null;
    
    constructor() {
        this.root = null;
    }
    
    // Inorder traversal (Left → Root → Right)
    public inorderTraversal(): T[] {
        const result: T[] = [];
        this.inorderHelper(this.root, result);
        return result;
    }
    
    private inorderHelper(node: TreeNode<T> | null, result: T[]): void {
        if (!node) return;
        
        this.inorderHelper(node.left, result);
        result.push(node.data);
        this.inorderHelper(node.right, result);
    }
    
    // Iterative inorder
    public inorderIterative(): T[] {
        const result: T[] = [];
        const stack: TreeNode<T>[] = [];
        let current = this.root;
        
        while (current || stack.length > 0) {
            while (current) {
                stack.push(current);
                current = current.left;
            }
            
            current = stack.pop()!;
            result.push(current.data);
            current = current.right;
        }
        
        return result;
    }
    
    // Preorder traversal (Root → Left → Right)
    public preorderTraversal(): T[] {
        const result: T[] = [];
        this.preorderHelper(this.root, result);
        return result;
    }
    
    private preorderHelper(node: TreeNode<T> | null, result: T[]): void {
        if (!node) return;
        
        result.push(node.data);
        this.preorderHelper(node.left, result);
        this.preorderHelper(node.right, result);
    }
    
    // Iterative preorder
    public preorderIterative(): T[] {
        if (!this.root) return [];
        
        const result: T[] = [];
        const stack: TreeNode<T>[] = [this.root];
        
        while (stack.length > 0) {
            const node = stack.pop()!;
            result.push(node.data);
            
            if (node.right) stack.push(node.right);
            if (node.left) stack.push(node.left);
        }
        
        return result;
    }
    
    // Postorder traversal (Left → Right → Root)
    public postorderTraversal(): T[] {
        const result: T[] = [];
        this.postorderHelper(this.root, result);
        return result;
    }
    
    private postorderHelper(node: TreeNode<T> | null, result: T[]): void {
        if (!node) return;
        
        this.postorderHelper(node.left, result);
        this.postorderHelper(node.right, result);
        result.push(node.data);
    }
    
    // Iterative postorder
    public postorderIterative(): T[] {
        if (!this.root) return [];
        
        const result: T[] = [];
        const stack: TreeNode<T>[] = [];
        let lastVisited: TreeNode<T> | null = null;
        let current: TreeNode<T> | null = this.root;
        
        while (stack.length > 0 || current) {
            if (current) {
                stack.push(current);
                current = current.left;
            } else {
                const peekNode = stack[stack.length - 1];
                
                if (peekNode.right && lastVisited !== peekNode.right) {
                    current = peekNode.right;
                } else {
                    result.push(peekNode.data);
                    lastVisited = stack.pop()!;
                }
            }
        }
        
        return result;
    }
    
    // Level order traversal (BFS)
    public levelOrderTraversal(): T[] {
        if (!this.root) return [];
        
        const result: T[] = [];
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const current = queue.shift()!;
            result.push(current.data);
            
            if (current.left) queue.push(current.left);
            if (current.right) queue.push(current.right);
        }
        
        return result;
    }
    
    // Level order by level
    public levelOrderByLevel(): T[][] {
        if (!this.root) return [];
        
        const result: T[][] = [];
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const levelSize = queue.length;
            const currentLevel: T[] = [];
            
            for (let i = 0; i < levelSize; i++) {
                const node = queue.shift()!;
                currentLevel.push(node.data);
                
                if (node.left) queue.push(node.left);
                if (node.right) queue.push(node.right);
            }
            
            result.push(currentLevel);
        }
        
        return result;
    }
    
    // Zigzag traversal
    public zigzagTraversal(): T[] {
        if (!this.root) return [];
        
        const result: T[] = [];
        const queue: TreeNode<T>[] = [this.root];
        let leftToRight = true;
        
        while (queue.length > 0) {
            const levelSize = queue.length;
            const currentLevel: T[] = [];
            
            for (let i = 0; i < levelSize; i++) {
                const node = queue.shift()!;
                currentLevel.push(node.data);
                
                if (node.left) queue.push(node.left);
                if (node.right) queue.push(node.right);
            }
            
            if (!leftToRight) {
                currentLevel.reverse();
            }
            
            result.push(...currentLevel);
            leftToRight = !leftToRight;
        }
        
        return result;
    }
    
    // Reverse level order
    public reverseLevelOrder(): T[] {
        if (!this.root) return [];
        
        const result: T[] = [];
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const current = queue.shift()!;
            result.push(current.data);
            
            if (current.right) queue.push(current.right);
            if (current.left) queue.push(current.left);
        }
        
        return result.reverse();
    }
    
    // Morris traversal (inorder without extra space)
    public morrisInorder(): T[] {
        const result: T[] = [];
        let current = this.root;
        
        while (current) {
            if (!current.left) {
                result.push(current.data);
                current = current.right;
            } else {
                // Find inorder predecessor
                let predecessor = current.left;
                while (predecessor.right && predecessor.right !== current) {
                    predecessor = predecessor.right;
                }
                
                if (!predecessor.right) {
                    // Make current the right child of its inorder predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Revert the changes
                    predecessor.right = null;
                    result.push(current.data);
                    current = current.right;
                }
            }
        }
        
        return result;
    }
    
    // Vertical order traversal
    public verticalOrder(): T[][] {
        if (!this.root) return [];
        
        const columnMap = new Map<number, T[]>();
        const queue: Array<[TreeNode<T>, number]> = [[this.root, 0]];
        let minColumn = 0;
        let maxColumn = 0;
        
        while (queue.length > 0) {
            const [node, column] = queue.shift()!;
            
            if (!columnMap.has(column)) {
                columnMap.set(column, []);
            }
            columnMap.get(column)!.push(node.data);
            
            minColumn = Math.min(minColumn, column);
            maxColumn = Math.max(maxColumn, column);
            
            if (node.left) {
                queue.push([node.left, column - 1]);
            }
            if (node.right) {
                queue.push([node.right, column + 1]);
            }
        }
        
        const result: T[][] = [];
        for (let i = minColumn; i <= maxColumn; i++) {
            if (columnMap.has(i)) {
                result.push(columnMap.get(i)!);
            }
        }
        
        return result;
    }
}
```

## Advanced Traversal Applications

### Find All Paths from Root to Leaves

```typescript
public findAllPaths(): T[][] {
    const result: T[][] = [];
    const currentPath: T[] = [];
    this.findPathsHelper(this.root, currentPath, result);
    return result;
}

private findPathsHelper(node: TreeNode<T> | null, currentPath: T[], result: T[][]): void {
    if (!node) return;
    
    currentPath.push(node.data);
    
    // If leaf node, add path to result
    if (!node.left && !node.right) {
        result.push([...currentPath]);
    } else {
        this.findPathsHelper(node.left, currentPath, result);
        this.findPathsHelper(node.right, currentPath, result);
    }
    
    currentPath.pop(); // Backtrack
}
```

### Boundary Traversal

```typescript
public boundaryTraversal(): T[] {
    if (!this.root) return [];
    
    const result: T[] = [this.root.data];
    
    // Left boundary (excluding leaf)
    this.leftBoundary(this.root.left, result);
    
    // Leaf nodes
    this.leafNodes(this.root, result);
    
    // Right boundary (excluding leaf, in reverse)
    const rightBoundary: T[] = [];
    this.rightBoundary(this.root.right, rightBoundary);
    result.push(...rightBoundary.reverse());
    
    return result;
}

private leftBoundary(node: TreeNode<T> | null, result: T[]): void {
    if (!node || (!node.left && !node.right)) return;
    
    result.push(node.data);
    
    if (node.left) {
        this.leftBoundary(node.left, result);
    } else {
        this.leftBoundary(node.right, result);
    }
}

private rightBoundary(node: TreeNode<T> | null, result: T[]): void {
    if (!node || (!node.left && !node.right)) return;
    
    result.push(node.data);
    
    if (node.right) {
        this.rightBoundary(node.right, result);
    } else {
        this.rightBoundary(node.left, result);
    }
}

private leafNodes(node: TreeNode<T> | null, result: T[]): void {
    if (!node) return;
    
    if (!node.left && !node.right) {
        result.push(node.data);
        return;
    }
    
    this.leafNodes(node.left, result);
    this.leafNodes(node.right, result);
}
```

## Usage Examples

### Go Examples

```go
func main() {
    bt := NewBinaryTree()
    
    // Build test tree: 1,2,3,4,5,6,7
    elements := []int{1, 2, 3, 4, 5, 6, 7}
    for _, element := range elements {
        bt.Insert(element)
    }
    
    fmt.Println("Inorder:", bt.InorderTraversal())
    fmt.Println("Preorder:", bt.PreorderTraversal())
    fmt.Println("Postorder:", bt.PostorderTraversal())
    fmt.Println("Level Order:", bt.LevelOrderTraversal())
    fmt.Println("Zigzag:", bt.ZigzagTraversal())
    fmt.Println("Reverse Level:", bt.ReverseLevelOrder())
    
    // By level
    levels := bt.LevelOrderByLevel()
    fmt.Println("By Level:", levels)
}
```

### TypeScript Examples

```typescript
const bt = new BinaryTree<number>();

// Build test tree
[1, 2, 3, 4, 5, 6, 7].forEach(x => bt.insert(x));

console.log("Inorder (Recursive):", bt.inorderTraversal());
console.log("Inorder (Iterative):", bt.inorderIterative());
console.log("Inorder (Morris):", bt.morrisInorder());

console.log("Preorder (Recursive):", bt.preorderTraversal());
console.log("Preorder (Iterative):", bt.preorderIterative());

console.log("Postorder (Recursive):", bt.postorderTraversal());
console.log("Postorder (Iterative):", bt.postorderIterative());

console.log("Level Order:", bt.levelOrderTraversal());
console.log("Level Order by Level:", bt.levelOrderByLevel());
console.log("Zigzag:", bt.zigzagTraversal());
console.log("Reverse Level Order:", bt.reverseLevelOrder());

console.log("All Paths:", bt.findAllPaths());
console.log("Boundary:", bt.boundaryTraversal());
console.log("Vertical Order:", bt.verticalOrder());
```

## Traversal Comparison

| Traversal | Order | Use Case | Time | Space |
|-----------|-------|----------|------|-------|
| Inorder | L-Root-R | BST sorted order | O(n) | O(h) |
| Preorder | Root-L-R | Tree copying | O(n) | O(h) |
| Postorder | L-R-Root | Tree deletion | O(n) | O(h) |
| Level Order | By levels | BFS, printing | O(n) | O(w) |
| Morris | Inorder | Space-optimal | O(n) | O(1) |

Where h = height, w = maximum width of tree

## Applications

- **Inorder**: Expression evaluation, BST validation
- **Preorder**: Tree serialization, prefix expressions
- **Postorder**: Directory size calculation, postfix expressions
- **Level Order**: Tree printing, shortest path in unweighted trees
- **Zigzag**: Spiral printing, alternating processing
- **Vertical**: Column-wise processing, coordinate-based operations