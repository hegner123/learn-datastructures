# Binary Tree Operations and Utility Methods

## Overview

Binary trees support various utility operations that provide information about tree structure, perform comparisons, and enable advanced manipulations. These operations are essential for tree analysis and algorithm implementation.

## Basic Tree Properties

### Size Operations

```go
// Size returns the number of nodes in the binary tree
func (bt *BinaryTree) Size() int {
    return bt.sizeHelper(bt.root)
}

func (bt *BinaryTree) sizeHelper(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    return 1 + bt.sizeHelper(node.left) + bt.sizeHelper(node.right)
}

// CountLeaves returns the number of leaf nodes
func (bt *BinaryTree) CountLeaves() int {
    return bt.countLeavesHelper(bt.root)
}

func (bt *BinaryTree) countLeavesHelper(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    if node.left == nil && node.right == nil {
        return 1
    }
    
    return bt.countLeavesHelper(node.left) + bt.countLeavesHelper(node.right)
}

// CountInternalNodes returns the number of non-leaf nodes
func (bt *BinaryTree) CountInternalNodes() int {
    return bt.Size() - bt.CountLeaves()
}
```

### Height and Depth Operations

```go
// Height returns the height of the binary tree
func (bt *BinaryTree) Height() int {
    return bt.heightHelper(bt.root)
}

func (bt *BinaryTree) heightHelper(node *TreeNode) int {
    if node == nil {
        return -1
    }
    
    leftHeight := bt.heightHelper(node.left)
    rightHeight := bt.heightHelper(node.right)
    
    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}

// Depth returns the depth of a specific node
func (bt *BinaryTree) Depth(data int) int {
    return bt.depthHelper(bt.root, data, 0)
}

func (bt *BinaryTree) depthHelper(node *TreeNode, data int, currentDepth int) int {
    if node == nil {
        return -1
    }
    
    if node.data == data {
        return currentDepth
    }
    
    leftDepth := bt.depthHelper(node.left, data, currentDepth+1)
    if leftDepth != -1 {
        return leftDepth
    }
    
    return bt.depthHelper(node.right, data, currentDepth+1)
}

// MaxWidth returns the maximum width at any level
func (bt *BinaryTree) MaxWidth() int {
    if bt.root == nil {
        return 0
    }
    
    maxWidth := 0
    queue := []*TreeNode{bt.root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        if levelSize > maxWidth {
            maxWidth = levelSize
        }
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.left != nil {
                queue = append(queue, node.left)
            }
            if node.right != nil {
                queue = append(queue, node.right)
            }
        }
    }
    
    return maxWidth
}
```

## Tree Balance and Structure Analysis

```go
// IsBalanced checks if the binary tree is balanced
func (bt *BinaryTree) IsBalanced() bool {
    _, balanced := bt.isBalancedHelper(bt.root)
    return balanced
}

func (bt *BinaryTree) isBalancedHelper(node *TreeNode) (int, bool) {
    if node == nil {
        return -1, true
    }
    
    leftHeight, leftBalanced := bt.isBalancedHelper(node.left)
    rightHeight, rightBalanced := bt.isBalancedHelper(node.right)
    
    if !leftBalanced || !rightBalanced {
        return 0, false
    }
    
    diff := leftHeight - rightHeight
    if diff < 0 {
        diff = -diff
    }
    
    balanced := diff <= 1
    height := leftHeight
    if rightHeight > leftHeight {
        height = rightHeight
    }
    
    return height + 1, balanced
}

// IsComplete checks if the tree is complete
func (bt *BinaryTree) IsComplete() bool {
    if bt.root == nil {
        return true
    }
    
    queue := []*TreeNode{bt.root}
    foundNonFull := false
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.left != nil {
            if foundNonFull {
                return false
            }
            queue = append(queue, node.left)
        } else {
            foundNonFull = true
        }
        
        if node.right != nil {
            if foundNonFull {
                return false
            }
            queue = append(queue, node.right)
        } else {
            foundNonFull = true
        }
    }
    
    return true
}

// IsPerfect checks if the tree is perfect
func (bt *BinaryTree) IsPerfect() bool {
    height := bt.Height()
    return bt.isPerfectHelper(bt.root, height, 0)
}

func (bt *BinaryTree) isPerfectHelper(node *TreeNode, height, level int) bool {
    if node == nil {
        return true
    }
    
    // Leaf node
    if node.left == nil && node.right == nil {
        return height == level
    }
    
    // Node with only one child
    if node.left == nil || node.right == nil {
        return false
    }
    
    return bt.isPerfectHelper(node.left, height, level+1) &&
           bt.isPerfectHelper(node.right, height, level+1)
}

// IsFull checks if every node has 0 or 2 children
func (bt *BinaryTree) IsFull() bool {
    return bt.isFullHelper(bt.root)
}

func (bt *BinaryTree) isFullHelper(node *TreeNode) bool {
    if node == nil {
        return true
    }
    
    // Leaf node
    if node.left == nil && node.right == nil {
        return true
    }
    
    // Node with both children
    if node.left != nil && node.right != nil {
        return bt.isFullHelper(node.left) && bt.isFullHelper(node.right)
    }
    
    // Node with only one child
    return false
}
```

## Search and Path Operations

```go
// Search finds a value in the binary tree
func (bt *BinaryTree) Search(data int) bool {
    return bt.searchHelper(bt.root, data)
}

func (bt *BinaryTree) searchHelper(node *TreeNode, data int) bool {
    if node == nil {
        return false
    }
    
    if node.data == data {
        return true
    }
    
    return bt.searchHelper(node.left, data) || bt.searchHelper(node.right, data)
}

// FindPath finds path from root to target node
func (bt *BinaryTree) FindPath(target int) []int {
    var path []int
    if bt.findPathHelper(bt.root, target, &path) {
        return path
    }
    return []int{}
}

func (bt *BinaryTree) findPathHelper(node *TreeNode, target int, path *[]int) bool {
    if node == nil {
        return false
    }
    
    *path = append(*path, node.data)
    
    if node.data == target {
        return true
    }
    
    if bt.findPathHelper(node.left, target, path) || 
       bt.findPathHelper(node.right, target, path) {
        return true
    }
    
    // Backtrack
    *path = (*path)[:len(*path)-1]
    return false
}

// FindLCA finds the Lowest Common Ancestor of two nodes
func (bt *BinaryTree) FindLCA(p, q int) *TreeNode {
    return bt.findLCAHelper(bt.root, p, q)
}

func (bt *BinaryTree) findLCAHelper(node *TreeNode, p, q int) *TreeNode {
    if node == nil {
        return nil
    }
    
    if node.data == p || node.data == q {
        return node
    }
    
    leftLCA := bt.findLCAHelper(node.left, p, q)
    rightLCA := bt.findLCAHelper(node.right, p, q)
    
    if leftLCA != nil && rightLCA != nil {
        return node
    }
    
    if leftLCA != nil {
        return leftLCA
    }
    return rightLCA
}

// Distance calculates distance between two nodes
func (bt *BinaryTree) Distance(p, q int) int {
    lca := bt.FindLCA(p, q)
    if lca == nil {
        return -1
    }
    
    distP := bt.distanceFromNode(lca, p)
    distQ := bt.distanceFromNode(lca, q)
    
    if distP == -1 || distQ == -1 {
        return -1
    }
    
    return distP + distQ
}

func (bt *BinaryTree) distanceFromNode(node *TreeNode, target int) int {
    if node == nil {
        return -1
    }
    
    if node.data == target {
        return 0
    }
    
    leftDist := bt.distanceFromNode(node.left, target)
    if leftDist != -1 {
        return leftDist + 1
    }
    
    rightDist := bt.distanceFromNode(node.right, target)
    if rightDist != -1 {
        return rightDist + 1
    }
    
    return -1
}
```

## Tree Transformation Operations

```go
// Mirror creates a mirror image of the binary tree
func (bt *BinaryTree) Mirror() {
    bt.mirrorHelper(bt.root)
}

func (bt *BinaryTree) mirrorHelper(node *TreeNode) {
    if node == nil {
        return
    }
    
    // Swap left and right children
    node.left, node.right = node.right, node.left
    
    // Recursively mirror subtrees
    bt.mirrorHelper(node.left)
    bt.mirrorHelper(node.right)
}

// Clone creates a deep copy of the tree
func (bt *BinaryTree) Clone() *BinaryTree {
    newTree := NewBinaryTree()
    newTree.root = bt.cloneHelper(bt.root)
    return newTree
}

func (bt *BinaryTree) cloneHelper(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }
    
    newNode := &TreeNode{
        data:  node.data,
        left:  bt.cloneHelper(node.left),
        right: bt.cloneHelper(node.right),
    }
    
    return newNode
}

// Flatten converts tree to linked list (preorder)
func (bt *BinaryTree) Flatten() {
    bt.flattenHelper(bt.root)
}

func (bt *BinaryTree) flattenHelper(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }
    
    leftTail := bt.flattenHelper(node.left)
    rightTail := bt.flattenHelper(node.right)
    
    if node.left != nil {
        leftTail.right = node.right
        node.right = node.left
        node.left = nil
    }
    
    if rightTail != nil {
        return rightTail
    }
    if leftTail != nil {
        return leftTail
    }
    return node
}
```

## TypeScript Implementation

### Complete Operations Implementation

```typescript
class BinaryTree<T> {
    private root: TreeNode<T> | null;
    
    constructor() {
        this.root = null;
    }
    
    // Get size (number of nodes)
    public size(): number {
        return this.sizeHelper(this.root);
    }
    
    private sizeHelper(node: TreeNode<T> | null): number {
        if (!node) return 0;
        return 1 + this.sizeHelper(node.left) + this.sizeHelper(node.right);
    }
    
    // Get height of tree
    public height(): number {
        return this.heightHelper(this.root);
    }
    
    private heightHelper(node: TreeNode<T> | null): number {
        if (!node) return -1;
        
        const leftHeight = this.heightHelper(node.left);
        const rightHeight = this.heightHelper(node.right);
        
        return Math.max(leftHeight, rightHeight) + 1;
    }
    
    // Count leaf nodes
    public countLeaves(): number {
        return this.countLeavesHelper(this.root);
    }
    
    private countLeavesHelper(node: TreeNode<T> | null): number {
        if (!node) return 0;
        if (!node.left && !node.right) return 1;
        
        return this.countLeavesHelper(node.left) + this.countLeavesHelper(node.right);
    }
    
    // Count internal nodes
    public countInternalNodes(): number {
        return this.size() - this.countLeaves();
    }
    
    // Get depth of specific node
    public depth(data: T): number {
        return this.depthHelper(this.root, data, 0);
    }
    
    private depthHelper(node: TreeNode<T> | null, data: T, currentDepth: number): number {
        if (!node) return -1;
        if (node.data === data) return currentDepth;
        
        const leftDepth = this.depthHelper(node.left, data, currentDepth + 1);
        if (leftDepth !== -1) return leftDepth;
        
        return this.depthHelper(node.right, data, currentDepth + 1);
    }
    
    // Get maximum width
    public maxWidth(): number {
        if (!this.root) return 0;
        
        let maxWidth = 0;
        const queue: TreeNode<T>[] = [this.root];
        
        while (queue.length > 0) {
            const levelSize = queue.length;
            maxWidth = Math.max(maxWidth, levelSize);
            
            for (let i = 0; i < levelSize; i++) {
                const node = queue.shift()!;
                
                if (node.left) queue.push(node.left);
                if (node.right) queue.push(node.right);
            }
        }
        
        return maxWidth;
    }
    
    // Check if tree is balanced
    public isBalanced(): boolean {
        const result = this.isBalancedHelper(this.root);
        return result.balanced;
    }
    
    private isBalancedHelper(node: TreeNode<T> | null): { height: number; balanced: boolean } {
        if (!node) return { height: -1, balanced: true };
        
        const left = this.isBalancedHelper(node.left);
        const right = this.isBalancedHelper(node.right);
        
        if (!left.balanced || !right.balanced) {
            return { height: 0, balanced: false };
        }
        
        const balanced = Math.abs(left.height - right.height) <= 1;
        const height = Math.max(left.height, right.height) + 1;
        
        return { height, balanced };
    }
    
    // Check if tree is complete
    public isComplete(): boolean {
        if (!this.root) return true;
        
        const queue: TreeNode<T>[] = [this.root];
        let foundNonFull = false;
        
        while (queue.length > 0) {
            const node = queue.shift()!;
            
            if (node.left) {
                if (foundNonFull) return false;
                queue.push(node.left);
            } else {
                foundNonFull = true;
            }
            
            if (node.right) {
                if (foundNonFull) return false;
                queue.push(node.right);
            } else {
                foundNonFull = true;
            }
        }
        
        return true;
    }
    
    // Check if tree is perfect
    public isPerfect(): boolean {
        const height = this.height();
        return this.isPerfectHelper(this.root, height, 0);
    }
    
    private isPerfectHelper(node: TreeNode<T> | null, height: number, level: number): boolean {
        if (!node) return true;
        
        if (!node.left && !node.right) {
            return height === level;
        }
        
        if (!node.left || !node.right) {
            return false;
        }
        
        return this.isPerfectHelper(node.left, height, level + 1) &&
               this.isPerfectHelper(node.right, height, level + 1);
    }
    
    // Check if tree is full
    public isFull(): boolean {
        return this.isFullHelper(this.root);
    }
    
    private isFullHelper(node: TreeNode<T> | null): boolean {
        if (!node) return true;
        
        if (!node.left && !node.right) return true;
        
        if (node.left && node.right) {
            return this.isFullHelper(node.left) && this.isFullHelper(node.right);
        }
        
        return false;
    }
    
    // Search for a value
    public search(data: T): boolean {
        return this.searchHelper(this.root, data);
    }
    
    private searchHelper(node: TreeNode<T> | null, data: T): boolean {
        if (!node) return false;
        if (node.data === data) return true;
        
        return this.searchHelper(node.left, data) || 
               this.searchHelper(node.right, data);
    }
    
    // Find path from root to target
    public findPath(target: T): T[] {
        const path: T[] = [];
        if (this.findPathHelper(this.root, target, path)) {
            return path;
        }
        return [];
    }
    
    private findPathHelper(node: TreeNode<T> | null, target: T, path: T[]): boolean {
        if (!node) return false;
        
        path.push(node.data);
        
        if (node.data === target) return true;
        
        if (this.findPathHelper(node.left, target, path) || 
            this.findPathHelper(node.right, target, path)) {
            return true;
        }
        
        path.pop(); // Backtrack
        return false;
    }
    
    // Find lowest common ancestor
    public findLCA(p: T, q: T): TreeNode<T> | null {
        return this.findLCAHelper(this.root, p, q);
    }
    
    private findLCAHelper(node: TreeNode<T> | null, p: T, q: T): TreeNode<T> | null {
        if (!node) return null;
        if (node.data === p || node.data === q) return node;
        
        const leftLCA = this.findLCAHelper(node.left, p, q);
        const rightLCA = this.findLCAHelper(node.right, p, q);
        
        if (leftLCA && rightLCA) return node;
        return leftLCA || rightLCA;
    }
    
    // Calculate distance between two nodes
    public distance(p: T, q: T): number {
        const lca = this.findLCA(p, q);
        if (!lca) return -1;
        
        const distP = this.distanceFromNode(lca, p);
        const distQ = this.distanceFromNode(lca, q);
        
        if (distP === -1 || distQ === -1) return -1;
        return distP + distQ;
    }
    
    private distanceFromNode(node: TreeNode<T> | null, target: T): number {
        if (!node) return -1;
        if (node.data === target) return 0;
        
        const leftDist = this.distanceFromNode(node.left, target);
        if (leftDist !== -1) return leftDist + 1;
        
        const rightDist = this.distanceFromNode(node.right, target);
        if (rightDist !== -1) return rightDist + 1;
        
        return -1;
    }
    
    // Mirror the tree
    public mirror(): void {
        this.mirrorHelper(this.root);
    }
    
    private mirrorHelper(node: TreeNode<T> | null): void {
        if (!node) return;
        
        [node.left, node.right] = [node.right, node.left];
        
        this.mirrorHelper(node.left);
        this.mirrorHelper(node.right);
    }
    
    // Clone the tree
    public clone(): BinaryTree<T> {
        const newTree = new BinaryTree<T>();
        newTree.root = this.cloneHelper(this.root);
        return newTree;
    }
    
    private cloneHelper(node: TreeNode<T> | null): TreeNode<T> | null {
        if (!node) return null;
        
        const newNode = new TreeNode<T>(node.data);
        newNode.left = this.cloneHelper(node.left);
        newNode.right = this.cloneHelper(node.right);
        
        return newNode;
    }
    
    // Get all ancestors of a node
    public findAncestors(target: T): T[] {
        const ancestors: T[] = [];
        this.findAncestorsHelper(this.root, target, ancestors);
        return ancestors;
    }
    
    private findAncestorsHelper(node: TreeNode<T> | null, target: T, ancestors: T[]): boolean {
        if (!node) return false;
        if (node.data === target) return true;
        
        if (this.findAncestorsHelper(node.left, target, ancestors) ||
            this.findAncestorsHelper(node.right, target, ancestors)) {
            ancestors.push(node.data);
            return true;
        }
        
        return false;
    }
    
    // Check if two trees are identical
    public isIdentical(other: BinaryTree<T>): boolean {
        return this.areNodesIdentical(this.root, other.root);
    }
    
    private areNodesIdentical(node1: TreeNode<T> | null, node2: TreeNode<T> | null): boolean {
        if (!node1 && !node2) return true;
        if (!node1 || !node2) return false;
        
        return node1.data === node2.data &&
               this.areNodesIdentical(node1.left, node2.left) &&
               this.areNodesIdentical(node1.right, node2.right);
    }
    
    // Get diameter of tree (longest path between any two nodes)
    public diameter(): number {
        const result = { diameter: 0 };
        this.diameterHelper(this.root, result);
        return result.diameter;
    }
    
    private diameterHelper(node: TreeNode<T> | null, result: { diameter: number }): number {
        if (!node) return -1;
        
        const leftHeight = this.diameterHelper(node.left, result);
        const rightHeight = this.diameterHelper(node.right, result);
        
        const currentDiameter = leftHeight + rightHeight + 2;
        result.diameter = Math.max(result.diameter, currentDiameter);
        
        return Math.max(leftHeight, rightHeight) + 1;
    }
    
    // Convert to array representation
    public toArray(): (T | null)[] {
        if (!this.root) return [];
        
        const result: (T | null)[] = [];
        const queue: (TreeNode<T> | null)[] = [this.root];
        
        while (queue.length > 0) {
            const current = queue.shift();
            
            if (current) {
                result.push(current.data);
                queue.push(current.left);
                queue.push(current.right);
            } else {
                result.push(null);
            }
        }
        
        // Remove trailing nulls
        while (result.length > 0 && result[result.length - 1] === null) {
            result.pop();
        }
        
        return result;
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
    
    // Tree properties
    fmt.Printf("Size: %d\n", bt.Size())
    fmt.Printf("Height: %d\n", bt.Height())
    fmt.Printf("Leaves: %d\n", bt.CountLeaves())
    fmt.Printf("Internal nodes: %d\n", bt.CountInternalNodes())
    fmt.Printf("Max width: %d\n", bt.MaxWidth())
    
    // Tree structure analysis
    fmt.Printf("Is balanced: %t\n", bt.IsBalanced())
    fmt.Printf("Is complete: %t\n", bt.IsComplete())
    fmt.Printf("Is perfect: %t\n", bt.IsPerfect())
    fmt.Printf("Is full: %t\n", bt.IsFull())
    
    // Search operations
    fmt.Printf("Search 5: %t\n", bt.Search(5))
    fmt.Printf("Depth of 5: %d\n", bt.Depth(5))
    fmt.Printf("Path to 5: %v\n", bt.FindPath(5))
    
    // Advanced operations
    lca := bt.FindLCA(4, 5)
    if lca != nil {
        fmt.Printf("LCA of 4 and 5: %d\n", lca.data)
    }
    fmt.Printf("Distance between 4 and 5: %d\n", bt.Distance(4, 5))
    
    // Clone and mirror
    cloned := bt.Clone()
    fmt.Printf("Cloned tree size: %d\n", cloned.Size())
    
    bt.Mirror()
    fmt.Printf("After mirroring, level order: %v\n", bt.LevelOrderTraversal())
}
```

### TypeScript Examples

```typescript
const bt = new BinaryTree<number>();

// Build test tree
[1, 2, 3, 4, 5, 6, 7].forEach(x => bt.insert(x));

// Tree properties
console.log("Size:", bt.size());
console.log("Height:", bt.height());
console.log("Leaves:", bt.countLeaves());
console.log("Internal nodes:", bt.countInternalNodes());
console.log("Max width:", bt.maxWidth());

// Structure analysis
console.log("Is balanced:", bt.isBalanced());
console.log("Is complete:", bt.isComplete());
console.log("Is perfect:", bt.isPerfect());
console.log("Is full:", bt.isFull());

// Search operations
console.log("Search 5:", bt.search(5));
console.log("Depth of 5:", bt.depth(5));
console.log("Path to 5:", bt.findPath(5));

// Advanced operations
const lca = bt.findLCA(4, 5);
console.log("LCA of 4 and 5:", lca?.data);
console.log("Distance between 4 and 5:", bt.distance(4, 5));
console.log("Ancestors of 5:", bt.findAncestors(5));
console.log("Diameter:", bt.diameter());

// Tree transformation
const cloned = bt.clone();
console.log("Cloned tree size:", cloned.size());
console.log("Trees are identical:", bt.isIdentical(cloned));

bt.mirror();
console.log("After mirroring:", bt.levelOrderTraversal());

console.log("Array representation:", bt.toArray());
```

## Operations Complexity

| Operation | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Size | O(n) | O(h) |
| Height | O(n) | O(h) |
| Search | O(n) | O(h) |
| IsBalanced | O(n) | O(h) |
| IsComplete | O(n) | O(w) |
| FindPath | O(n) | O(h) |
| FindLCA | O(n) | O(h) |
| Distance | O(n) | O(h) |
| Mirror | O(n) | O(h) |
| Clone | O(n) | O(n) |
| Diameter | O(n) | O(h) |

Where n = number of nodes, h = height, w = maximum width

## Applications

- **Size/Height**: Memory estimation, tree visualization
- **Balance checking**: Performance optimization, tree maintenance
- **Path finding**: Navigation, relationship analysis
- **LCA**: Range queries, tree-based algorithms
- **Mirroring**: Image processing, symmetry operations
- **Cloning**: Backup, parallel processing
- **Structure validation**: Data integrity, format compliance