# Binary Tree Deletion Methods

## Overview

Binary tree deletion is more complex than insertion since removing nodes affects tree structure. Different strategies exist based on what should happen to child nodes when a parent is deleted.

## Deletion Strategies

### 1. Delete Leaf Nodes Only

Simplest deletion - only removes nodes with no children.

```go
// DeleteLeaf removes a leaf node (node with no children)
func (bt *BinaryTree) DeleteLeaf(data int) bool {
    if bt.root == nil {
        return false
    }
    
    // Special case: root is leaf
    if bt.root.data == data && bt.root.left == nil && bt.root.right == nil {
        bt.root = nil
        return true
    }
    
    return bt.deleteLeafHelper(bt.root, data)
}

func (bt *BinaryTree) deleteLeafHelper(node *TreeNode, data int) bool {
    if node == nil {
        return false
    }
    
    // Check left child
    if node.left != nil && node.left.data == data {
        if node.left.left == nil && node.left.right == nil {
            node.left = nil
            return true
        }
    }
    
    // Check right child
    if node.right != nil && node.right.data == data {
        if node.right.left == nil && node.right.right == nil {
            node.right = nil
            return true
        }
    }
    
    // Recursively search in subtrees
    return bt.deleteLeafHelper(node.left, data) || 
           bt.deleteLeafHelper(node.right, data)
}
```

### 2. Delete Node and Replace with Child

When deleting a node with one child, replace it with that child.

```go
// DeleteAndPromoteChild removes node and promotes its child
func (bt *BinaryTree) DeleteAndPromoteChild(data int) bool {
    if bt.root == nil {
        return false
    }
    
    // Special case: deleting root
    if bt.root.data == data {
        if bt.root.left == nil && bt.root.right == nil {
            bt.root = nil
        } else if bt.root.left == nil {
            bt.root = bt.root.right
        } else if bt.root.right == nil {
            bt.root = bt.root.left
        } else {
            // Root has two children - need different strategy
            return false
        }
        return true
    }
    
    return bt.deleteAndPromoteHelper(bt.root, data)
}

func (bt *BinaryTree) deleteAndPromoteHelper(node *TreeNode, data int) bool {
    if node == nil {
        return false
    }
    
    // Check left child
    if node.left != nil && node.left.data == data {
        target := node.left
        if target.left == nil && target.right == nil {
            // Leaf node
            node.left = nil
        } else if target.left == nil {
            // Only right child
            node.left = target.right
        } else if target.right == nil {
            // Only left child
            node.left = target.left
        } else {
            // Has two children - cannot promote
            return false
        }
        return true
    }
    
    // Check right child
    if node.right != nil && node.right.data == data {
        target := node.right
        if target.left == nil && target.right == nil {
            // Leaf node
            node.right = nil
        } else if target.left == nil {
            // Only right child
            node.right = target.right
        } else if target.right == nil {
            // Only left child
            node.right = target.left
        } else {
            // Has two children - cannot promote
            return false
        }
        return true
    }
    
    // Recursively search in subtrees
    return bt.deleteAndPromoteHelper(node.left, data) || 
           bt.deleteAndPromoteHelper(node.right, data)
}
```

### 3. Delete Subtree (Node and All Descendants)

Removes entire subtree rooted at the target node.

```go
// DeleteSubtree removes the node and all its descendants
func (bt *BinaryTree) DeleteSubtree(data int) bool {
    if bt.root == nil {
        return false
    }
    
    // Special case: deleting root (clears entire tree)
    if bt.root.data == data {
        bt.root = nil
        return true
    }
    
    return bt.deleteSubtreeHelper(bt.root, data)
}

func (bt *BinaryTree) deleteSubtreeHelper(node *TreeNode, data int) bool {
    if node == nil {
        return false
    }
    
    // Check left child
    if node.left != nil && node.left.data == data {
        node.left = nil // This removes entire left subtree
        return true
    }
    
    // Check right child
    if node.right != nil && node.right.data == data {
        node.right = nil // This removes entire right subtree
        return true
    }
    
    // Recursively search in subtrees
    return bt.deleteSubtreeHelper(node.left, data) || 
           bt.deleteSubtreeHelper(node.right, data)
}
```

### 4. Delete and Restructure (Advanced)

For nodes with two children, replace with successor or predecessor.

```go
// DeleteWithRestructure handles all deletion cases including nodes with two children
func (bt *BinaryTree) DeleteWithRestructure(data int) bool {
    if bt.root == nil {
        return false
    }
    
    bt.root = bt.deleteNodeHelper(bt.root, data)
    return true
}

func (bt *BinaryTree) deleteNodeHelper(node *TreeNode, data int) *TreeNode {
    if node == nil {
        return nil
    }
    
    if node.data == data {
        // Case 1: Leaf node
        if node.left == nil && node.right == nil {
            return nil
        }
        
        // Case 2: One child
        if node.left == nil {
            return node.right
        }
        if node.right == nil {
            return node.left
        }
        
        // Case 3: Two children - replace with inorder successor
        successor := bt.findMinNode(node.right)
        node.data = successor.data
        node.right = bt.deleteNodeHelper(node.right, successor.data)
        
        return node
    }
    
    // Recursively delete from subtrees
    node.left = bt.deleteNodeHelper(node.left, data)
    node.right = bt.deleteNodeHelper(node.right, data)
    
    return node
}

// findMinNode finds the leftmost (minimum) node in subtree
func (bt *BinaryTree) findMinNode(node *TreeNode) *TreeNode {
    for node.left != nil {
        node = node.left
    }
    return node
}
```

### 5. Delete by Position

Delete node at specific level-order position.

```go
// DeleteByPosition removes node at specific position (0-indexed level order)
func (bt *BinaryTree) DeleteByPosition(position int) bool {
    if bt.root == nil || position < 0 {
        return false
    }
    
    if position == 0 {
        // Deleting root - need to restructure
        if bt.root.left == nil && bt.root.right == nil {
            bt.root = nil
        } else if bt.root.left == nil {
            bt.root = bt.root.right
        } else if bt.root.right == nil {
            bt.root = bt.root.left
        } else {
            // Find rightmost node to replace root
            rightmost := bt.findRightmostNode()
            bt.deleteRightmostNode()
            rightmost.left = bt.root.left
            rightmost.right = bt.root.right
            bt.root = rightmost
        }
        return true
    }
    
    // Find parent of target position
    parentPos := (position - 1) / 2
    isLeftChild := position%2 == 1
    
    parent := bt.getNodeAtPosition(parentPos)
    if parent == nil {
        return false
    }
    
    if isLeftChild && parent.left != nil {
        parent.left = nil
        return true
    } else if !isLeftChild && parent.right != nil {
        parent.right = nil
        return true
    }
    
    return false
}

// Helper to get node at specific level-order position
func (bt *BinaryTree) getNodeAtPosition(position int) *TreeNode {
    if bt.root == nil || position < 0 {
        return nil
    }
    
    queue := []*TreeNode{bt.root}
    currentPos := 0
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if currentPos == position {
            return node
        }
        
        if node.left != nil {
            queue = append(queue, node.left)
        }
        if node.right != nil {
            queue = append(queue, node.right)
        }
        
        currentPos++
    }
    
    return nil
}

func (bt *BinaryTree) findRightmostNode() *TreeNode {
    if bt.root == nil {
        return nil
    }
    
    queue := []*TreeNode{bt.root}
    var last *TreeNode
    
    for len(queue) > 0 {
        last = queue[0]
        queue = queue[1:]
        
        if last.left != nil {
            queue = append(queue, last.left)
        }
        if last.right != nil {
            queue = append(queue, last.right)
        }
    }
    
    return last
}

func (bt *BinaryTree) deleteRightmostNode() {
    if bt.root == nil {
        return
    }
    
    if bt.root.left == nil && bt.root.right == nil {
        bt.root = nil
        return
    }
    
    queue := []*TreeNode{bt.root}
    var prev *TreeNode
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.left != nil {
            if node.left.left == nil && node.left.right == nil && len(queue) == 0 {
                node.left = nil
                return
            }
            queue = append(queue, node.left)
        }
        
        if node.right != nil {
            if node.right.left == nil && node.right.right == nil && len(queue) == 0 {
                node.right = nil
                return
            }
            queue = append(queue, node.right)
        }
        
        prev = node
    }
}
```

## TypeScript Implementation

### Complete Deletion Implementation

```typescript
class BinaryTree<T> {
    private root: TreeNode<T> | null;
    
    constructor() {
        this.root = null;
    }
    
    // Delete leaf nodes only
    public deleteLeaf(data: T): boolean {
        if (!this.root) return false;
        
        // Special case: root is leaf
        if (this.root.data === data && !this.root.left && !this.root.right) {
            this.root = null;
            return true;
        }
        
        return this.deleteLeafHelper(this.root, data);
    }
    
    private deleteLeafHelper(node: TreeNode<T> | null, data: T): boolean {
        if (!node) return false;
        
        // Check left child
        if (node.left?.data === data && !node.left.left && !node.left.right) {
            node.left = null;
            return true;
        }
        
        // Check right child
        if (node.right?.data === data && !node.right.left && !node.right.right) {
            node.right = null;
            return true;
        }
        
        return this.deleteLeafHelper(node.left, data) || 
               this.deleteLeafHelper(node.right, data);
    }
    
    // Delete subtree
    public deleteSubtree(data: T): boolean {
        if (!this.root) return false;
        
        if (this.root.data === data) {
            this.root = null;
            return true;
        }
        
        return this.deleteSubtreeHelper(this.root, data);
    }
    
    private deleteSubtreeHelper(node: TreeNode<T> | null, data: T): boolean {
        if (!node) return false;
        
        if (node.left?.data === data) {
            node.left = null;
            return true;
        }
        
        if (node.right?.data === data) {
            node.right = null;
            return true;
        }
        
        return this.deleteSubtreeHelper(node.left, data) || 
               this.deleteSubtreeHelper(node.right, data);
    }
    
    // Delete with restructuring
    public deleteWithRestructure(data: T): boolean {
        if (!this.root) return false;
        
        this.root = this.deleteNodeHelper(this.root, data);
        return true;
    }
    
    private deleteNodeHelper(node: TreeNode<T> | null, data: T): TreeNode<T> | null {
        if (!node) return null;
        
        if (node.data === data) {
            // Case 1: Leaf node
            if (!node.left && !node.right) {
                return null;
            }
            
            // Case 2: One child
            if (!node.left) return node.right;
            if (!node.right) return node.left;
            
            // Case 3: Two children - replace with inorder successor
            const successor = this.findMinNode(node.right);
            node.data = successor.data;
            node.right = this.deleteNodeHelper(node.right, successor.data);
            
            return node;
        }
        
        node.left = this.deleteNodeHelper(node.left, data);
        node.right = this.deleteNodeHelper(node.right, data);
        
        return node;
    }
    
    private findMinNode(node: TreeNode<T>): TreeNode<T> {
        while (node.left) {
            node = node.left;
        }
        return node;
    }
    
    // Delete by position
    public deleteByPosition(position: number): boolean {
        if (!this.root || position < 0) return false;
        
        if (position === 0) {
            if (!this.root.left && !this.root.right) {
                this.root = null;
            } else if (!this.root.left) {
                this.root = this.root.right;
            } else if (!this.root.right) {
                this.root = this.root.left;
            } else {
                // Complex root deletion - promote rightmost node
                const rightmost = this.findRightmostNode();
                this.deleteRightmostNode();
                if (rightmost) {
                    rightmost.left = this.root.left;
                    rightmost.right = this.root.right;
                    this.root = rightmost;
                }
            }
            return true;
        }
        
        const parentPos = Math.floor((position - 1) / 2);
        const isLeftChild = position % 2 === 1;
        
        const parent = this.getNodeAtPosition(parentPos);
        if (!parent) return false;
        
        if (isLeftChild && parent.left) {
            parent.left = null;
            return true;
        } else if (!isLeftChild && parent.right) {
            parent.right = null;
            return true;
        }
        
        return false;
    }
    
    private getNodeAtPosition(position: number): TreeNode<T> | null {
        if (!this.root || position < 0) return null;
        
        const queue: TreeNode<T>[] = [this.root];
        let currentPos = 0;
        
        while (queue.length > 0) {
            const node = queue.shift()!;
            
            if (currentPos === position) {
                return node;
            }
            
            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
            
            currentPos++;
        }
        
        return null;
    }
    
    private findRightmostNode(): TreeNode<T> | null {
        if (!this.root) return null;
        
        const queue: TreeNode<T>[] = [this.root];
        let last: TreeNode<T> | null = null;
        
        while (queue.length > 0) {
            last = queue.shift()!;
            
            if (last.left) queue.push(last.left);
            if (last.right) queue.push(last.right);
        }
        
        return last;
    }
    
    private deleteRightmostNode(): void {
        if (!this.root) return;
        
        if (!this.root.left && !this.root.right) {
            this.root = null;
            return;
        }
        
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const node = queue.shift()!;
            
            if (node.left) {
                if (!node.left.left && !node.left.right && queue.length === 0) {
                    node.left = null;
                    return;
                }
                queue.push(node.left);
            }
            
            if (node.right) {
                if (!node.right.left && !node.right.right && queue.length === 0) {
                    node.right = null;
                    return;
                }
                queue.push(node.right);
            }
        }
    }
    
    // Clear entire tree
    public clear(): void {
        this.root = null;
    }
    
    // Delete all nodes with specific value
    public deleteAll(data: T): number {
        let count = 0;
        while (this.deleteWithRestructure(data)) {
            count++;
        }
        return count;
    }
}
```

## Usage Examples

### Go Examples

```go
func main() {
    bt := NewBinaryTree()
    
    // Build test tree
    elements := []int{1, 2, 3, 4, 5, 6, 7}
    for _, element := range elements {
        bt.Insert(element)
    }
    
    fmt.Println("Original tree size:", bt.Size())
    
    // Delete leaf node
    bt.DeleteLeaf(4)
    fmt.Println("After deleting leaf 4:", bt.Size())
    
    // Delete subtree
    bt.DeleteSubtree(2)
    fmt.Println("After deleting subtree at 2:", bt.Size())
    
    // Delete with restructure
    bt.DeleteWithRestructure(1)
    fmt.Println("After deleting root with restructure:", bt.Size())
}
```

### TypeScript Examples

```typescript
const bt = new BinaryTree<number>();

// Build test tree
[1, 2, 3, 4, 5, 6, 7].forEach(x => bt.insert(x));

console.log("Original tree size:", bt.size());

// Delete leaf node
bt.deleteLeaf(4);
console.log("After deleting leaf 4:", bt.size());

// Delete subtree
bt.deleteSubtree(2);
console.log("After deleting subtree at 2:", bt.size());

// Delete by position
bt.deleteByPosition(0); // Delete root
console.log("After deleting root by position:", bt.size());

// Clear entire tree
bt.clear();
console.log("After clearing:", bt.size());
```

## Deletion Strategy Comparison

| Strategy | Use Case | Preserves Structure | Time Complexity |
|----------|----------|-------------------|----------------|
| Delete Leaf | Safe removal | Yes | O(n) |
| Promote Child | Single child nodes | Partially | O(n) |
| Delete Subtree | Remove branch | No | O(1) |
| Restructure | Complete deletion | Yes | O(n) |
| By Position | Index-based | Partially | O(n) |

## Best Practices

1. **Choose appropriate strategy**:
   - Leaf deletion for safe operations
   - Subtree deletion for branch removal
   - Restructure for maintaining tree integrity

2. **Handle edge cases**:
   - Empty trees
   - Root deletion
   - Nodes with two children

3. **Memory management**:
   - Ensure deleted nodes are properly cleaned up
   - Avoid memory leaks in manual memory management

4. **Maintain tree properties**:
   - Consider impact on tree balance
   - Preserve any required invariants